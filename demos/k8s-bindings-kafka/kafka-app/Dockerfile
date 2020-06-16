FROM golang as build-stage
WORKDIR /app

COPY app.go /

RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kafka-app

FROM alpine
COPY --from=build-stage /kafka-app /
CMD ["/kafka-app"]