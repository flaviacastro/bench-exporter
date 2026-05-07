FROM alpine:3.22

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY ${TARGETOS}/${TARGETARCH}/bench-exporter .

ENTRYPOINT ["./bench-exporter"]
