FROM golang:alpine AS builder
WORKDIR /app
ADD ./ /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags musl -a -o graphite-sender .

FROM alpine
ENV TARGET="localhost:2003" PREFIX="" FREQUENCY="30s"
RUN addgroup -S graphite && adduser -S -G graphite graphite
COPY --from=builder /app/graphite-sender /graphite-sender
USER graphite
LABEL maintainer="Alejandro Galue <agalue@opennms.org>" \
      name="Graphite Sample Generator"
ENTRYPOINT /graphite-sender -target "$TARGET" -prefix "$PREFIX" -frequency "$FREQUENCY"
