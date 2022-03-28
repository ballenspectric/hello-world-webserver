FROM golang:latest AS builder
LABEL maintainer="foss@spectric.com"
COPY hello.go .
RUN go build -ldflags="-w -s" -a -tags netgo -o hello hello.go

FROM scratch
COPY --from=builder /go/hello /hello
EXPOSE 8080
CMD ["/hello"]
