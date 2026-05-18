FROM node:20-alpine AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM golang:1.25 AS builder

# ENV GOPROXY="https://goproxy.cn,direct"
ENV CGO_ENABLED=0

WORKDIR /app
COPY . /app

COPY --from=frontend-builder /app/frontend/dist ./web/

RUN make build


FROM jrottenberg/ffmpeg:7.1-ubuntu2404 AS ffmpeg


FROM ubuntu:24.04

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get -yqq update && \
    apt-get install -yq --no-install-recommends python3 pipx && rm -rf /var/lib/apt/lists/*

ENV PATH=/root/.local/bin:$PATH
RUN pipx install mkdocs-material --include-deps # --index-url https://pypi.tuna.tsinghua.edu.cn/simple

COPY --from=builder /app/my-geektime /usr/bin/my-geektime
COPY --from=ffmpeg /usr/share/fonts /usr/share/fonts
COPY --from=ffmpeg /usr/share/fontconfig /usr/share/fontconfig
COPY --from=ffmpeg /usr/bin/fc-* /usr/bin/
COPY --from=ffmpeg /usr/local /usr/local/
ENV LD_LIBRARY_PATH=/usr/local/lib:/usr/local/lib64

EXPOSE 8090

ENTRYPOINT ["my-geektime"]