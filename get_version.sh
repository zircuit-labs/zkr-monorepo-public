#!/bin/bash
GITCOMMIT="$(git rev-parse HEAD)"
GITDATE="$(git show -s --format='%cd' --date=iso-strict)"
GITBRANCH="$(git branch --show-current)"
VERSION="v0.0.0"
META=""

echo "{
    \"git_commit\": \"$GITCOMMIT\",
    \"git_date\": \"$GITDATE\",
    \"git_branch\": \"$GITBRANCH\",
    \"version\": \"$VERSION\",
    \"meta\": \"$META\",
    \"prover_commit\": \"$MODULAR_PROVER_COMMIT\",
    \"l2geth_commit\": \"$L2GETH_COMMIT\"
}" > version.json
