SET TIMEZONE TO 'UTC';

CREATE OR REPLACE FUNCTION block_explorer.epoch_to_date(epoch BIGINT)
    RETURNS DATE AS
$$
begin
    RETURN to_timestamp(epoch)::date;
END;
$$ LANGUAGE plpgsql IMMUTABLE;

ALTER TABLE block_explorer."Transactions"
    ADD COLUMN transaction_date DATE GENERATED ALWAYS AS (block_explorer.epoch_to_date("timestamp")) STORED;

CREATE INDEX idx_transactions_transaction_date ON "block_explorer"."Transactions" (transaction_date);

-------------------------transactions_count_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.transactions_count_by_day AS
SELECT transaction_date,
       COUNT(*) AS value
FROM block_explorer."Transactions"
GROUP BY transaction_date
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_transactions_count_by_day ON block_explorer.transactions_count_by_day (transaction_date);

-------------------------unique_addresses_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.unique_addresses_by_day AS
WITH uniqueaddresses AS (SELECT from_addr AS addr, transaction_date
                         FROM block_explorer."Transactions"
                         UNION ALL
                         SELECT to_addr AS addr, transaction_date
                         FROM block_explorer."Transactions")
SELECT transaction_date,
       COUNT(DISTINCT addr) AS value
FROM uniqueaddresses
GROUP BY transaction_date
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_unique_addresses_by_day ON block_explorer.unique_addresses_by_day (transaction_date);
CREATE INDEX idx_transactions_from_addr_date ON block_explorer."Transactions" (from_addr, transaction_date);
CREATE INDEX idx_transactions_to_addr_date ON block_explorer."Transactions" (to_addr, transaction_date);

-------------------------cumulative_unique_addresses_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.cumulative_unique_addresses_by_day AS
SELECT transaction_date,
       SUM(value) OVER (
           ORDER BY transaction_date ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS value
FROM block_explorer.unique_addresses_by_day
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_cumulative_unique_addresses_by_day ON block_explorer.cumulative_unique_addresses_by_day (transaction_date);

-------------------------deployed_contracts_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.deployed_contracts_by_day AS
SELECT transaction_date,
       COUNT(*) AS value
FROM block_explorer."Transactions"
WHERE (is_smart_contract_creation)
GROUP BY transaction_date
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_deployed_contracts_by_day ON block_explorer.deployed_contracts_by_day (transaction_date);

-------------------------total_gas_used_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.total_gas_used_by_day AS
SELECT transaction_date,
       SUM(CASE WHEN gas_used IS NOT NULL AND gas_used != '' THEN CAST(gas_used AS DECIMAL(21, 0)) ELSE 0 END) AS value
FROM block_explorer."Transactions"
GROUP BY "transaction_date"
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_total_gas_used_by_day ON block_explorer.total_gas_used_by_day (transaction_date);


-------------------------total_transaction_fee_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.total_transaction_fee_by_day AS
SELECT transaction_date,
       SUM(CASE WHEN fee IS NOT NULL AND fee != '' THEN CAST(fee AS DECIMAL(21, 0)) ELSE 0 END) AS value
FROM block_explorer."Transactions"
GROUP BY transaction_date
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_total_transaction_fee_by_day ON block_explorer.total_transaction_fee_by_day (transaction_date);

-------------------------average_transaction_fee_by_day-------------------------

CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.average_transaction_fee_by_day AS
SELECT
    transaction_date,
    AVG(CASE WHEN fee IS NOT NULL AND fee != '' THEN CAST(fee AS DECIMAL(21, 0)) ELSE 0 END) AS value
FROM
    block_explorer."Transactions"
GROUP BY
    transaction_date
ORDER BY
    transaction_date DESC;

CREATE UNIQUE INDEX idx_average_transaction_fee_by_day ON block_explorer.average_transaction_fee_by_day (transaction_date);



