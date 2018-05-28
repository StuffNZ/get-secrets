FROM golang:1.10-alpine as builder

RUN apk upgrade --no-cache --update && \
    apk add --no-cache --update ca-certificates git make

COPY . /app

RUN cd /app && make clean test-default all

FROM alpine

RUN apk upgrade --no-cache --update && \
    apk add --no-cache --update ca-certificates

RUN addgroup app && \
    adduser -D -G app -h /app -s /bin/sh app

COPY --from=builder /app/bin/bitbucket.org/mexisme/get-secrets /app/
RUN chown -R app:app /app && chmod +x /app/get-secrets

# Not sure we want to do this; think we'd just share the dir
# VOLUME /app/secrets

USER app
WORKDIR /app

ENTRYPOINT ["/app/get-secrets"]
