# Databases

This directory contains database schemas and scripts to run the database for the proof-orchestrator and block explorer api.


#### Start proof-orchestrator DB

Starts a postgres docker container and runs all the sql files located under `proof-orchestrator/schemas/*/*.sql` and `proof-orchestrator/*.sql`.

```zsh
./docker-run.sh
```

#### Connect to proof-orchestrator DB

Connect to the proof-orchestrator db using the shell script. Or see the connection URL string in `connect-db.sh` script.

```zsh
./connect-db.sh
```

#### Start block-explorer DB

Starts a postgres docker container and runs all the sql files located under `block-explorer/schemas/*/*.sql` and `block-explorer/*.sql`.

```zsh
./block-explorer/docker-run.sh
```

#### Connect to block-explorer DB

Connect to the block-explorer db using the shell script. Or see the connection URL string in `connect-db.sh` script.

```zsh
./block-explorer/connect-db.sh
```
