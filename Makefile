ifneq (,$(wildcard ./.env))
	include .env
	export
endif

ifeq ($(OS),Windows_NT)
	CURRENT_UID := 0
	CURRENT_GID := 0
else
	CURRENT_UID := $(shell id -u)
	CURRENT_GID := $(shell id -g)
endif

GIT_USER_NAME := $(shell git config --get user.name)
GIT_USER_EMAIL := $(shell git config --get user.email)

export CURRENT_UID
export CURRENT_GID
export GIT_USER_NAME
export GIT_USER_EMAIL

ARGS=
DOCKER_FLAGS=-it
SCAFFOLDER_RUNNER_IMAGE=registry.gitlab.com/2ndwatch/docker-images/ms-scaffolder-runner/ms-scaffolder-runner:1

build: dockerinit
ifdef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v ${VOLUME_PATH} \
	-v "${CURDIR}":/service \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${BUILD_IMAGE} \
	bash -c \
		'cd /service && \
		./scripts/ms-scaffolder/build.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

dockerbuild: build
ifdef BUILD_IMAGE
	docker build -t ${PROJECT_SLUG} "${CURDIR}"
else
	$(error This command cannot be used until after running `make setup`)
endif

dockerbuilddebug: build
ifdef DEBUG_IMAGE
	docker build -t ${PROJECT_SLUG} "${CURDIR}" --build-arg BASE_IMAGE=${DEBUG_IMAGE}
else
	$(error This command can only be used if 'DEBUG_IMAGE' is specified in .env)
endif

dockerinit:
	docker login registry.gitlab.com
	docker pull ${SCAFFOLDER_RUNNER_IMAGE} || echo "Continue on pull fail incase we are testing a local image"
ifdef BUILD_IMAGE
	docker pull ${BUILD_IMAGE} || echo "Continue on pull fail incase we are testing a local image"
endif

protogen: dockerinit
ifdef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v ${VOLUME_PATH} \
	-v "${CURDIR}":/service \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${BUILD_IMAGE} \
	bash -c \
		'cd /service && \
		./scripts/ms-scaffolder/protogen.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

run: dockerbuild
ifdef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v "${HOME}/.aws":/root/.aws \
	${PROJECT_SLUG}
else
	$(error This command cannot be used until after running `make setup`)
endif

setup: dockerinit
ifndef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v "${CURDIR}":/service \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${SCAFFOLDER_RUNNER_IMAGE} \
	bash -c \
		'git config --global user.name "${GIT_USER_NAME}" && \
		git config --global user.email "${GIT_USER_EMAIL}" && \
		cd /service && \
		mkdir .ms-scaffolder && \
		cp ./scripts/ms-scaffolder/setup.sh ./.ms-scaffolder/setup-temp.sh && \
		./.ms-scaffolder/setup-temp.sh $(ARGS)'
else
	$(error You have already set up this project.)
endif

shbuilder: dockerinit
ifdef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v ${VOLUME_PATH} \
	-v "${CURDIR}":/service \
	-v "${HOME}/.aws":/root/.aws:ro \
	-v "${HOME}/.gitconfig":/etc/gitconfig:ro \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${BUILD_IMAGE} \
	bash
else
	$(error This command cannot be used until after running `make setup`)
endif

shrunner: dockerinit
	docker run --rm $(DOCKER_FLAGS) \
	-v "${CURDIR}":/service \
	-v "${HOME}/.aws":/root/.aws:ro \
	-v "${HOME}/.gitconfig":/etc/gitconfig:ro \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${SCAFFOLDER_RUNNER_IMAGE} \
	bash

shservice: dockerbuilddebug
ifdef BUILD_IMAGE
	docker run --rm --entrypoint=sh $(DOCKER_FLAGS) \
	-v "${HOME}/.aws":/root/.aws:ro \
	-v "${HOME}/.gitconfig":/etc/gitconfig:ro \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${PROJECT_SLUG}
else
	$(error This command cannot be used until after running `make setup`)
endif

test: dockerinit
ifdef BUILD_IMAGE
	docker run --rm $(DOCKER_FLAGS) \
	-v ${VOLUME_PATH} \
	-v "${CURDIR}":/service \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${BUILD_IMAGE} \
	bash -c \
		'cd /service && \
		./scripts/ms-scaffolder/test.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

update: dockerinit
ifdef BUILD_IMAGE
	git clone git@gitlab.com:2ndwatch/microservices/templates/ms-scaffolder.git "${CURDIR}/.ms-scaffolder"
	docker run --rm $(DOCKER_FLAGS) \
	-v "${CURDIR}":/service \
	-v "${HOME}/.gitconfig":/etc/gitconfig \
	-u ${CURRENT_UID}:${CURRENT_GID} \
	${SCAFFOLDER_RUNNER_IMAGE} \
	bash -c \
		'cd /service && \
		mkdir .ms-scaffolder && \
		cp ./scripts/ms-scaffolder/update.sh ./.ms-scaffolder/update-temp.sh && \
		./.ms-scaffolder/update-temp.sh $(ARGS)'
else
	$(error This command cannot be used until after running `make setup`)
endif

.PHONY: build dockerbuild dockerbuilddebug dockerinit protogen run setup shbuilder shrunner shservice test update