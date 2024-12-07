# Obtain an absolute path to the directory of the Makefile.
# Assume the Makefile is in the root of the repository.
REPODIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
UIDGID := $(shell stat -c '%u:%g' ${REPODIR})

# Prefer podman if installed, otherwise use docker.
# Note: Setting the var at runtime will always override.
CONTAINER_ENGINE ?= $(if $(shell command -v podman), podman, docker)
CONTAINER_RUN_ARGS ?= $(if $(filter ${CONTAINER_ENGINE}, podman), --log-driver=none, --user "${UIDGID}")

IMAGE := ghcr.io/isovalent/ebpf-docs
VERSION := $(shell cat ${REPODIR}/tools/image-version)

PROD := false
GH_TOKEN := ""

.DEFAULT_TARGET := serve
default: serve ;

.PHONY: build-container
build-container: 
	$(eval EPOCH := $(shell date +%s))
	${CONTAINER_ENGINE} build -f ${REPODIR}/tools/Dockerfile -t ${IMAGE}:$(EPOCH) ${REPODIR}
	echo $(EPOCH) > ${REPODIR}/tools/image-version

.PHONY: push-container
push-container:
	${CONTAINER_ENGINE} push ${IMAGE}:${VERSION}

.PHONY: container-shell
container-shell:
	${CONTAINER_ENGINE} run --rm -it -v "${REPODIR}:/docs" -e "GH_TOKEN=${GH_TOKEN}" \
	-e "AS_USER=$$(id -u $${USER})" -e "AS_GROUP=$$(id -g $${USER})" \
	-w /docs "${IMAGE}:${VERSION}"

.PHONY: html
html:
	${CONTAINER_ENGINE} run --rm -it -v "${REPODIR}:/docs" \
	-e "PROD=${PROD}" -e "GH_TOKEN=${GH_TOKEN}" \
	-e "AS_USER=$$(id -u $${USER})" -e "AS_GROUP=$$(id -g $${USER})" \
	-w /docs "${IMAGE}:${VERSION}" "mkdocs build -d /docs/out"

.PHONY: clear-html
clear-html:
	rm -r out/*

.PHONY: serve
serve:
	${CONTAINER_ENGINE} run --rm -it -p 8000:8000 -v "${REPODIR}:/docs" \
	-e "PROD=${PROD}" -e "GH_TOKEN=${GH_TOKEN}" \
	-w /docs -e "AS_USER=$$(id -u $${USER})" -e "AS_GROUP=$$(id -g $${USER})" \
	"${IMAGE}:${VERSION}" "mkdocs serve -a 0.0.0.0:8000 --watch /docs/docs"

.PHONY: build-spellcheck
build-spellcheck:
	${CONTAINER_ENGINE} run --rm -v "${REPODIR}:/docs" -w /docs golang:latest bash -c \
	"CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/spellcheck /docs/tools/spellcheck/."

.PHONY: build-gen-tools
build-gen-tools:
	${CONTAINER_ENGINE} run --rm -v "${REPODIR}:/docs" -w /docs golang:latest bash -c \
	"CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/libbpf-tag-gen /docs/tools/libbpf-tag-gen/. && \
	CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/helper-ref-gen /docs/tools/helper-ref-gen/. && \
	CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/feature-gen /docs/tools/feature-gen/. && \
	CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/kfunc-gen /docs/tools/kfunc-gen/. && \
	CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/mtu-calc /docs/tools/mtu-calc/. && \
	CGO_ENABLED=0 go build -buildvcs=false -o /docs/tools/bin/helper-def-scraper /docs/tools/helper-def-scraper/."

LIBBPF_REF := $(shell cat ${REPODIR}/tools/libbpf-ref)

.PHONY: generate-docs
generate-docs: build-gen-tools
	${CONTAINER_ENGINE} run --rm -v "${REPODIR}:/docs" \
		-w /docs -e "AS_USER=$$(id -u $${USER})" -e "AS_GROUP=$$(id -g $${USER})" "${IMAGE}:${VERSION}" \
		"/docs/tools/bin/helper-ref-gen --project-root /docs && \
		/docs/tools/bin/libbpf-tag-gen --project-root /docs --libbpf-ref '${LIBBPF_REF}' && \
		/docs/tools/bin/feature-gen --project-root /docs --tags && \
		/docs/tools/bin/feature-gen --project-root /docs --timeline && \
		/docs/tools/bin/kfunc-gen --project-root /docs && \
		/docs/tools/bin/mtu-calc --project-root /docs && \
		/docs/tools/bin/helper-def-scraper --helper-path /docs/docs/linux/helper-function --libbpf-ref '${LIBBPF_REF}'"

.PHONY: update-libbpf-ref
update-libbpf-ref:
	curl -L \
		-H "Accept: application/vnd.github+json" \
		-H "X-GitHub-Api-Version: 2022-11-28" \
		https://api.github.com/repos/libbpf/libbpf/commits | \
	jq -r '.[0].sha' > ${REPODIR}/tools/libbpf-ref

.PHONY: spellcheck
spellcheck: build-spellcheck html
	${CONTAINER_ENGINE} run --rm -v "${REPODIR}:/docs" \
		-w /docs -e "AS_USER=$$(id -u $${USER})" -e "AS_GROUP=$$(id -g $${USER})" "${IMAGE}:${VERSION}" \
		"/docs/tools/bin/spellcheck --project-root /docs"
