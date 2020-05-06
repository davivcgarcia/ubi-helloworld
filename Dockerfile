FROM registry.access.redhat.com/ubi8/go-toolset AS builder
ENV GOPATH /go
WORKDIR /go/src/github.com/davivcgarcia/ubi-helloworld
COPY app.go .
USER root
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /opt/
COPY --from=builder /go/src/github.com/davivcgarcia/ubi-helloworld/app .
EXPOSE 8080
USER 1001
CMD ["/opt/app"] 