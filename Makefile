OLD_SHA:=$(shell shasum -a 256 a_main-packr.go | cut -d' ' -f1)
NEW_SHA= $(shell shasum -a 256 a_main-packr.go | cut -d' ' -f1)

all: test install post
test:
	packr
	go test -v -count 1 ./...
	go build
	docker-compose --no-ansi -f docker-compose.yml up -d --force-recreate
	sleep 5
	./modelgen -c root:@localhost:3307 -d modelgen_tests -p models generate
	golint -set_exit_status generated_models
	rm -rf modelgen
	rm -rf ./generated_models
test-ci:
	go test -v -count 1 ./...
	go build
	docker-compose --no-ansi -f docker-compose.yml up -d --force-recreate
	sleep 30 # annoying, but for ci.
	./modelgen -c root:@localhost:3307 -d modelgen_tests -p models generate
	golint -set_exit_status generated_models
	rm -rf modelgen
	rm -rf ./generated_models
clean:
	docker rm -f modelgen-tests
install:
	packr && go install
post:
	@if [ "$(NEW_SHA)" != "$(OLD_SHA)" ]; then\
        echo "sha comparison failed on a_main-packr.go";\
		exit 1;\
    fi
