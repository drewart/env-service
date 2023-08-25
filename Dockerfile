FROM golang:1.20-alpine3.17 as build

WORKDIR /usr/src/app

COPY . .
RUN go build -v -o service service.go

FROM alpine:3.17

COPY --from=build /usr/src/app/service /bin/service

EXPOSE 8080

CMD /bin/service