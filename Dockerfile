# You should override $PARENT at build-time to name the upper-level container
#     e.g. node:7-alpine
# You may need to override $DOCKER_BASE if you're using this repo as a Submodule of another builder repo
# Override $CONFIG_FILE to use a different config. file

# Note: This will only work for recent versions of Docker
# Note: Your $PARENT base OS/Distribution (Debian or Alpine) must be compatible with the Golang builder base.
#      A Golang binary built with Debian won't usually work on Alpine out-of-the-box, for example

# Debian-based builder:
#ARG PARENT_BUILD=golang:1.10
#ARG PARENT=debian

# Alpine-based builder:
ARG PARENT_BUILD=golang:1.11-alpine
ARG PARENT=alpine

##########

FROM $PARENT_BUILD as builder

ARG DOCKER_BASE=.

RUN if [ -f /etc/debian_version ]; then \
      apt-get update && apt-get upgrade -y && \
      apt-get install -y git make; \
    \
    elif [ -f /etc/alpine-release ]; then \
      apk upgrade --no-cache --update && \
      apk add --no-cache --update ca-certificates git make build-base; \
    fi

COPY $DOCKER_BASE/ /app

RUN cd /app && make clean test-default all

##########

FROM $PARENT

ARG DOCKER_BASE=.
ARG CONFIG_FILE=$DOCKER_BASE/get-secrets.toml

# Note: These will need to be overridden in later stages
# Note: SECRETS_EVAL_DOTENV is the default;  the `.env` files are read and copied into the env vars
#       before running the subsequent command
ENV SECRETS_BASE=/app/secrets \
    SECRETS_S3_DOTENV_PATH=""

RUN if [ -f /etc/debian_version ]; then \
      apt-get update && apt-get upgrade -y && \
      apt-get install -y ca-certificates && \
      addgroup app && \
      adduser --disabled-password --ingroup app --home /app --shell /bin/sh app; \
      \
    elif [ -f /etc/alpine-release ]; then \
      apk upgrade --no-cache --update && \
      apk add --no-cache --update ca-certificates && \
      addgroup app && \
      adduser -D -G app -h /app -s /bin/sh app; \
      \
    elif [ -f /etc/centos-release ]; then \
      yum -y update && yum clean all && \
      yum -y update ca-certificates && yum clean all && \
      groupadd app && \
      adduser -g app -d /app -s /bin/sh app; \
    fi

RUN mkdir $SECRETS_BASE && \
    chown app:app $SECRETS_BASE

# We have `NO_SUCH_FILE` to make sure the var is set:
COPY ${CONFIG_FILE:-MISSING_CONFIG_FILE} /app/.secrets.toml
COPY --from=builder /app/bin/get-secrets /app/
RUN chmod ugo=r /app/.secrets.toml && chmod ugo=rx /app/get-secrets

# Not sure we want to do this; think we'd just share the dir:
# VOLUME /app/secrets

USER app
WORKDIR /app

ENTRYPOINT ["/app/get-secrets"]
