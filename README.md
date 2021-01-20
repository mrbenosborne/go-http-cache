# Go HTTP Cache
A HTTP cache server to improve the performance of applications with frequent/shared access to a resource.

![Build & Publish](https://github.com/mrbenosborne/go-http-cache/workflows/Build%20&%20Publish/badge.svg) ![Go](https://github.com/mrbenosborne/go-http-cache/workflows/Go/badge.svg)

# Docker
A docker image is available from [Docker Hub](https://hub.docker.com/repository/docker/mrbenosborne/go-http-cache) and is updated whenever changes are pushed into
the *main* branch.

## Run
Quickly test it out by spinning up a docker container.

```
docker run --name go-http-cache -p 80:8901 mrbenosborne/go-http-cache:latest
```

# Postman
Head over to our Postman collection for all the documented requests/responses.

[View Postman API](https://documenter.getpostman.com/view/14265644/TVzYgZvF)