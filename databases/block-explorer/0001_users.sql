-- Read only user for DEVs
CREATE USER devreadonly WITH LOGIN PASSWORD 'SANTIZED';

GRANT CONNECT ON DATABASE block_explorer TO devreadonly;

GRANT pg_read_all_data, pg_monitor TO devreadonly;

