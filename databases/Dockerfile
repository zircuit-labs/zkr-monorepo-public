FROM public.ecr.aws/docker/library/postgres:16-alpine
# We need the schema change files to be uniquely named and ordered alphabetically in order they should execute
COPY proof-orchestrator/schemas/*/*.sql /docker-entrypoint-initdb.d/
COPY proof-orchestrator/seed-data* /docker-entrypoint-initdb.d/
