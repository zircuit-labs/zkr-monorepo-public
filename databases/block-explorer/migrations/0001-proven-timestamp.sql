-- 0001-proven-timestamp.sql

ALTER TABLE block_explorer."Transactions" ADD proven_timestamp int8 NULL;
ALTER TABLE block_explorer."Transactions" ADD finalization_period_seconds int8 NULL;
