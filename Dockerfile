FROM alpine:3.22

WORKDIR /app

COPY bench-exporter .

ENTRYPOINT ["./bench-exporter"]
