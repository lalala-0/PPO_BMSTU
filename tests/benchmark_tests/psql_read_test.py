from gevent import monkey
monkey.patch_all()

import psycopg2
from locust import User, task, between, events
import time
from prometheus_client import start_http_server, Histogram, Counter, Gauge

# Запуск Prometheus HTTP сервера на порту 8001
start_http_server(8001, addr='0.0.0.0')

# Определение метрик Prometheus для PostgreSQL
pg_request_time = Histogram('postgresql_request_duration_seconds', 'PostgreSQL query duration',
                            buckets=[0.1, 0.5, 1, 2, 5, 10])  # добавлены подходящие для процентовые интервалы
pg_errors = Counter('postgresql_query_errors_total', 'Total number of PostgreSQL query errors')
pg_rps = Counter('postgresql_requests_total', 'PostgreSQL requests count')
pg_concurrent_queries = Gauge('postgresql_concurrent_queries', 'Number of concurrent PostgreSQL queries')

class PostgresBenchmark(User):
    wait_time = between(1, 2)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        try:
            # Установление соединения с PostgreSQL
            self.pg_conn = psycopg2.connect(
                dbname="postgres", user="postgres", password="postgres", host="127.0.0.1", port="5444")
            self.pg_cursor = self.pg_conn.cursor()
            print("PostgreSQL connection established")
        except Exception as e:
            pg_errors.inc()
            events.request.fire(
                request_type="POSTGRESQL",
                name="PostgreSQL ratings.query",
                response_time=0,
                response_length=0,
                exception=e
            )
            print(f"PostgreSQL query error: {e}")
            self.pg_conn = None

    @task
    def test_postgresql_query(self):
        if self.pg_conn is None:
            pg_errors.inc()
            return

        pg_concurrent_queries.inc()  # Увеличиваем количество параллельных запросов

        start_time = time.time()
        try:
            self.pg_cursor.execute("SELECT * FROM ratings ORDER BY name limit 1000;")
            result = self.pg_cursor.fetchall()
            duration = time.time() - start_time

            pg_request_time.observe(duration)  # Фиксируем длительность запроса для процентилей
            pg_rps.inc()  # Увеличиваем количество запросов в секунду

            events.request.fire(
                request_type="POSTGRESQL",
                name="PostgreSQL ratings.query",
                response_time=duration * 1000,  # Время в миллисекундах
                response_length=len(result),
                success=True
            )
            print(f"PostgreSQL query time: {duration} seconds, rows: {len(result)}")
        except Exception as e:
            pg_errors.inc()
            events.request.fire(
                request_type="POSTGRESQL",
                name="PostgreSQL ratings.query",
                response_time=0,
                response_length=0,
                exception=e
            )
            print(f"PostgreSQL query error: {e}")
        finally:
            pg_concurrent_queries.dec()  # Уменьшаем количество параллельных запросов
