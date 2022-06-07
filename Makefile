ifneq (,$(wildcard ./.env))
    include .env
    export
endif

SCAFFOLDER_RUNNER_IMAGE=registry.gitlab.com/2ndwatch/docker-images/ms-scaffolder-runner/ms-scaffolder-runner:1
ARGS=

dockerlogin:
	docker login registry.gitlab.com

execrunner: dockerlogin
	docker run --rm -it -v "${CURDIR}":/app ${SCAFFOLDER_RUNNER_IMAGE} bash

setup: dockerlogin
	docker run --rm -it -v "${CURDIR}":/app ${SCAFFOLDER_RUNNER_IMAGE} bash -c 'cd /app && ./scripts/ms-scaffolder/setup.sh $(ARGS)'

update: dockerlogin
ifdef BUILD_IMAGE
	git clone --branch init git@gitlab.com:2ndwatch/microservices/templates/ms-scaffolder.git "${CURDIR}/.ms-scaffolder"
	docker run --rm -it -v "${CURDIR}":/app ${SCAFFOLDER_RUNNER_IMAGE} bash -c \
	'cd /app && \
	cp ./scripts/ms-scaffolder/update.sh ./.ms-scaffolder/update-temp.sh && \
	./.ms-scaffolder/update-temp.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

protogen: dockerlogin
ifdef BUILD_IMAGE
	docker run --rm -it -v "${CURDIR}":/app ${SCAFFOLDER_RUNNER_IMAGE} bash -c 'cd /app && ./scripts/ms-scaffolder/protogen.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

build: protogen
ifdef BUILD_IMAGE
	docker run --rm -it -v "${CURDIR}":/app ${BUILD_IMAGE} bash -c 'cd /app && ./scripts/ms-scaffolder/build.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

test: build
ifdef BUILD_IMAGE
	docker run --rm -it -v "${CURDIR}":/app ${BUILD_IMAGE} bash -c 'cd /app && ./scripts/ms-scaffolder/test.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

run: build
ifdef BUILD_IMAGE
	docker run --rm -it -v "${CURDIR}":/app ${BUILD_IMAGE} bash -c 'cd /app && ./scripts/ms-scaffolder/run.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

.PHONY: all dockerlogin execrunner setup update protogen build test run 
