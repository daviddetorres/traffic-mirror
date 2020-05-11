# Traffic Mirror
An API based server traffic generator. 

Traffic Mirror can generate different sizes of correct responses and errors with different codes. This is useful if you are testing load balancers or kubernetes services.

# Installation
From docker image:
```
docker run -d --rm \
        --name traffic-mirror \
        -p 8080:8080 daviddetorres/traffic-mirror:latest
```

Building your source code
```
git clone https://github.com/daviddetorres/traffic-mirror
go build main.go -o traffic-mirror
```

# Usage
Onces started, traffic-mirror opens the port 8080 and waits for new petitions:
* /ok/milliseconds/bytes - Returns CODE:200 with random latency from 0 to 'milliseconds' and with a body filled with 0's 'bytes' times.
* /error/code - Returns error with the specified 'code'.

You can generate requests with curl: 
```
curl http://localhost:8080/ok/500/1000
```

Or with load testers like [bombardier](https://github.com/codesenberg/bombardier)
```
docker run -ti --rm \
    --name bombardier \
    --net='host' \
    alpine/bombardier -r 11 -c 10 -d 10s -l http://localhost:8080/ok/100/10000
```

## Locust script
If you want something more sophisticated to generate queries, you can use the locust script and customize it as you wish. 

In command line:
```
locust -f traffic-generator.py --no-web -c 10 -r 10 -H http://localhost:8080
```

With a Docker image:
```
docker run -ti --rm \
  --net='host' \
  --volume $PWD:/mnt/locust \
  -e LOCUSTFILE_PATH=/mnt/locust/traffic-generator.py \
  -e TARGET_URL=http://localhost:8080 \
  -e LOCUST_OPTS="--no-web -c 10 -r 10 " \
  locustio/locust
```
