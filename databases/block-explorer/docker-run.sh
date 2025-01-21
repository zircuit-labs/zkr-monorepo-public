docker build . -t block-explorer-db
port=5433
echo "
------------------------------------------------------------
Starting DB on port $port...
------------------------------------------------------------
"
docker run --publish 5433:5432 -e POSTGRES_PASSWORD=beadmin123 -e POSTGRES_USER=beadmin -e POSTGRES_DB=blockexplorer_test block-explorer-db
