docker build . -t proof-orchestrator-db
port=8880
echo "
------------------------------------------------------------
Starting DB on port $port...
------------------------------------------------------------
"
docker run --publish 8880:5432 -e POSTGRES_PASSWORD=local -e POSTGRES_USER=poadmin -e POSTGRES_DB=proof_orchestrator proof-orchestrator-db
