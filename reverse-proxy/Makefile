

TAG=micro-reverse-proxy
REPO=ghcr.io/github/0ex/${TAG}
USER=akvadrako

all: build login push

login:
	secret-tool lookup username akvadrako service api.github.com \
	    | docker login ghcr.io -u akvadrako --password-stdin

build:
	go build -o proxy
	docker build -t ${TAG} .

push:
	docker tag ${TAG} ${REPO}
	docker push ${REPO}
