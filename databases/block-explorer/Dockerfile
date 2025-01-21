FROM postgres:16.2

# We need the schema change files to be uniquely named and ordered alphabetically in order they should execute
COPY ./migrations/*.sql /docker-entrypoint-initdb.d/
