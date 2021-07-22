FROM golang:1.16 AS build
ARG GOPROXY="https://goproxy.io,direct"
ARG CGO_ENABLED=0
WORKDIR /go/src/app
COPY . .
RUN go build
RUN ls

FROM alpine:3
COPY --from=build /go/src/app/seedpod /usr/local/bin
EXPOSE 8080
CMD ["/usr/local/bin/seedpod", "serve"]