from locust import Locust, TaskSet, task, between, HttpLocust

class MyTaskSet(TaskSet):
    headers = {"Host": "traffic-mirror.io"}

    @task(100)
    def NoLatencyLightWeight(self):
        self.client.request("GET","/ok/0/500", 
                            headers=self.headers)

    @task(100)
    def MediumLatencyLightWeight(self):
        self.client.request("GET","/ok/500/500", 
                            headers=self.headers)

    @task(10)
    def HighLatencyLightWeight(self):
        self.client.request("GET","/ok/5000/500", 
                            headers=self.headers)
    
    @task(100)
    def NoLatencyMediumWeight(self):
        self.client.request("GET","/ok/0/50000", 
                            headers=self.headers)
    
    @task(100)
    def MediumLatencyMediumWeight(self):
        self.client.request("GET","/ok/500/50000", 
                            headers=self.headers)
    
    @task(10)
    def HighLatencyMediumWeight(self):
        self.client.request("GET","/ok/5000/50000", 
                            headers=self.headers)

    @task(20)
    def NoLatencyHeavyWeight(self):
        self.client.request("GET","/ok/0/5000000", 
                            headers=self.headers)
    
    @task(20)
    def MediumLatencyHeavyWeight(self):
        self.client.request("GET","/ok/500/5000000", 
                            headers=self.headers)
   
    @task(10)
    def HighLatencyHeavyWeight(self):
        self.client.request("GET","/ok/5000/5000000", 
                            headers=self.headers)
    
    @task(100)
    def LightWeightPost(self):
        self.client.request("POST","/ok/0/0",  
                            data={"data":"0"*500}, 
                            headers=self.headers)
    
    @task(100)
    def MediumWeightPost(self):
        self.client.request("POST","/ok/0/0", 
                            data={"data":"0"*50000}, 
                            headers=self.headers)
    
    @task(20)
    def HeavyWeightPost(self):
        self.client.request("POST","/ok/0/0", 
                            data={"data":"0"*5000000}, 
                            headers=self.headers)
   
    @task(10)
    def Error404(self):
      self.client.request("GET","/error/404", 
                        headers=self.headers)
   
    @task(10)
    def Error403(self):
      self.client.request("GET","/error/403", 
                        headers=self.headers)

    @task(10)
    def Error500(self):
      self.client.request("GET","/error/500", 
                        headers=self.headers)
    
    @task(10)
    def Error503(self):
      self.client.request("GET","/error/503", 
                        headers=self.headers)

class TrafficMirrorUser(HttpLocust):
    task_set = MyTaskSet    
    wait_time = between(5, 15)