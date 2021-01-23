# Micro

Micro is a Dead simple boilerplate (Scaffolding?) for Go-based microservices. It comes with a pre-configured proxy load balancer (envoy) and built-in documentation using OpenAPI and Swagger.

# Usage

### Without Docker 
```bash 
# Without docker
cd service && go mod download && make swagger && go build -o micro
./micro --env development --port :8080
```
![alt text](https://github.com/dkvilo/micro/blob/master/images/local.png?raw=true)
![alt text](https://github.com/dkvilo/micro/blob/master/images/redoc.png?raw=true)

### With Docker, and without envoy
```bash 
# With docker, no proxy (production env, no swagger needed)
# build
docker build -t micro-go ./service
# Run
docker run -ti -p 80:8080 micro-go ./micro --port :8080 --env production
```
![alt text](https://github.com/dkvilo/micro/blob/master/images/no-proxy.png?raw=true)

### With Envoy proxy
```bash
# Build
docker-compose build
# Run
docker-compose up
```
![alt text](https://github.com/dkvilo/micro/blob/master/images/proxy.png?raw=true)

- Dependencies:
	- julienschmidt/httprouter
	- go-swagger/go-swagger
	- go-openapi/runtime

