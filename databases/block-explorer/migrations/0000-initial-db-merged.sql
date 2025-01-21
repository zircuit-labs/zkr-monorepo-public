-- 0000-initial-db.sql


-- Create Schema

CREATE SCHEMA IF NOT EXISTS block_explorer;

-- Create users

--
-- USER           | Service                             | Roles
-- ---------------|-------------------------------------|----------------------------------------------
-- beadmin        | No specific service                 | Admin (full access)
-- beapi          | Block Explorer API                  | Read-only access
-- bel1consumer   | Block Explorer L1 Consumer          | CRUD permissions
-- bel2consumer   | Block Explorer L2 Consumer          | CRUD permissions
-- benatsconsumer | Block Explorer NATS Consumer        | CRUD permissions
-- befakegenerator| Block Explorer Fake Data Generator  | CRUD permissions

-- Creating roles (users)

-- Passwords should be random generated for each environment and stored in the correct vaults.
CREATE USER beadmin WITH PASSWORD 'SANITIZED';
CREATE USER beapi WITH PASSWORD 'SANITIZED';
CREATE USER bel1consumer WITH PASSWORD 'SANITIZED';
CREATE USER bel2consumer WITH PASSWORD 'SANITIZED';
CREATE USER benatsconsumer WITH PASSWORD 'SANITIZED';
CREATE USER befakegenerator WITH PASSWORD 'SANITIZED';

-- Setting up permissions

-- beadmin: Admin (full access)
ALTER SCHEMA block_explorer OWNER TO beadmin;
GRANT ALL PRIVILEGES ON SCHEMA block_explorer TO beadmin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA block_explorer TO beadmin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA block_explorer TO beadmin;

ALTER DEFAULT PRIVILEGES IN SCHEMA block_explorer
    GRANT ALL ON TABLES TO beadmin;
ALTER DEFAULT PRIVILEGES IN SCHEMA block_explorer
    GRANT ALL ON SEQUENCES TO beadmin;

SET ROLE beadmin;

-- beapi: Read-only access
GRANT USAGE ON SCHEMA block_explorer TO beapi;
GRANT SELECT ON ALL TABLES IN SCHEMA block_explorer TO beapi;

ALTER DEFAULT PRIVILEGES IN SCHEMA block_explorer
    GRANT SELECT ON TABLES TO beapi;

-- Set permissions for CRUD users
CREATE OR REPLACE FUNCTION block_explorer.set_permissions_crud_user(role_name text)
    RETURNS void LANGUAGE plpgsql AS $$
BEGIN
    -- Grant CRUD permissions on current tables in block_explorer schema
    EXECUTE format('GRANT USAGE ON SCHEMA block_explorer TO %I', role_name);
    EXECUTE format('GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA block_explorer TO %I', role_name);
    EXECUTE format('GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA block_explorer TO %I', role_name);

    -- Set default privileges for future tables and sequences in block_explorer schema
    EXECUTE format('ALTER DEFAULT PRIVILEGES IN SCHEMA block_explorer GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO %I', role_name);
    EXECUTE format('ALTER DEFAULT PRIVILEGES IN SCHEMA block_explorer GRANT USAGE, SELECT, UPDATE ON SEQUENCES TO %I', role_name);

END;
$$;

-- CRUD users
SELECT block_explorer.set_permissions_crud_user('bel1consumer');
SELECT block_explorer.set_permissions_crud_user('bel2consumer');
SELECT block_explorer.set_permissions_crud_user('benatsconsumer');
SELECT block_explorer.set_permissions_crud_user('befakegenerator');

CREATE TABLE block_explorer."EventLogsData"
(
    event_log_id serial8 NOT NULL,
    "name"       varchar NULL,
    value        varchar NULL,
    "type"       varchar NULL,
    CONSTRAINT "EventLogsData_pkey" PRIMARY KEY (event_log_id)
);

CREATE TABLE block_explorer."InternalTransactions"
(
    "transaction"      varchar NULL,
    sender             varchar NULL,
    receiver           varchar NULL,
    type_trace_address varchar NULL,
    "type"             varchar NULL,
    value              varchar NULL
);

CREATE TABLE block_explorer."Batches"
(
    hash         varchar PRIMARY KEY,
    size         int       NULL,
    proof        varchar   NULL,
    status       varchar   NULL,
    created_at   timestamp NULL,
    published_at timestamp NULL
);

CREATE TABLE block_explorer."Transactions"
(
    l2_tx_hash                 varchar PRIMARY KEY,
    l2_block_index             int8    NULL,
    batch_hash                 varchar NULL,
    tx_index                   int8    NULL,
    tx_type                    varchar NULL,
    nonce                      varchar NULL,
    l1_block_index             int8    NULL,
    l1_tx_hash                 varchar NULL,
    l1_tx_origin               varchar NULL,
    confirmed_by               varchar NULL,
    status                     varchar NULL,
    failure_reason             varchar NULL,
    from_addr                  varchar NULL,
    to_addr                    varchar NULL,
    "timestamp"                int8    NULL,
    fee                        varchar NULL,
    value                      varchar NULL,
    gas_limit                  varchar NULL,
    gas_used                   varchar NULL,
    gas_price                  varchar NULL,
    block_id                   varchar NULL,
    input_data                 varchar NULL,
    is_smart_contract_creation bool    NULL
);

CREATE TABLE block_explorer."Accounts"
(
    address                varchar NOT NULL,
    byte_code              varchar NULL,
    deployment_transaction varchar NULL,
    contract_creator       varchar NULL,
    CONSTRAINT "Accounts_pkey" PRIMARY KEY (address)
);

CREATE TABLE block_explorer."EventLogs"
(
    id            serial8 NOT NULL,
    index         int     NOT NULL,
    "transaction" varchar NULL,
    address       varchar NULL,
    signature     varchar NULL,
    topics        jsonb   NULL,
    raw_data      varchar NULL,
    CONSTRAINT "EventLogs_pkey" PRIMARY KEY (id),
    CONSTRAINT "EventLogs_transaction_fkey" FOREIGN KEY (transaction) REFERENCES block_explorer."Transactions" (l2_tx_hash) ON DELETE CASCADE
);

CREATE TABLE block_explorer."Blocks"
(
    block_index int8 PRIMARY KEY,
    batch_hash  varchar NULL,
    "timestamp" int8    NULL,
    id          varchar NULL,
    parent_id   varchar NULL,
    status      varchar NULL
);
-- 0001-internal-tx-new-fields.sql


ALTER TABLE block_explorer."InternalTransactions"
    ADD gas_limit varchar NULL;
ALTER TABLE block_explorer."InternalTransactions"
    ADD block_index int8 NULL;
ALTER TABLE block_explorer."InternalTransactions"
    ADD "timestamp" int8 NULL;
-- 0002-tx-nonce-version.sql


ALTER TABLE block_explorer."Transactions"
    ADD message_nonce varchar NULL;
ALTER TABLE block_explorer."Transactions"
    ADD messenger_version varchar NULL;
ALTER TABLE block_explorer."Transactions"
    ADD withdrawal_hash varchar NULL;

-- This index is designed to capture unique combinations of message_nonce, messenger_version, and withdrawal_hash under the following conditions:
-- 1. If values are provided for message_nonce and messenger_version, withdrawal_hash should contain an empty string to enforce uniqueness (L1L2).
-- 2. If a value is provided for withdrawal_hash, then message_nonce and messenger_version should contain empty strings to enforce uniqueness (L2L1).
-- 3. If all three fields (message_nonce, messenger_version, and withdrawal_hash) are null, no uniqueness is enforced, allowing multiple such entries (L2).
-- This ensures that the unique combination of f2, f3, and f4 is only applied when all three fields are not null
CREATE UNIQUE INDEX transaction_crosschain_unique_nonce_tx_type
    ON block_explorer."Transactions" (message_nonce, messenger_version, withdrawal_hash, tx_type)
    WHERE message_nonce IS NOT NULL AND messenger_version IS NOT NULL AND withdrawal_hash IS NOT NULL;
-- 0003-tx-l1-new-fields.sql


ALTER TABLE block_explorer."Transactions"
    ADD l1_fee_scalar varchar NULL;
ALTER TABLE block_explorer."Transactions"
    ADD l1_gas_price varchar NULL;
ALTER TABLE block_explorer."Transactions"
    ADD l1_gas_used varchar NULL;
-- 0004-token-mapping-tables.sql


SET ROLE beadmin;

CREATE TABLE block_explorer."TokenMapping"
(
    id          bigserial PRIMARY KEY,
    symbol      varchar   NOT NULL,
    l1_address  varchar   NOT NULL,
    l2_address  varchar   NOT NULL,
    decimals_l1 int8      NOT NULL,
    decimals_l2 int8      NOT NULL,
    "timestamp" timestamp DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO block_explorer."TokenMapping" (symbol, l1_address, l2_address, decimals_l1, decimals_l2)
VALUES
    ('USDC', '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48', 6, 18),
    ('ETH', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2', '0x0000000000000000000000000000000000000000', 18, 18);
-- 0005-l2l1-proven.sql


ALTER TABLE block_explorer."Transactions"
    ADD proven bool NOT NULL DEFAULT FALSE;
ALTER TABLE block_explorer."Transactions"
    ADD finalized bool NOT NULL DEFAULT FALSE;
-- 0006-txn-l1-submission-tx-hash.sql


ALTER TABLE block_explorer."Transactions"
    ADD l1_submission_hash varchar NULL;
-- 0007-erc20transfers.sql


SET ROLE beadmin;

CREATE TABLE block_explorer."ERC20Tokens"
(
    addr     varchar PRIMARY KEY,
    name     varchar NOT NULL,
    symbol   varchar NOT NULL,
    decimals int8    NOT NULL
);

CREATE TABLE block_explorer."ERC20Transfers"
(
    id         bigserial PRIMARY KEY,
    l2_tx_hash varchar NOT NULL,
    token_addr varchar NOT NULL,
    from_addr  varchar NOT NULL,
    to_addr    varchar NOT NULL,
    value      varchar NOT NULL,
    timestamp  int8    NOT NULL
);

ALTER TABLE block_explorer."ERC20Transfers" ADD CONSTRAINT erc20transfers_fk FOREIGN KEY (token_addr) REFERENCES block_explorer."ERC20Tokens"(addr);

-- 0008-block-new-fields.sql


SET ROLE beadmin;

ALTER TABLE block_explorer."Blocks"
    ADD gas_used varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD gas_limit varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD extra_data varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD proof varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD public_inputs varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD zkevm_circuits varchar NULL;
ALTER TABLE block_explorer."Blocks"
    ADD prover_version varchar NULL;


CREATE TABLE block_explorer."BlockSettings"
(
    id             bigserial PRIMARY KEY,
    public_inputs  varchar NOT NULL,
    zkevm_circuits varchar NULL,
    prover_version varchar NULL
);

INSERT INTO block_explorer."BlockSettings" (id, public_inputs, zkevm_circuits, prover_version)
VALUES (1,
        '0x123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef01',
        'circuit_v1_0_1',
        'prover_1.2.3');
-- 0009-txn-l1-state-batch-index.sql


ALTER TABLE block_explorer."Transactions"
    DROP l1_submission_hash;

ALTER TABLE block_explorer."Batches"
    ADD l1_state_batch_index varchar NULL;

ALTER TABLE block_explorer."Batches"
    ADD l1_submission_hash varchar NULL;
-- 0010-transaction-burnt-fee.sql


ALTER TABLE block_explorer."Transactions"
    ADD burnt_fee varchar NULL;
-- 0011-batch-new-fields.sql


SET ROLE beadmin;

ALTER TABLE block_explorer."Batches"
    ADD public_inputs varchar NULL;
ALTER TABLE block_explorer."Batches"
    ADD zkevm_circuits varchar NULL;
ALTER TABLE block_explorer."Batches"
    ADD prover_version varchar NULL;


CREATE TABLE block_explorer."BatchSettings"
(
    id             bigserial PRIMARY KEY,
    public_inputs  varchar NOT NULL,
    zkevm_circuits varchar NULL,
    prover_version varchar NULL
);

INSERT INTO block_explorer."BatchSettings" (id, public_inputs, zkevm_circuits, prover_version)
VALUES (1,
        '0xdd3439ff73342459762e3e80ff67ecb466e6595ac6304b511201b628dd65bbae',
        'circuit_v2_2_2',
        'prover_4.4.4');

-- Recreate created_at and published_at as int.
ALTER TABLE block_explorer."Batches"
    DROP created_at;

ALTER TABLE block_explorer."Batches"
    ADD created_at int8 NULL;

ALTER TABLE block_explorer."Batches"
    DROP published_at;

ALTER TABLE block_explorer."Batches"
    ADD published_at int8 NULL;
-- 0012-block-parent-block-index.sql


ALTER TABLE block_explorer."Blocks"
    ADD parent_block_index int8 NULL;
-- 0013-batch-output-root.sql


ALTER TABLE block_explorer."Batches"
    ADD output_root varchar NULL;
-- 0014-transactions-transaction-batch-index.sql


ALTER TABLE block_explorer."Transactions"
    ADD l1_txn_batch_index int8 NULL;
-- 0015-batch-l1-timestamp.sql


ALTER TABLE block_explorer."Batches"
    ADD l1_timestamp int8 NULL;

-- 0016-optimize-internal-transactions-upsert.sql

CREATE INDEX idx_transactions_l2_block_index ON block_explorer."Transactions" (l2_block_index);
CREATE INDEX idx_internal_transactions_transaction ON block_explorer."InternalTransactions" (transaction);

-- 0017-optimize-batches-endpoint.sql

CREATE INDEX idx_transactions_batch_hash ON block_explorer."Transactions" (batch_hash);

-- 0018-optimize-internal-transactions-select.sql

CREATE INDEX idx_internal_transactions_sender_receiver ON block_explorer."InternalTransactions" (LOWER((sender)::text), (LOWER((receiver)::text)));

-- 0019-optimize-transactions-filter-address-select.sql

CREATE INDEX idx_transactions_to_addr ON block_explorer."Transactions" (LOWER(to_addr));
CREATE INDEX idx_transactions_from_addr ON block_explorer."Transactions" (LOWER(from_addr));

-- 0020-optimize-event-logs-reindex.sql

CREATE INDEX idx_event_logs_transaction ON block_explorer."EventLogs" ("transaction");

-- 0021-create-recent-transctions-endpoint.sql

ALTER TABLE block_explorer."Transactions" ADD finalized_timestamp int8 NULL;

-- 0022-optimize-get-transaction.sql

CREATE INDEX idx_transactions_l2_tx_hash_lower_l2_block_index_tx_index_lower
    ON block_explorer."Transactions" (LOWER(l2_tx_hash), l2_block_index DESC, tx_index DESC);

-- 0023-optimize-get-transaction-with-address-filter.sql

CREATE INDEX idx_transactions_l1_tx_origin_lower
    ON block_explorer."Transactions" (LOWER(l1_tx_origin));

CREATE TABLE block_explorer."BackIndexingFirstBlock"
(
    block_index int8,
    chain char(2) PRIMARY KEY CHECK (chain = 'L1' OR chain = 'L2')
);

-- 0024-create-transactions-index.sql

CREATE INDEX idx_transactions_on_l2_block_index_and_tx_type ON block_explorer."Transactions"(l2_block_index DESC, tx_type);

-- 0025-create-index-single-tx-filter.sql

CREATE INDEX idx_transactions_l2_tx_hash_l2_block_index_tx_index ON block_explorer."Transactions" (LOWER(l2_tx_hash) DESC, l2_block_index DESC, tx_index DESC);

-- 0026-create-effective_gas_price.sql

ALTER TABLE block_explorer."Transactions"
    ADD effective_gas_price varchar NULL;

-- 0027-create-gas-watcher.sql

CREATE TABLE block_explorer."Gas"
(
    transaction_hash             varchar(66) PRIMARY KEY,
    watcher_label                varchar(255),
    transaction_to               varchar(42),
    transaction_from             varchar(42),
    transaction_gas              bigint,
    transaction_gas_price        varchar(255),
    transaction_gas_tip_cap      varchar(255),
    transaction_gas_fee_cap      varchar(255),
    transaction_blob_gas         bigint,
    transaction_blob_gas_fee_cap varchar(255),
    receipt_status               bigint,
    receipt_cumulative_gas_used  bigint,
    receipt_gas_used             bigint,
    receipt_effective_gas_price  varchar(255),
    receipt_blob_gas_used        bigint,
    receipt_blob_gas_price       varchar(255),
    block_number                 bigint,
    block_timestamp              bigint,
    block_gas_used               bigint,
    block_gas_limit              bigint,
    block_excess_blob_gas        bigint,
    block_blob_gas_used          bigint,
    block_base_fee               varchar(255)
);

-- 0028-add_mint.sql

ALTER TABLE block_explorer."Transactions"
    ADD mint varchar NULL;

ALTER TABLE block_explorer."Transactions"
    ADD withdrawal_value varchar NULL;

ALTER TABLE block_explorer."Transactions"
    DROP message_nonce;

ALTER TABLE block_explorer."Transactions"
    DROP messenger_version;

CREATE UNIQUE INDEX transaction_crosschain_unique_nonce_tx_type
    ON block_explorer."Transactions" (withdrawal_hash, tx_type)
    WHERE withdrawal_hash IS NOT NULL;

-- 0029-add-l1-deposit-gas-fee-burned
ALTER TABLE block_explorer."Transactions"
    ADD l1_deposit_burnt_fee varchar NULL;
