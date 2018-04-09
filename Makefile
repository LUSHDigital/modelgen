OLD_SHA:=$(shell shasum -a 256 main-packr.go | cut -d' ' -f1)
NEW_SHA= $(shell shasum -a 256 main-packr.go | cut -d' ' -f1)

all: test install post
test:
	docker-compose --no-ansi -f docker-compose.yml up -d --force-recreate
	while ! docker exec -i modelgen-tests mysql -uroot <<< "select true" | grep TRUE; do sleep 1; done
	modelgen -c root:@localhost:3306 -d modelgen_tests -p models generate
	golint -set_exit_status generated_models
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
