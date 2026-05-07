FROM alpine:3.22

WORKDIR /app

COPY linux/${TARGETARCH}/bench-exporter .

ENTRYPOINT ["./bench-exporter"]
