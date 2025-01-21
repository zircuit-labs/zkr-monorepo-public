-- Drop the existing index
DROP INDEX IF EXISTS idx_average_transaction_fee_by_day;

-- Drop the existing materialized view
DROP MATERIALIZED VIEW IF EXISTS block_explorer.average_transaction_fee_by_day;

-- Recreate the updated materialized view
CREATE MATERIALIZED VIEW IF NOT EXISTS block_explorer.average_transaction_fee_by_day AS
SELECT transaction_date,
       CEIL(AVG(CASE WHEN fee IS NOT NULL AND fee != '' THEN CAST(fee AS DECIMAL(21, 0)) ELSE 0 END)) AS value
FROM block_explorer."Transactions"
GROUP BY transaction_date
ORDER BY transaction_date DESC;

-- Recreate the unique index
CREATE UNIQUE INDEX idx_average_transaction_fee_by_day ON block_explorer.average_transaction_fee_by_day (transaction_date);
