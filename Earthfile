# Earthfile
VERSION 0.8

# version of go compiler to use
# remember to update both SHA versions as well
# when updating major version, also be sure to look at the .github/worflows files
ARG --global GO_VERSION=1.23.1
ARG --global GO_SHA_AMD64=49bbb517cfa9eee677e1e7897f7cf9cfdbcf49e05f61984a2789136de359f9bd
ARG --global GO_SHA_ARM64=faec7f7f8ae53fda0f3d408f52182d942cc89ef5b7d3d9f23ff117437d4b2d2f

# ------------------------- Base Images -----------------------------

# The final image on which we will deploy based on wolfi-base
docker-base:
    FROM cgr.dev/chainguard/wolfi-base
    RUN apk add --no-cache libgcc ca-certificates wget jq opencl
    RUN update-ca-certificates
    ENV LD_LIBRARY_PATH=/usr/lib/

INSTALL_GO_CMD:
    FUNCTION
    IF [ "$(uname -m)" = "x86_64" ]
        ENV GO_TAR="go${GO_VERSION}.linux-amd64.tar.gz"
        ENV GO_URL="https://golang.org/dl/${GO_TAR}"
        ENV GO_SHA="${GO_SHA_AMD64}"
    ELSE
        ENV GO_TAR="go${GO_VERSION}.linux-arm64.tar.gz"
        ENV GO_URL="https://golang.org/dl/${GO_TAR}"
        ENV GO_SHA="${GO_SHA_ARM64}"
    END
    WORKDIR /Downloads
    RUN wget -nv "${GO_URL}"
    RUN echo "${GO_SHA} ${GO_TAR}" | sha256sum -c
    RUN tar -C /usr/ -xzf "${GO_TAR}"
    RUN rm "${GO_TAR}"
    IF [ -d "/usr/go/bin" ]
        ENV PATH=$PATH:/usr/go/bin
        ENV GOBIN=/usr/bin
    ELSE
        ENV PATH=$PATH:/usr/local/go/bin
        ENV GOBIN=/usr/local/go/bin
    END
    RUN go version
    RUN which go

# Add tools to the deploy image to use as our go builder
go-builder:
    FROM +docker-base
    RUN apk add --no-cache git build-base
    DO +INSTALL_GO_CMD

# use the streamlined docker-in-docker image for dockerized e2e testing
e2e-docker:
    FROM earthly/dind:alpine
    RUN apk add --no-cache git
    DO +INSTALL_GO_CMD

# Getting Node 16 and the right glibc for foundry together is a PITA with newer base images
# Hence the need for an entirely new builder based on debian 11
yarn-base:
    FROM debian:11
    RUN apt-get update
    RUN apt-get install -yq --no-install-recommends \
        bash \
        build-essential \
        ca-certificates \
        curl \
        gnupg2 \
        git \
        wget \
        lcov \
        jq \
        python3
    RUN mkdir -p /etc/apt/keyrings
    RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
    RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_18.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
    RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
    RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list
    RUN apt-get update && apt-get install -yq --no-install-recommends nodejs yarn
    RUN node -v
    RUN curl -L https://foundry.paradigm.xyz | bash
    ENV PATH="${PATH}:/root/.foundry/bin"
    COPY versions.json /src/
    COPY ops/scripts/install-foundry.sh /src/ops/scripts/
    RUN /src/ops/scripts/install-foundry.sh
    DO +INSTALL_GO_CMD

# ------------------------- Intermediate Images -----------------------------

## --- Go-Builder Based Intermediate Images

# Download the go deps in an earlier stage as they seldom change
GO_BUILD_DEPENDENCIES_CMD:
    FUNCTION
    WORKDIR /src
    COPY ./go.mod /src
    COPY ./go.sum /src
    RUN go mod download
    RUN go mod verify
    RUN go install github.com/mfridman/tparse@latest
    RUN rm -rf ~/.gitconfig # Don't save the token in the image

# NOTE: best practice is normally to copy *ONLY* the required source files since
# every change here changes the input to the next stage and can break caching.
# However, most of our code here is highly interconnected and it's easier to copy it all
GO_COPY_CODE_CMD:
    FUNCTION
    COPY ./common /src/common
    COPY ./databases /src/databases
    COPY ./op-batcher /src/op-batcher
    COPY ./op-bindings /src/op-bindings
    COPY ./op-chain-ops /src/op-chain-ops
    COPY ./op-e2e /src/op-e2e
    COPY ./op-node /src/op-node
    COPY ./op-service /src/op-service
    COPY ./zr-proof-orchestrator /src/zr-proof-orchestrator

go-build-copy-source:
    FROM +go-builder
    DO +GO_BUILD_DEPENDENCIES_CMD
    DO +GO_COPY_CODE_CMD

## --- Docker-in-Docker Based Intermediate Images

e2e-docker-copy-source:
    FROM +e2e-docker
    DO +GO_BUILD_DEPENDENCIES_CMD
    DO +GO_COPY_CODE_CMD
    COPY ./databases /src/databases

## --- Yarn Based Intermediate Images

YARN_BUILD_DEPENDENCIES_CMD:
    FUNCTION
    COPY ./.nvmrc /src/.nvmrc
    COPY ./.yarn /src/.yarn
    COPY ./.yarnrc /src/.yarnrc
    COPY ./package.json /src/package.json
    COPY ./tsconfig.json /src/tsconfig.json
    COPY ./yarn.lock /src/yarn.lock
    COPY ./.eslintrc.js /src/.eslintrc.js
    COPY ./.prettierrc.js /src/.prettierrc.js
    COPY ./packages/contracts-bedrock/package.json /src/packages/contracts-bedrock/package.json
    WORKDIR /src/packages/contracts-bedrock
    RUN yarn install --frozen-lockfile

YARN_COPY_CODE_CMD:
    FUNCTION
    COPY ./packages/contracts-bedrock/ /src/packages/contracts-bedrock
    COPY ./common /src/common
    COPY ./op-bindings /src/op-bindings
    COPY ./op-chain-ops /src/op-chain-ops
    COPY ./op-node /src/op-node
    COPY ./op-service /src/op-service
    COPY ./zr-proof-orchestrator /src/zr-proof-orchestrator

yarn-copy-source:
    FROM +yarn-base
    DO +GO_BUILD_DEPENDENCIES_CMD
    DO +YARN_BUILD_DEPENDENCIES_CMD
    DO +YARN_COPY_CODE_CMD
    # also compile the contracts since otherwise all steps that depend on this
    # target will also compile them separately and it usually won't get cached
    WORKDIR /src/packages/contracts-bedrock
    RUN foundryup
    RUN yarn build

# ------------------------- Build Targets -------------------------


# compile ./main.go and create binary for it
GO_BUILD_CMD:
    FUNCTION
    ARG --required servicename
    RUN env CGO_ENABLED=1 go build -v -o "$servicename" ./main.go
    SAVE ARTIFACT "$servicename"

# Build every executable that we need.
go-build:
    FROM +go-build-copy-source
    WORKDIR /src/op-batcher/cmd
    DO +GO_BUILD_CMD --servicename="op-batcher"
    WORKDIR /src/op-node/cmd
    DO +GO_BUILD_CMD --servicename="op-node"

# ------------------------ Docker Targets ------------------------

# Generate the version information (locally so we don't need to copy the .git/)
# this generates a json file to avoid compiling the information with the code
# which allows us to better cache compilation results since now source files
# are unchanged (directly) by a changed git hash
# To use this, add the file to the docker image separate from a compiled binary
# specifically: `github.com/zircuit-labs/zkr-go-common/version` expects to find it in `/etc/version.json`
git-version-info:
    LOCALLY
    RUN bash get_version.sh
    SAVE ARTIFACT version.json

# all of the go services from this repository follow the same pattern
# compile ./main.go and create binary for it
DOCKER_BUILD_CMD:
    FUNCTION
    ARG --required servicename
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    ENV PATH=$PATH:/usr/bin/
    COPY +go-build/$servicename /usr/bin
    COPY +git-version-info/version.json /etc/version.json
    ENTRYPOINT /usr/bin/$servicename
    SAVE IMAGE "$servicename":latest
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        SAVE IMAGE --push "${DOCKER_REGISTRY_URL}/${servicename}:${VERSION_TAG}"
    END

docker-op-batcher:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +docker-base
    DO +DOCKER_BUILD_CMD --servicename="op-batcher" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-op-node:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +docker-base
    DO +DOCKER_BUILD_CMD --servicename="op-node" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

# L2 geth is already built and in the `l2geth-ccc` ECR repo
# This just adds `wget` and the entrypoint file to it
# and optionally pushes that to the `l2-geth` ECR repo
docker-l2-geth:
    FROM ../l2-geth-public+docker-l2geth-no-ccc
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    ENV GETH_MINER_RECOMMIT=2s
    RUN apk add --no-cache jq wget
    COPY ops/entrypoint-l2.sh /entrypoint.sh
    ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
    SAVE IMAGE l2-geth:latest
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        SAVE IMAGE --push "${DOCKER_REGISTRY_URL}/l2-geth:${VERSION_TAG}"
    END

DOCKER_GENESIS_CMD:
    FUNCTION
    ARG --required network
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    COPY +git-version-info/version.json /etc/version.json
    COPY ./packages/contracts-bedrock /src/packages/contracts-bedrock
    RUN echo "$network-sepolia" > /tmp/network_name
    COPY ./ops/entrypoint-genesis.sh /entrypoint.sh
    ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
    SAVE IMAGE ${network}-genesis:latest
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        SAVE IMAGE --push "${DOCKER_REGISTRY_URL}/${network}-genesis:${VERSION_TAG}"
    END

DOCKER_GENESIS_MAINNET_CMD:
    FUNCTION
    ARG --required network
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    COPY +go-build/op-node /usr/bin
    COPY +git-version-info/version.json /etc/version.json
    COPY ./packages/contracts-bedrock /src/packages/contracts-bedrock
    RUN echo "$network" > /tmp/network_name
    COPY ./ops/entrypoint-genesis.sh /entrypoint.sh
    ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
    SAVE IMAGE ${network}-genesis:latest
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        SAVE IMAGE --push "${DOCKER_REGISTRY_URL}/${network}-genesis:${VERSION_TAG}"
    END

docker-alphanet-genesis:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +yarn-copy-source
    DO +DOCKER_GENESIS_CMD --network="alphanet" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-betanet-genesis:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +yarn-copy-source
    DO +DOCKER_GENESIS_CMD --network="betanet" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-betanet2-genesis:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +yarn-copy-source
    DO +DOCKER_GENESIS_CMD --network="betanet2" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-testnet-genesis:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +yarn-copy-source
    DO +DOCKER_GENESIS_CMD --network="testnet" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-mainnet-genesis:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM +yarn-copy-source
    DO +DOCKER_GENESIS_MAINNET_CMD --network="mainnet" --DOCKER_REGISTRY_URL=$DOCKER_REGISTRY_URL --VERSION_TAG=$VERSION_TAG

docker-all:
    BUILD +docker-alphanet-genesis
    BUILD +docker-betanet-genesis
    BUILD +docker-betanet2-genesis
    BUILD +docker-testnet-genesis
    BUILD +docker-mainnet-genesis
    BUILD +docker-l2-geth
    BUILD +docker-op-batcher
    BUILD +docker-op-node

# When actually pushing images to cloud providers,
# do it sequentially to prevent rate limiting and other errors.
# Ideally if the images are already built (with +docker-all)
# this should be exceptionally fast
docker-all-sequential:
    WAIT
        BUILD +docker-betanet-genesis
    END
    WAIT
        BUILD +docker-betanet2-genesis
    END
    WAIT
        BUILD +docker-alphanet-genesis
    END
    WAIT
        BUILD +docker-l2-geth
    END
    WAIT
        BUILD +docker-op-batcher
    END
    WAIT
        BUILD +docker-op-node
    END
    WAIT
        BUILD +docker-testnet-genesis
    END
    WAIT
        BUILD +docker-mainnet-genesis
    END

# ------------------------- Test Targets -------------------------

# execute all unit tests
# skip tests that contain `_Integration_` as those are run separately
# save coverage information as `./.coverage/unit`
go-test-coverage-unit:
    FROM +go-build-copy-source
    WORKDIR /src/
    # the tests use the L1 allocs, deployment addresses and deploy config of devnet
    COPY +devnet-l1-genesis/.devnet/* /src/.devnet/
    COPY +devnet-l1-genesis/deploy-config/* /src/packages/contracts-bedrock/deploy-config/
    # TODO: Add -race flag back to the tests
    RUN go test -json -shuffle=on -coverpkg="$(go list ./... | paste -d, -s -)" -covermode=atomic -coverprofile=/coverage.unit -skip _Integration_ ./... | tparse -all -progress
    SAVE ARTIFACT --force /coverage.unit AS LOCAL ./.coverage/unit.gocov

# merges all coverage data in local `./.coverage` dir and outputs as `./.coverage/combined-coverage`
go-test-coverage-merge:
    FROM +go-build-copy-source
    # genhtml also needs the perl date package
    RUN apk add --no-cache lcov perl-date-format

    RUN go install github.com/wadey/gocovmerge@latest
    RUN go install github.com/jandelgado/gcov2lcov@latest
    COPY ./.coverage /coverage
    # combine just the go coverage (requires following convention to use gocov ending)
    RUN gocovmerge /coverage/*.gocov > /combined-coverage
    # convert it to lcov and filter out unwanted paths via the `--remove` flag
    RUN gcov2lcov -infile=/combined-coverage -outfile=/unfiltered-combined-coverage.lcov
    RUN lcov \
        --remove /unfiltered-combined-coverage.lcov \
        --output-file /combined-coverage.lcov \
        --substitute 's#/src/##g' \
        "*op-bindings*" \
        "*mocks*" "*mock_*.go" \
        "*op-e2e*"
    SAVE ARTIFACT --force /combined-coverage.lcov AS LOCAL ./.coverage/combined-coverage.lcov

# for local use only: generates coverage information using lcov tools
go-test-coverage-local-html:
    FROM +go-test-coverage-merge
    RUN genhtml -o /coveragehtml/ /combined-coverage.lcov --show-details --legend --prefix /src
    SAVE ARTIFACT --force /coveragehtml AS LOCAL ./.coveragehtml_go

# run fuzz tests
go-test-fuzz:
    FROM +go-build-copy-source
    WORKDIR /src/
    # NOTE: These have been taken from the various Makefile targets and placed alphabetical order
    # so that future changes are easier to see in a diff.
    # op-batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelCloseTimeout /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelConfig_CheckTimeout /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzChannelZeroCloseTimeout /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationTimeoutMaxChannelDuration /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationTimeoutZeroMaxChannelDuration /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDurationZero /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzSeqWindowClose /src/op-batcher/batcher
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzSeqWindowZeroTimeoutClose /src/op-batcher/batcher
    # op-chain-ops
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz=FuzzAliasing /src/op-chain-ops/crossdomain
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz=FuzzEncodeDecodeLegacyWithdrawal /src/op-chain-ops/crossdomain
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz=FuzzEncodeDecodeWithdrawal /src/op-chain-ops/crossdomain
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz=FuzzVersionedNonce /src/op-chain-ops/crossdomain
    # op-node
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzBatchRoundTrip /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDecodeDepositTxDataToL1Info /src/op-node/rollup/driver
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDeriveDepositsBadVersion /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDeriveDepositsRoundTrip /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzFrameUnmarshalBinary /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzL1InfoAgainstContract /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzL1InfoBedrockRoundTrip /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzL1InfoEcotoneRoundTrip /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseFrames /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseL1InfoDepositTxDataBadLength /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseL1InfoDepositTxDataValid /src/op-node/rollup/derive
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzRejectCreateBlockBadTimestamp /src/op-node/rollup/driver
    RUN	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzUnmarshallLogEvent /src/op-node/rollup/derive
    # op-service
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDetectNonBijectivity /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzEncodeDecodeBlob /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzEncodeScalar /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadMarshalUnmarshalV1 /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadMarshalUnmarshalV2 /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadMarshalUnmarshalV3 /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadUnmarshal /src/op-service/eth
    RUN go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzOBP01 /src/op-service/eth

go-test-integration-op-e2e:
    FROM +go-build
    WORKDIR /src
    RUN go test -json -shuffle=on -coverpkg="$(go list ./... | paste -d, -s -)" -covermode=atomic -coverprofile=/coverage.integration-op-e2e -run _Integration_Ope2e$ ./... | tparse -all -notests -progress
    SAVE ARTIFACT --force /coverage.integration-op-e2e AS LOCAL ./.coverage/integration-op-e2e.gocov

# ------------------------- Yarn/Node Targets -------------------------


faultmon-builder:
    FROM +yarn-base
    WORKDIR /src
    COPY yarn.lock .nvmrc package.json tsconfig.json /src/
    COPY ./packages/chain-mon /src/packages/chain-mon
    WORKDIR /src/packages/chain-mon
    RUN yarn install --immutable
    RUN npm install -g typescript
    RUN yarn build
    SAVE ARTIFACT /src/node_modules top_node_modules
    SAVE ARTIFACT /src/packages/chain-mon/dist dist
    SAVE ARTIFACT /src/packages/chain-mon/node_modules node_modules
    SAVE ARTIFACT /src/packages/chain-mon/src src
    SAVE ARTIFACT /src/packages/chain-mon/package.json package.json

faultmon-docker:
    ARG DOCKER_REGISTRY_URL
    ARG VERSION_TAG
    FROM cgr.dev/chainguard/wolfi-base
    RUN apk add --no-cache typescript yarn npm
    RUN npm install --global tsx
    COPY +faultmon-builder/top_node_modules ./node_modules
    COPY +faultmon-builder/dist ./app/dist
    COPY +faultmon-builder/node_modules ./app/node_modules
    COPY +faultmon-builder/src ./app/src
    COPY +faultmon-builder/package.json ./app
    WORKDIR ./app/src
    ENTRYPOINT ["yarn", "run", "start:fault-mon"]
    SAVE IMAGE faultmon:latest
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        SAVE IMAGE --push "${DOCKER_REGISTRY_URL}/faultmon:${VERSION_TAG}"
    END


# run the tests and collect coverage info
yarn-coverage:
    FROM +yarn-copy-source
    WORKDIR /src/packages/contracts-bedrock
    RUN --no-cache yarn coverage --report lcov \
        --report-file /unfiltered_coverage.lcov \
        --no-match-path 'test/invariants/*' \
        --no-match-test 'testFuzz*|test_callWithMinGas_noLeakageLow_succeeds|test_callWithMinGas_noLeakageHigh_succeeds'
    RUN lcov --rc lcov_branch_coverage=1 --remove /unfiltered_coverage.lcov --output-file /coverage.lcov "*test*" "*scripts*"
    SAVE ARTIFACT --force /coverage.lcov AS LOCAL ./.coverage/coverage_sol.lcov

# for local use only: generates coverage information using lcov tools
yarn-coverage-local-html:
    FROM +yarn-coverage
    RUN genhtml -o /coveragehtml/ /coverage.lcov --show-details --legend --branch-coverage
    SAVE ARTIFACT --force /coveragehtml AS LOCAL ./.coveragehtml_sol

# run just the tests since coverage won't actually fail if tests fail
yarn-test:
    FROM +yarn-copy-source
    WORKDIR /src/packages/contracts-bedrock
    # also print contract sizes here to fail if any contract exceeds the size limit
    RUN yarn build --sizes
    RUN --no-cache yarn test -vvv

# ------------------------- Linting / Static Analysis -------------------------

# Run govulncheck
go-lint-govulncheck:
    FROM +go-build-copy-source
    RUN go install golang.org/x/vuln/cmd/govulncheck@latest
    WORKDIR /src
    RUN govulncheck ./...

# Run staticcheck - is part of golangci-lint but unfortunately not the same thing
go-lint-staticcheck:
    FROM +go-build-copy-source
    RUN go install honnef.co/go/tools/cmd/staticcheck@latest
    COPY staticcheck.conf /src/staticcheck.conf
    WORKDIR /src
    RUN staticcheck ./...

# Run revive - is part of golangci-lint but unfortunately difficult to configure there
go-lint-revive:
    FROM +go-build-copy-source
    RUN go install github.com/mgechev/revive@latest
    COPY revive.toml /src/revive.toml
    WORKDIR /src
    RUN revive -set_exit_status -formatter stylish -config /src/revive.toml -exclude op-node/... -exclude op-batcher/... -exclude op-e2e/... -exclude op-service/... -exclude op-chain-ops/... ./...

# Mirror is not yet part of other linters
# This looks for redundant conversions between types
go-lint-mirror:
    FROM +go-build-copy-source
    RUN go install github.com/butuzov/mirror/cmd/mirror@latest
    WORKDIR /src
    RUN mirror ./...

lint-e2e-test:
    FROM +yarn-copy-source
    WORKDIR /src/packages/contracts-bedrock
    RUN --no-cache yarn lint:e2e:check
    RUN --no-cache yarn lint:e2e-sls:check

# Run golangci-lint - massive set of checks; highly configurable.
go-lint-golangci:
    FROM +go-build-copy-source
    RUN apk add --no-cache curl
    RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOBIN v1.60.3
    COPY .golangci.toml /src/.golangci.toml # linter settings
    WORKDIR /src
    RUN golangci-lint run --config=.golangci.toml

# ------------------------- Locally Modify Code -------------------------

# Update all dependent code
# WARNING: This can absolutely break the build, so be cautious.
# When this happens, check each and every diff in go.mod
go-local-mod-update:
    LOCALLY
    RUN go get -d -u ./... && go mod tidy

# Install gofumpt locally (if needed),
# and format all the go code nice and neat
go-local-fmt:
    LOCALLY
    IF ! [ -x "$(command -v gofumpt)" ]
        RUN go install mvdan.cc/gofumpt@latest
    END
    RUN gofumpt -l -w .

# Recreates the go bindings from the bedrock packages
go-local-update-bindings:
    FROM +yarn-copy-source
    RUN go install github.com/zircuit-labs/l2-geth-public/cmd/abigen@${L2GETH_COMMIT}
    WORKDIR /src/op-bindings
    RUN make bindings
    SAVE ARTIFACT --force /src/op-bindings/bindings AS LOCAL ./op-bindings/bindings

# Compiles the openapi specification for the Block Explorer V1 API into a single static HTML page using the redoc cli docker image.
# The generated HTML page is standalone and doesn't require any external dependencies, this page is embedded into the
# block explorer API binary and is served at the GET /v1/docs endpoint.
build-blockexplorer-docs:
    LOCALLY
    RUN docker run --rm -v ./zr-blockexplorer-api:/data ghcr.io/redocly/redoc/cli:latest build /data/docs/api_v1.yaml -o /data/internal/api/v1/docs.html

build-zr-eth-analytics-docs:
    LOCALLY
    RUN docker run --rm -v ./zr-eth-analytics:/data ghcr.io/redocly/redoc/cli:latest build /data/docs/api.yaml -o /data/internal/api/handlers/docs.html

# update and download all submodules
git-update-submodules:
    LOCALLY
    RUN git submodule sync --recursive
    RUN git -c protocol.version=2 submodule update --init --force --depth=1 --recursive

# ------------------------- Devnet -------------------------

devnet-up:
    WAIT
        BUILD +devnet-start-l1
    END
    WAIT
        BUILD +devnet-start-l2
    END

devnet-start-l1:
    ARG COMPOSE_PROFILES
    ARG DOCKER_REGISTRY_URL
    WAIT
        BUILD +devnet-pull-containers
    END
    # bring l1 up
    WAIT
        BUILD +devnet-l1-up
    END
    # save the current block that we base the L2 genesis start on
    LOCALLY
    RUN cast block --json --full --rpc-url http://localhost:8545 > ./.devnet/l1-starting-block.json
    # bring up the l2 components
    WAIT
        BUILD +devnet-l2-genesis
    END

devnet-start-l2:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    ARG COMPOSE_PROFILES
    ARG LOG_LEVEL
    ARG LOG_FORMAT
    ARG LOG_COLOR
    ARG METRICS_ENABLED
    ARG PPROF_ENABLED
    ARG NATS_ADDRESS
    ARG PROOFSTORE_DBDSN
    ARG BLOBSTORE_ENDPOINT
    ARG BLOBSTORE_ACCESSKEYID
    ARG BLOBSTORE_SECRETACCESSKEY
    ARG BLOBSTORE_BUCKET
    ARG BLOBSTORE_REGION
    ARG BLOBSTORE_S3FORCEPATHSTYLE
    ARG BLOBSTORE_DISABLESSL
    ARG PROPOSER_L2OUTPUTORACLEADDRESS
    ARG PROPOSER_MNEMONIC
    ARG PROPOSER_HDPATH
    ARG TASK_BLOCKSTOCONSUMELIMIT
    ARG TASK_WAITFORBATCHES
    ARG TASK_USEFINALIZED
    ARG BE_DB_HOST
    ARG BE_DB_NAME
    ARG BE_DB_PORT
    ARG BE_API_DB_USER
    ARG BE_API_DB_PASS
    ARG BE_L1C_DB_USER
    ARG BE_L1C_DB_PASS
    ARG BE_L2C_DB_USER
    ARG BE_L2C_DB_PASS
    ARG BE_NATSC_DB_USER
    ARG BE_NATSC_DB_PASS
    ARG INDEXER_CONSUMER_SEQUENCER
    ARG L1_URL_HTTP
    ARG L1_URL_WS
    ARG L1_URL_BEACON
    ARG L2_URL_HTTP
    ARG L2_URL_WS
    ARG ROLLUP_RPC
    ARG OP_NODE_L2_ENGINE_RPC
    ARG OP_NODE_ROLLUP_CONFIG
    ARG OP_NODE_NETWORK
    ARG OP_NODE_RPC_ENABLE_ADMIN
    ARG OP_NODE_L1_TRUST_RPC
    ARG OP_NODE_L1_RPC_KIND
    ARG OP_NODE_L1_RPC_RATE_LIMIT
    ARG OP_NODE_L1_RPC_MAX_BATCH_SIZE
    ARG OP_NODE_L1_HTTP_POLL_INTERVAL
    ARG OP_NODE_L2_ENGINE_AUTH
    ARG OP_NODE_VERIFIER_L1_CONFS
    ARG OP_NODE_SEQUENCER_ENABLED
    ARG OP_NODE_SEQUENCER_STOPPED
    ARG OP_NODE_SEQUENCER_MAX_SAFE_LAG
    ARG OP_NODE_SEQUENCER_L1_CONFS
    ARG OP_NODE_L1_EPOCH_POLL_INTERVAL
    ARG OP_NODE_SNAPSHOT_LOG
    ARG OP_NODE_HEARTBEAT_ENABLED
    ARG OP_NODE_HEARTBEAT_MONIKER
    ARG OP_NODE_HEARTBEAT_URL
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC_TRUST_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_AUTH
    ARG OP_NODE_P2P_SEQUENCER_KEY
    ARG OP_NODE_P2P_PEER_SCORING
    ARG OP_NODE_P2P_PEER_BANNING
    ARG SEQUENCER_ROLLUP_HOSTNAME
    ARG SEQUENCER_ROLLUP_DISCOVERY_PORT
    ARG OP_NODE_P2P_PRIV_PATH
    ARG OP_NODE_SYNCMODE
    ARG MNEMONIC
    ARG OP_PROPOSER_HD_PATH
    ARG OP_PROPOSER_L2OO_ADDRESS
    ARG OP_PROPOSER_POLL_INTERVAL
    ARG OP_PROPOSER_ALLOW_NON_FINALIZED
    ARG OP_PROPOSER_NUM_CONFIRMATIONS
    ARG OP_PROPOSER_SAFE_ABORT_NONCE_TOO_LOW_COUNT
    ARG OP_PROPOSER_SAFE_ABORT_STUCK_IN_GAS_FEE_INCREASE_LOOP
    ARG OP_PROPOSER_RESUBMISSION_TIMEOUT
    ARG OP_PROPOSER_NETWORK_TIMEOUT
    ARG OP_PROPOSER_TXMGR_TX_SEND_TIMEOUT
    ARG OP_PROPOSER_TXMGR_TX_NOT_IN_MEMPOOL_TIMEOUT
    ARG OP_PROPOSER_TXMGR_RECEIPT_QUERY_INTERVAL
    ARG OP_BATCHER_HD_PATH
    ARG OP_BATCHER_SUB_SAFETY_MARGIN
    ARG OP_BATCHER_POLL_INTERVAL
    ARG OP_BATCHER_MAX_PENDING_TX
    ARG OP_BATCHER_MAX_CHANNEL_DURATION
    ARG OP_BATCHER_MAX_L1_TX_SIZE_BYTES
    ARG OP_BATCHER_STOPPED
    ARG OP_BATCHER_NUM_CONFIRMATIONS
    ARG OP_BATCHER_SAFE_ABORT_NONCE_TOO_LOW_COUNT
    ARG OP_BATCHER_SAFE_ABORT_STUCK_IN_GAS_FEE_INCREASE_LOOP
    ARG OP_BATCHER_RESUBMISSION_TIMEOUT
    ARG OP_BATCHER_NETWORK_TIMEOUT
    ARG OP_BATCHER_TXMGR_TX_SEND_TIMEOUT
    ARG OP_BATCHER_TXMGR_TX_NOT_IN_MEMPOOL_TIMEOUT
    ARG OP_BATCHER_TXMGR_RECEIPT_QUERY_INTERVAL
    ARG OP_BATCHER_RPC_ENABLE_ADMIN
    ARG OP_BATCHER_TARGET_L1_TX_SIZE_BYTES
    ARG OP_BATCHER_TARGET_NUM_FRAMES
    ARG OP_BATCHER_APPROX_COMPR_RATIO
    ARG OP_BATCHER_COMPRESSOR
    ARG OP_BATCHER_BATCH_TYPE
    ARG OP_BATCHER_DATA_AVAILABILITY_TYPE
    ARG DOCKER_REGISTRY_URL
    WAIT
        BUILD +docker-all
        BUILD +devnet-download-params
    END
    LOCALLY

    WORKDIR ./ops-bedrock
    RUN bash ./start_system.sh

devnet-canyon-hardfork:
    ARG COMPOSE_PROFILES
    ARG LOG_LEVEL
    ARG LOG_FORMAT
    ARG LOG_COLOR
    ARG METRICS_ENABLED
    ARG PPROF_ENABLED
    ARG NATS_ADDRESS
    ARG PROOFSTORE_DBDSN
    ARG BLOBSTORE_ENDPOINT
    ARG BLOBSTORE_ACCESSKEYID
    ARG BLOBSTORE_SECRETACCESSKEY
    ARG BLOBSTORE_BUCKET
    ARG BLOBSTORE_REGION
    ARG BLOBSTORE_S3FORCEPATHSTYLE
    ARG BLOBSTORE_DISABLESSL
    ARG PROPOSER_L2OUTPUTORACLEADDRESS
    ARG PROPOSER_MNEMONIC
    ARG PROPOSER_HDPATH
    ARG TASK_BLOCKSTOCONSUMELIMIT
    ARG TASK_WAITFORBATCHES
    ARG BE_DB_HOST
    ARG BE_DB_NAME
    ARG BE_DB_PORT
    ARG BE_API_DB_USER
    ARG BE_API_DB_PASS
    ARG BE_L1C_DB_USER
    ARG BE_L1C_DB_PASS
    ARG BE_L2C_DB_USER
    ARG BE_L2C_DB_PASS
    ARG BE_NATSC_DB_USER
    ARG BE_NATSC_DB_PASS
    ARG INDEXER_CONSUMER_SEQUENCER
    ARG L1_URL_HTTP
    ARG L1_URL_WS
    ARG L1_URL_BEACON
    ARG L2_URL_HTTP
    ARG L2_URL_WS
    ARG ROLLUP_RPC
    ARG OP_NODE_L2_ENGINE_RPC
    ARG OP_NODE_ROLLUP_CONFIG
    ARG OP_NODE_NETWORK
    ARG OP_NODE_RPC_ENABLE_ADMIN
    ARG OP_NODE_L1_TRUST_RPC
    ARG OP_NODE_L1_RPC_KIND
    ARG OP_NODE_L1_RPC_RATE_LIMIT
    ARG OP_NODE_L1_RPC_MAX_BATCH_SIZE
    ARG OP_NODE_L1_HTTP_POLL_INTERVAL
    ARG OP_NODE_L2_ENGINE_AUTH
    ARG OP_NODE_VERIFIER_L1_CONFS
    ARG OP_NODE_SEQUENCER_ENABLED
    ARG OP_NODE_SEQUENCER_STOPPED
    ARG OP_NODE_SEQUENCER_MAX_SAFE_LAG
    ARG OP_NODE_SEQUENCER_L1_CONFS
    ARG OP_NODE_L1_EPOCH_POLL_INTERVAL
    ARG OP_NODE_SNAPSHOT_LOG
    ARG OP_NODE_HEARTBEAT_ENABLED
    ARG OP_NODE_HEARTBEAT_MONIKER
    ARG OP_NODE_HEARTBEAT_URL
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC_TRUST_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_AUTH
    ARG OP_NODE_P2P_SEQUENCER_KEY
    ARG OP_NODE_P2P_PEER_SCORING
    ARG OP_NODE_P2P_PEER_BANNING
    ARG SEQUENCER_ROLLUP_HOSTNAME
    ARG SEQUENCER_ROLLUP_DISCOVERY_PORT
    ARG OP_NODE_P2P_PRIV_PATH
    ARG OP_NODE_SYNCMODE
    ARG MNEMONIC
    ARG OP_PROPOSER_HD_PATH
    ARG OP_PROPOSER_L2OO_ADDRESS
    ARG OP_PROPOSER_POLL_INTERVAL
    ARG OP_PROPOSER_ALLOW_NON_FINALIZED
    ARG OP_PROPOSER_NUM_CONFIRMATIONS
    ARG OP_PROPOSER_SAFE_ABORT_NONCE_TOO_LOW_COUNT
    ARG OP_PROPOSER_SAFE_ABORT_STUCK_IN_GAS_FEE_INCREASE_LOOP
    ARG OP_PROPOSER_RESUBMISSION_TIMEOUT
    ARG OP_PROPOSER_NETWORK_TIMEOUT
    ARG OP_PROPOSER_TXMGR_TX_SEND_TIMEOUT
    ARG OP_PROPOSER_TXMGR_TX_NOT_IN_MEMPOOL_TIMEOUT
    ARG OP_PROPOSER_TXMGR_RECEIPT_QUERY_INTERVAL
    ARG OP_BATCHER_HD_PATH
    ARG OP_BATCHER_SUB_SAFETY_MARGIN
    ARG OP_BATCHER_POLL_INTERVAL
    ARG OP_BATCHER_MAX_PENDING_TX
    ARG OP_BATCHER_MAX_CHANNEL_DURATION
    ARG OP_BATCHER_MAX_L1_TX_SIZE_BYTES
    ARG OP_BATCHER_STOPPED
    ARG OP_BATCHER_NUM_CONFIRMATIONS
    ARG OP_BATCHER_SAFE_ABORT_NONCE_TOO_LOW_COUNT
    ARG OP_BATCHER_SAFE_ABORT_STUCK_IN_GAS_FEE_INCREASE_LOOP
    ARG OP_BATCHER_RESUBMISSION_TIMEOUT
    ARG OP_BATCHER_NETWORK_TIMEOUT
    ARG OP_BATCHER_TXMGR_TX_SEND_TIMEOUT
    ARG OP_BATCHER_TXMGR_TX_NOT_IN_MEMPOOL_TIMEOUT
    ARG OP_BATCHER_TXMGR_RECEIPT_QUERY_INTERVAL
    ARG OP_BATCHER_RPC_ENABLE_ADMIN
    ARG OP_BATCHER_TARGET_L1_TX_SIZE_BYTES
    ARG OP_BATCHER_TARGET_NUM_FRAMES
    ARG OP_BATCHER_APPROX_COMPR_RATIO
    ARG OP_BATCHER_COMPRESSOR
    ARG OP_BATCHER_BATCH_TYPE
    ARG OP_BATCHER_DATA_AVAILABILITY_TYPE
    ARG DOCKER_REGISTRY_URL
    ARG FORK_TRIGGER=true
    LOCALLY
    RUN ./ops-bedrock/canyon_hardfork.sh

devnet-down:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    LOCALLY
    WORKDIR ./ops-bedrock
    RUN COMPOSE_PROFILES=replica docker compose down

devnet-stop:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    LOCALLY
    WORKDIR ./ops-bedrock
    RUN docker compose stop op-batcher
    RUN docker compose stop op-node
    RUN docker compose stop $(docker compose config --services | grep -Ev "\<(l1|create-beacon-chain-genesis|geth-genesis|beacon-chain|validator|op-node|l2-replica)\>")

devnet-start-replica:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    ARG LOG_LEVEL
    ARG METRICS_ENABLED
    ARG PPROF_ENABLED
    ARG L1_URL_HTTP
    ARG L1_URL_WS
    ARG L1_URL_BEACON
    ARG L2_URL_HTTP
    ARG L2_URL_WS
    ARG OP_NODE_ROLLUP_CONFIG
    ARG OP_NODE_RPC_ENABLE_ADMIN
    ARG OP_NODE_L1_TRUST_RPC
    ARG OP_NODE_L1_RPC_KIND
    ARG OP_NODE_L1_RPC_RATE_LIMIT
    ARG OP_NODE_L1_RPC_MAX_BATCH_SIZE
    ARG OP_NODE_L1_HTTP_POLL_INTERVAL
    ARG OP_NODE_VERIFIER_L1_CONFS
    ARG OP_NODE_SEQUENCER_STOPPED
    ARG OP_NODE_SEQUENCER_MAX_SAFE_LAG
    ARG OP_NODE_SEQUENCER_L1_CONFS
    ARG OP_NODE_L1_EPOCH_POLL_INTERVAL
    ARG OP_NODE_SNAPSHOT_LOG
    ARG OP_NODE_HEARTBEAT_ENABLED
    ARG OP_NODE_HEARTBEAT_MONIKER
    ARG OP_NODE_HEARTBEAT_URL
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC
    ARG OP_NODE_L2_BACKUP_UNSAFE_SYNC_RPC_TRUST_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_RPC
    ARG OP_NODE_REPLICA_L2_ENGINE_AUTH
    ARG OP_NODE_P2P_PEER_SCORING
    ARG OP_NODE_P2P_PEER_BANNING
    ARG SEQUENCER_ROLLUP_HOSTNAME
    ARG SEQUENCER_ROLLUP_DISCOVERY_PORT
    ARG OP_NODE_P2P_PRIV_PATH
    ARG DOCKER_REGISTRY_URL
    LOCALLY

    WORKDIR ./ops-bedrock
    RUN docker compose up l2-replica op-node-replica -d

devnet-stop-replica:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    LOCALLY
    WORKDIR ./ops-bedrock
    RUN docker compose stop op-node-replica
    RUN docker compose stop l2-replica

devnet-clean:
    WAIT
        BUILD +devnet-down
    END
    LOCALLY
    RUN rm -rf ./.devnet
    RUN rm -rf ./packages/contracts-bedrock/deployments/devnetL1
    RUN docker image ls 'ops-bedrock*' --format='{{.Repository}}' | xargs -r docker rmi
    RUN docker volume ls --filter name=ops-bedrock --format='{{.Name}}' | xargs -r docker volume rm

devnet-pull-containers:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    LOCALLY
    WORKDIR ./ops-bedrock
    RUN docker compose pull --ignore-pull-failures

devnet-download-params:
    LOCALLY
    RUN mkdir -p ./params
    IF [ "${DOCKER_REGISTRY_URL}" != "" ]
        RUN bash ./ops-bedrock/params.sh
    END

devnet-base:
    FROM +yarn-copy-source
    COPY ./packages/contracts-bedrock /src/packages/contracts-bedrock
    COPY ./zr-proof-orchestrator/common/types /src/zr-proof-orchestrator/common/types
    COPY ./bedrock-devnet /src/bedrock-devnet
    COPY ./ops-bedrock /src/ops-bedrock


devnet-l1-genesis:
    FROM +devnet-base
    WORKDIR /src
    RUN PYTHONPATH=/src/bedrock-devnet python3 ./bedrock-devnet/main.py --monorepo-dir=. --allocs
    RUN --no-cache PYTHONPATH=/src/bedrock-devnet python3 ./bedrock-devnet/main.py --monorepo-dir=. --genesis-l1
    SAVE ARTIFACT --force /src/.devnet/allocs-l1.json /.devnet/allocs-l1.json AS LOCAL ./.devnet/allocs-l1.json
    SAVE ARTIFACT --force /src/.devnet/allocs-l2.json /.devnet/allocs-l2.json AS LOCAL ./.devnet/allocs-l2.json
    SAVE ARTIFACT --force /src/.devnet/allocs-l2-delta.json /.devnet/allocs-l2-delta.json AS LOCAL ./.devnet/allocs-l2-delta.json
    SAVE ARTIFACT --force /src/.devnet/addresses.json /.devnet/addresses.json AS LOCAL ./.devnet/addresses.json
    SAVE ARTIFACT --force /src/.devnet/genesis-l1.json /.devnet/genesis-l1.json AS LOCAL ./.devnet/genesis-l1.json
    SAVE ARTIFACT --force /src/packages/contracts-bedrock/deployments/devnetL1/ /deployments/ AS LOCAL ./packages/contracts-bedrock/deployments/
    SAVE ARTIFACT --force /src/packages/contracts-bedrock/deploy-config/devnetL1.json /deploy-config/ AS LOCAL ./packages/contracts-bedrock/deploy-config/devnetL1.json

devnet-l1-up:
    ARG FORK
    ARG BEACON_CHAIN_GENESIS_IMAGE
    ARG GETH_IMAGE
    ARG PRYSM_IMAGE
    ARG PRYSM_VALIDATOR_IMAGE
    WAIT
        BUILD +devnet-l1-genesis
    END
    LOCALLY
    WORKDIR ./ops-bedrock
    RUN bash ./start_l1.sh

devnet-l2-genesis:
    FROM +devnet-l1-genesis
    COPY ./.devnet/l1-starting-block.json /src/.devnet/l1-starting-block.json
    WORKDIR /src
    RUN PYTHONPATH=/src/bedrock-devnet python3 ./bedrock-devnet/main.py --monorepo-dir=. --genesis-l2
    SAVE ARTIFACT --force /src/.devnet/genesis-l2.json AS LOCAL ./.devnet/genesis-l2.json
    SAVE ARTIFACT --force /src/.devnet/rollup.json AS LOCAL ./.devnet/rollup.json

# ------------------------- Documentation Targets -------------------------
contract-addresses-markdown:
    FROM +go-build-copy-source
    COPY ./docs /src/docs
    COPY ./packages/contracts-bedrock/deployments /src/packages/contracts-bedrock/deployments
    WORKDIR /src/docs
    RUN go run ./generate_address_table.go /src/packages/contracts-bedrock/deployments Mainnet > /contract_addresses_mainnet.md
    RUN go run ./generate_address_table.go /src/packages/contracts-bedrock/deployments Testnet > /contract_addresses_testnet.md
    SAVE ARTIFACT --force /contract_addresses_mainnet.md AS LOCAL ./docs/contract_addresses_mainnet.md
    SAVE ARTIFACT --force /contract_addresses_testnet.md AS LOCAL ./docs/contract_addresses_testnet.md
