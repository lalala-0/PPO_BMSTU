from gevent import monkey
monkey.patch_all()

import pymongo
from locust import User, task, between, events
import time
from prometheus_client import start_http_server, Histogram, Counter, Gauge

# Запуск Prometheus HTTP сервера на порту 8002
start_http_server(8002, addr='0.0.0.0')

# Определение метрик Prometheus для MongoDB
mongo_request_time = Histogram('mongodb_request_duration_seconds', 'MongoDB query duration',
                               buckets=[0.1, 0.5, 1, 2, 5, 10])
mongo_errors = Counter('mongodb_query_errors_total', 'Total number of MongoDB query errors')
mongo_rps = Counter('mongodb_requests_total', 'MongoDB requests count')
mongo_concurrent_queries = Gauge('mongodb_concurrent_queries', 'Number of concurrent MongoDB queries')

class MongoBenchmark(User):
    wait_time = between(1, 2)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        try:
            self.mongo_client = pymongo.MongoClient("mongodb://localhost:27027/")
            self.mongo_db = self.mongo_client["ppo_db"]
            self.mongo_collection = self.mongo_db["ratings"]
            print("MongoDB connection established")
        except Exception as e:
            mongo_errors.inc()
            events.request.fire(
                request_type="MONGODB",
                name="MongoDB ratings.query",
                response_time=0,
                response_length=0,
                exception=e
            )
            print(f"MongoDB query error: {e}")
            self.mongo_client = None

    @task
    def test_mongodb_query(self):
        if self.mongo_client is None:
            mongo_errors.inc()
            return

        mongo_concurrent_queries.inc()

        start_time = time.time()
        try:
            ratings = list(self.mongo_collection.find().limit(1000))
            duration = time.time() - start_time

            # Запись в Prometheus
            mongo_request_time.observe(duration)  # Запись в гистограмму
            mongo_rps.inc()  # Запись RPS

            events.request.fire(
                request_type="MONGODB",
                name="MongoDB ratings.query",
                response_time=duration * 1000,  # в миллисекундах
                response_length=len(ratings),
                success=True
            )
            print(f"MongoDB query time: {duration} seconds, rows: {len(ratings)}")
        except Exception as e:
            mongo_errors.inc()
            events.request.fire(
                request_type="MONGODB",
                name="MongoDB ratings.query",
                response_time=0,
                response_length=0,
                exception=e
            )
            print(f"MongoDB query error: {e}")
        finally:
            mongo_concurrent_queries.dec()
