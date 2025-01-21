-- Drop the existing index
DROP INDEX IF EXISTS idx_cumulative_unique_addresses_by_day;
DROP INDEX IF EXISTS idx_unique_addresses_by_day;

-- Drop the existing materialized view
DROP MATERIALIZED VIEW IF EXISTS block_explorer.cumulative_unique_addresses_by_day;
DROP MATERIALIZED VIEW IF EXISTS block_explorer.unique_addresses_by_day;

-- Recreate the updated unique addresses materialized view
CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.unique_addresses_by_day AS
WITH uniqueaddresses AS (SELECT from_addr AS addr, transaction_date
                         FROM block_explorer."Transactions"
                         UNION ALL
                         SELECT to_addr AS addr, transaction_date
                         FROM block_explorer."Transactions"),
     firstseendate AS (SELECT addr,
                              MIN(transaction_date) AS first_transaction_date
                       FROM uniqueaddresses
                       GROUP BY addr)
SELECT first_transaction_date AS transaction_date,
       COUNT(DISTINCT addr)   AS value
FROM firstseendate
GROUP BY first_transaction_date
ORDER BY first_transaction_date DESC;

CREATE UNIQUE INDEX idx_unique_addresses_by_day ON block_explorer.unique_addresses_by_day (transaction_date);

-- Recreate the updated cumulative unique addresses materialized view
CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.cumulative_unique_addresses_by_day AS
SELECT transaction_date,
       SUM(value) OVER (
           ORDER BY transaction_date ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS value
FROM block_explorer.unique_addresses_by_day
ORDER BY transaction_date DESC;

CREATE UNIQUE INDEX idx_cumulative_unique_addresses_by_day ON block_explorer.cumulative_unique_addresses_by_day (transaction_date);
