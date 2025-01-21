
JQ_TARGETS := build test test-integration devnet-up devnet-down devnet-clean devnet-logs alphanet-up alphanet-down alphanet-clean alphanet-logs indexer-up

# Requires at least Python v3.9; specify a minor version below if needed
PYTHON?=python3

$(JQ_TARGETS): check-jq

.PHONY: $(JQ_TARGETS) check-jq

check-jq:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is required but it's not installed. Aborting."; exit 1; }


# ---- basics ------
nuke: devnet-clean
	git clean -Xdf
.PHONY: nuke


build:
	earthly --secret GITHUB_TOKEN +go-build
.PHONY: build

test:
	earthly --secret GITHUB_TOKEN +go-test-coverage-unit
.PHONY: test

test-contracts:
	earthly --secret GITHUB_TOKEN +yarn-test
.PHONY: test

fuzz:
	earthly --secret GITHUB_TOKEN +go-test-fuzz
.PHONY: fuzz

test-integration:
	earthly -P --secret GITHUB_TOKEN +go-test-integration-zr-blockexplorer-api
	earthly -P --secret GITHUB_TOKEN +go-test-integration-zr-proof-orchestrator
	earthly -P --secret GITHUB_TOKEN +go-test-e2e-zr-proof-orchestrator
.PHONY: test-integration

lint:
	earthly --secret GITHUB_TOKEN +go-lint-govulncheck
	earthly --secret GITHUB_TOKEN +go-lint-staticcheck
	earthly --secret GITHUB_TOKEN +go-lint-revive
	earthly --secret GITHUB_TOKEN +go-lint-mirror
	earthly --secret GITHUB_TOKEN +go-lint-golangci
.PHONY: test

# ---- devnet ------

devnet-up: submodules
	# if DOCKER_REGISTRY_URL is omitted, the mock binaries will be used
	[ -f .arg ] && echo using .arg file || cp .arg.devnet .arg
	earthly --secret GITHUB_TOKEN +devnet-up
.PHONY: devnet-up

devnet-pectra-up: submodules
	# if DOCKER_REGISTRY_URL is omitted, the mock binaries will be used
	[ -f .arg ] && echo using .arg file || cp .arg.pectra.devnet .arg
	earthly --secret GITHUB_TOKEN +devnet-up
.PHONY: devnet-pectra-up

devnet-down:
	earthly --secret GITHUB_TOKEN +devnet-down
.PHONY: devnet-down

devnet-start:
	earthly --secret GITHUB_TOKEN +devnet-start-l2
.PHONY: devnet-start

devnet-stop:
	earthly --secret GITHUB_TOKEN +devnet-stop
.PHONY: devnet-stop

devnet-start-replica:
	earthly --secret GITHUB_TOKEN +devnet-start-replica
.PHONY: devnet-start-replica

devnet-stop-replica:
	earthly --secret GITHUB_TOKEN +devnet-stop-replica
.PHONY: devnet-stop-replica

devnet-clean:
	earthly --secret GITHUB_TOKEN +devnet-clean
.PHONY: devnet-clean

devnet-logs:
	@(cd ./ops-bedrock && docker-compose logs -f)
.PHONY: devnet-logs

op-bindings:
	earthly --secret GITHUB_TOKEN +go-local-update-bindings
.PHONY: op-bindings

indexer-up:
	earthly --secret GITHUB_TOKEN +indexer-up
.PHONY: indexer-up

build-blockexplorer-docs:
	earthly +build-blockexplorer-docs
.PHONY: build-blockexplorer-docs

build-zr-eth-analytics-docs:
	earthly +build-zr-eth-analytics-docs
.PHONY: build-zr-eth-analytics-docs

update-op-geth:
	./ops/scripts/update-op-geth.py
.PHONY: update-op-geth

submodules:
	git submodule update --init --recursive
.PHONY: submodules
