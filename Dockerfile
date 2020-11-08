# docker file for ayolescore app
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/simple-21
WORKDIR /go/src/github.com/renosyah/simple-21
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /api
RUN rm -rf /cmd
RUN rm -rf /model
RUN rm -rf /router
RUN rm -rf /vendor
RUN rm -rf /util
RUN rm .dockerignore
RUN rm .gitignore
RUN rm Dockerfile
RUN rm go.mod
RUN rm go.sum
RUN rm heroku.yml
RUN rm main.go
EXPOSE 8000
EXPOSE 80
CMD ./main --config=.server.toml
MAINTAINER syahputrareno975@gmail.com
