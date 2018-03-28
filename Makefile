OLD_SHA:=$(shell shasum -a 256 main-packr.go | cut -d' ' -f1)
NEW_SHA= $(shell shasum -a 256 main-packr.go | cut -d' ' -f1)

all: test install post
test:
	go test -v -cover ./...
install:
	packr && go install
post:
	@if [ "$(NEW_SHA)" != "$(OLD_SHA)" ]; then\
        echo "sha comparison failed on main-packr.go";\
		exit 1;\
    fi
