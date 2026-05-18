PROJECT:=my-geektime

.PHONY: build web githook run website image all


all: build

build:
	go build -ldflags "-X main.buildTime=`date +%Y%m%d.%H:%M:%S` -X main.buildCommit=`git rev-parse --short=12 HEAD` -X main.buildBranch=`git branch --show-current`"

web:
	cd frontend && npm install && npm run build && rm -rf ../web/index.html && rm -rf ../web/assets && cp -r dist/* ../web/	

githook:
	git config core.hooksPath .githooks
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	chmod 777 .githooks/commit-msg
	chmod 777 .githooks/pre-commit
	chmod 777 .githooks/pre-push

run: build
	gofmt -w ./
	my-geektime \
    --help

website:
	pip3 install mkdocs-material
	mkdocs gh-deploy --force --no-history

image:
	docker buildx build --platform linux/amd64,linux/arm64 -t zkep/mygeektime:latest --push .

