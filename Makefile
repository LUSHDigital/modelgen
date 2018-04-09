OLD_SHA:=$(shell shasum -a 256 main-packr.go | cut -d' ' -f1)
NEW_SHA= $(shell shasum -a 256 main-packr.go | cut -d' ' -f1)

all: test install post
test:
	go test -v ./...
	go build
	docker-compose --no-ansi -f docker-compose.yml up -d --force-recreate
	sleep 10
	./modelgen -c root:@localhost:3307 -d modelgen_tests -p models generate
	golint -set_exit_status generated_models
	rm -rf modelgen
	rm -rf ./generated_models
clean:
	docker stop modelgen-tests
	docker rm modelgen-tests
install:
	packr && go install
post:
	@if [ "$(NEW_SHA)" != "$(OLD_SHA)" ]; then\
        echo "sha comparison failed on main-packr.go";\
		exit 1;\
    fi
