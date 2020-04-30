# Sample-web golang api for cpu intense

# Install dependencies
```
go mod download
```

# Build
```
go build
```

# Run
```
./sample-web
```

# Using docker
```
docker build . -t sample-web:latest

docker run --publish 8080:8080 sample-web:latest
```

# How to test?
Just send a post request to <url>:8080 with the following details

```
POST http://localhost:8080/login

{
    "username": "admin",
    "password": "admin"
}
```

# What it does?
It is quite simple, it just executes encrypting and decrypting a string. And generates the token. The intention of this app is to put some load on CPU which helps some use cases such as autoscaling pods/vms/hardware/resources.

Note: This also does write logs to stdout - can be useful for any other scenarios.

# Benchmark
When I run the api directly on the host, it is taking around 20-25sec and CPU % is 80-90%.

But If I run the api on docker, it is taking around only 4sec and CPU % is 125%(docker process)

Try putting load using load-test