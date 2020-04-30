from locust import HttpLocust, TaskSet, task, between

class WebsiteTasks(TaskSet):
    headers = {
        "Content-Type": "application/json"
    }
    creds = {
        "username": "admin",
        "password": "admin"
    }
    def on_start(self):
        self.client.get("/")
    
    @task
    def login(self):
        self.client.get("/")
        
    @task
    def a(self):
        self.client.get("/")
        

class WebsiteUser(HttpLocust):
    task_set = WebsiteTasks
    wait_time = between(5, 300)