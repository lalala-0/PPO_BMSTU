from locust import HttpUser, TaskSet, task, between

class UserBehavior(TaskSet):
    @task
    def get_ratings(self):
        # Выполнение GET-запроса
        self.client.get("/ui/ratings/27a04043-5b70-44d7-8a3b-58387a53d228")

class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 3)  # Задержка между запросами от 1 до 3 секунд
    host = "http://127.0.0.1:8081"  # Базовый URL
