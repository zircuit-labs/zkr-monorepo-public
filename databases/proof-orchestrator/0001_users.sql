CREATE USER poadmin WITH PASSWORD 'SANITIZED';

GRANT ALL ON DATABASE proof_orchestrator TO poadmin;

GRANT ALL ON SCHEMA proof_state TO poadmin;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA proof_state TO poadmin;

-- Read only user for block explorer service
CREATE USER poreadonly WITH PASSWORD 'SANITIZED' GRANT CONNECT ON DATABASE proof_orchestrator TO poreadonly;

-- This assumes you're actually connected to proof_orchestrator when running the grant
GRANT USAGE ON SCHEMA proof_state TO poreadonly;

GRANT SELECT ON ALL TABLES IN SCHEMA proof_state TO poreadonly;

-- Read only user for DEVs
CREATE USER devreadonly WITH LOGIN PASSWORD 'SANITIZED';

GRANT CONNECT ON DATABASE proof_orchestrator TO devreadonly;

GRANT pg_read_all_data, pg_monitor TO devreadonly;

