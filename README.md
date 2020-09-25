# go-mvc-rabbitmq

- This application shows us to usage of rabbitmq with GoLang.

 # Getting Started

## 1. Download or clone this repo with

```
https://github.com/oktydag/go-mvc-rabbitmq.git
```


# Steps For Run
## 1.  Download RabbitMQ Image via Docker

Run RabbitMQ container with default **Username: guest** and **Password: guest**

```
docker run -d --hostname my-rabbit --name my-rabbit -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

If you want to decleare your own username and password;

```
docker run -d --hostname my-rabbit --name my-rabbit -e RABBITMQ_DEFAULT_USER=oktydag -e RABBITMQ_DEFAULT_PASS=123456 -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Then you can look deafult RabbitMQ port : 

```
http://localhost:15672/
```

## 2.  Run it

```
$ go run main.go
```

## 3. Call Services

**Publish Message like this;**
```
$ curl -X POST http://localhost:3000/publish  
```

**Receive Message like this;**
```
$ curl -X GET  http://localhost:3000/receive   
```