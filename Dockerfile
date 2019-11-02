FROM golang:1.11 as builder
RUN git clone https://github.com/joeydouglas/go-container-example.git && \
    cd /go/go-container-example
WORKDIR /go/go-container-example
# COPY . /go/src/github.com/joeydouglas/lenslocked.com/
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -o go-container-example

FROM ubuntu:18.04
WORKDIR /root
COPY --from=builder /go/go-container-example .
CMD ["./go-container-example"]