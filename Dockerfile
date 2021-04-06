FROM golang:1.16-alpine
LABEL maintainer="Jintae, Kim <6199@outlook.kr>"

COPY . /app
ENV HOME=/app

# Build Argument Set
ARG ELS_HOST=${ELS_HOST}
ARG ELS_USER=${ELS_USER}
ARG ELS_PASSWORD=${ELS_PASSWORD}
ARG ELS_INDEX=${ELS_INDEX}

# Env Set
ENV GIN_MODE=release
ENV PORT=8000
ENV ELS_HOST=${ELS_HOST}
ENV ELS_USER=${ELS_USER}
ENV ELS_PASSWORD=${ELS_PASSWORD}
ENV ELS_INDEX=${ELS_INDEX}

# Timezone Set
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

# Build
WORKDIR ${HOME}
RUN apk --no-cache add tzdata && go build main.go wire_gen.go

EXPOSE $PORT
ENTRYPOINT ["./main"]
