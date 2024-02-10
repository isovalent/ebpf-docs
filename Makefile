# Obtain an absolute path to the directory of the Makefile.
# Assume the Makefile is in the root of the repository.
REPODIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
UIDGID := $(shell stat -c '%u:%g' ${REPODIR})

# Prefer podman if installed, otherwise use docker.
# Note: Setting the var at runtime will always override.
CONTAINER_ENGINE ?= $(if $(shell command -v podman), podman, docker)
CONTAINER_RUN_ARGS ?= $(if $(filter ${CONTAINER_ENGINE}, podman), --log-driver=none, --user "${UIDGID}")

IMAGE := ebpf-mkdocs
VERSION := latest
PROD := false
GH_TOKEN := ""

.DEFAULT_TARGET = build-container

.PHONY: build-container
build-container:
	${CONTAINER_ENGINE} build -f ${REPODIR}/tools/Dockerfile \
	--build-arg="UID=$$(id -u $${USER})" --build-arg="GID=$$(id -g $${USER})" -t ${IMAGE}:${VERSION} ${REPODIR}

.PHONY: container-shell
container-shell: build-container
	${CONTAINER_ENGINE} run --rm -it -v "${REPODIR}:/docs" -e "GH_TOKEN=${GH_TOKEN}" -u $$(id -u $${USER}):$$(id -g $${USER}) -w /docs "${IMAGE}:${VERSION}"

.PHONY: html
html: build-container
	${CONTAINER_ENGINE} run --rm -it -v "${REPODIR}:/docs" \
	-e "PROD=${PROD}" -e "GH_TOKEN=${GH_TOKEN}" \
	-w /docs -u $$(id -u $${USER}):$$(id -g $${USER}) --entrypoint "bash" "${IMAGE}:${VERSION}" -c "mkdocs build -d /docs/out"

.PHONY: clear-html
clear-html:
	rm -r out/*

.PHONY: serve
serve: build-container
	${CONTAINER_ENGINE} run --rm -it -p 8000:8000 -v "${REPODIR}:/docs" \
	-e "PROD=${PROD}" -e "GH_TOKEN=${GH_TOKEN}" \
	-w /docs -u $$(id -u $${USER}):$$(id -g $${USER}) --entrypoint "bash" "${IMAGE}:${VERSION}" -c "mkdocs serve -a 0.0.0.0:8000 --watch /docs/docs"

.PHONY: generate-docs
generate-docs:
	cd ${REPODIR}/tools/helper-ref-gen; go run main.go --project-root "${REPODIR}"
	cd ${REPODIR}/tools/feature-tag-gen; go run main.go --project-root "${REPODIR}"
	cd ${REPODIR}/tools/helper-def-scraper; go run main.go --helper-path "${REPODIR}/docs/linux/helper-function"
