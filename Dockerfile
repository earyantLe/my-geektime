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


FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/my-geektime /usr/bin/my-geektime

EXPOSE 8090

ENTRYPOINT ["my-geektime"]