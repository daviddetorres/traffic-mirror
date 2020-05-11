from locust import Locust, TaskSet, task, between, HttpLocust

class MyTaskSet(TaskSet):
  @task(100)
  def NoLatencyLightWeight(self):
      self.client.get("/ok/0/500")

  @task(100)
  def MediumLatencyLightWeight(self):
      self.client.get("/ok/500/500")

  @task(10)
  def HighLatencyLightWeight(self):
      self.client.get("/ok/5000/500")

  @task(100)
  def NoLatencyMediumWeight(self):
      self.client.get("/ok/0/50000")

  @task(100)
  def MediumLatencyMediumWeight(self):
      self.client.get("/ok/500/50000")

  @task(10)
  def HighLatencyMediumWeight(self):
      self.client.get("/ok/5000/50000")
      
  @task(20)
  def NoLatencyHeavyWeight(self):
      self.client.get("/ok/0/5000000")

  @task(20)
  def MediumLatencyHeavyWeight(self):
      self.client.get("/ok/500/5000000")

  @task(10)
  def HighLatencyHeavyWeight(self):
      self.client.get("/ok/5000/5000000")
  @task(10)
  def Error404(self):
    self.client.get("/error/404")

  @task(10)
  def Error403(self):
    self.client.get("/error/403")
  
  @task(10)
  def Error500(self):
    self.client.get("/error/500")

  @task(10)
  def Error503(self):
    self.client.get("/error/503")

class TrafficMirrorUser(HttpLocust):
  task_set = MyTaskSet
  wait_time = between(5, 15)