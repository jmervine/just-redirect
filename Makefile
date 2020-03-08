default: test build

test:
	go test -v -race ./...

build:
	docker build -t jmervine/just-redirect:latest .
	docker tag jmervine/just-redirect:latest jmervine/just-redirect:$(shell git reflog | head -n1 | cut -d' ' -f1)

push:
	docker push jmervine/just-redirect:latest
	docker push jmervine/just-redirect:$(shell git reflog | head -n1 | cut -d' ' -f1)

