FROM golang:1.19.0-alpine3.15 as build

WORKDIR /go/src/github.com/Ferza17/event-driven-product-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
EXPOSE 3000 3000

RUN go build -o product-service

FROM alpine:3.12.1
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/Ferza17/event-driven-product-service .
CMD [ "./PRODUCT-SERVICE","grpc" ]

