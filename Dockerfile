FROM alpine:3.22

WORKDIR /app

COPY linux/amd64/bench-exporter .

ENTRYPOINT ["./bench-exporter"]
