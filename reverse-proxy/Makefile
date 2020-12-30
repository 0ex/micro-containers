

TAG=micro-reverse-proxy
REPO=ghcr.io/0ex/${TAG}
USER=akvadrako

all: build login push

login:
	secret-tool lookup username akvadrako service api.github.com \
	    | docker login ghcr.io -u akvadrako --password-stdin

build:
	CGO_ENABLED=0 GOOS=linux go build -a -o proxy \
	    -tags netgo \
	    -ldflags '-w -extldflags "-static"' 
	docker build -t ${TAG} .

push:
	docker tag ${TAG} ${REPO}
	docker push ${REPO}
