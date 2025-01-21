CREATE SCHEMA IF NOT EXISTS proof_state;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS proof_state.batches(
  "hash" CITEXT PRIMARY KEY,
  "finalized" BOOLEAN NOT NULL DEFAULT FALSE,
  "size" INT NOT NULL DEFAULT 0,
  "proof" JSONB NULL,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  "published_at" TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS proof_state.blocks(
  "hash" CITEXT PRIMARY KEY,
  "parent_hash" CITEXT NOT NULL,
  "number" INT NOT NULL CONSTRAINT number_unsigned CHECK (number >= 0),
  "proof" JSONB NULL,
  "json_path" TEXT NULL,
  "batch_hash" CITEXT,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  CONSTRAINT fk_block_batchhash FOREIGN KEY (batch_hash) REFERENCES proof_state.batches(hash) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS proof_state.component_proofs(
  "block_hash" CITEXT NOT NULL,
  "type" CITEXT NOT NULL,
  "proof" JSONB NULL,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  PRIMARY KEY ("block_hash", "type"),
  CONSTRAINT fk_component_batchhash FOREIGN KEY (block_hash) REFERENCES proof_state.blocks(hash) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS proof_state.locks(
  "id" TEXT NOT NULL PRIMARY KEY,
  "max_capacity" INTEGER NOT NULL CHECK (max_capacity >= 0),
  "capacity" INTEGER NOT NULL CHECK (capacity >= 0),
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  "updated_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  CHECK (capacity <= max_capacity)
);

CREATE TABLE IF NOT EXISTS proof_state.lock_holders(
  "instance_id" TEXT NOT NULL,
  "lock_id" TEXT NOT NULL,
  "service_name" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  "expires_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  "weight" INTEGER NOT NULL CHECK (weight >= 0),
  PRIMARY KEY ("instance_id", "lock_id")
);

-- 0001-blocks-number-index.sql
-- To speed up the query for the last published block
CREATE INDEX blocks_number_index ON proof_state.blocks(number DESC);

-- 0002-batch-hash-index.sql
-- To speed up the query for updating batch size
CREATE INDEX batch_hash_index ON proof_state.blocks(batch_hash);

-- 0003-block-hash-index.sql
-- To speed up the query for deleting component proofs
CREATE INDEX block_hash_index ON proof_state.component_proofs(block_hash);

-- 0004-update-component_proofs.sql
-- Add the nonce and prover version, and update the primary key to include them
-- Changing a PK requires dropping then adding it again
ALTER TABLE proof_state.component_proofs ADD "nonce" INT NOT NULL DEFAULT 0 CONSTRAINT nonce_unsigned CHECK (nonce >= 0);
ALTER TABLE proof_state.component_proofs ADD "prover_version" CITEXT DEFAULT '' NOT NULL;
ALTER TABLE proof_state.component_proofs DROP CONSTRAINT component_proofs_pkey;
ALTER TABLE proof_state.component_proofs ADD PRIMARY KEY ("block_hash", "type", "prover_version", "nonce");

-- 0005-batch-number.sql
-- Add a batch number to every batch. Although tempting to make this auto-generated, we want sequential only to a point:
-- batches should have continuous numbers (except for those that came before this was added which may remain at 0)
ALTER TABLE proof_state.batches ADD "number" INT NOT NULL DEFAULT 0 CONSTRAINT number_unsigned CHECK (number >= 0);

-- 0006-block-and-batch-nonce.sql
-- Add a nonce to the batches and blocks tables. Each item is still unique per hash, the nonce is used
-- only to indicate which proof version should be accepted.
ALTER TABLE proof_state.blocks ADD "nonce" INT NOT NULL DEFAULT 0 CONSTRAINT nonce_unsigned CHECK (nonce >= 0);
ALTER TABLE proof_state.batches ADD "nonce" INT NOT NULL DEFAULT 0 CONSTRAINT nonce_unsigned CHECK (nonce >= 0);

-- 0007-create-block-proofs-table.sql
-- Create a separate table for block proofs
CREATE TABLE IF NOT EXISTS proof_state.block_proofs(
  "block_hash" CITEXT NOT NULL,
  "prover_version" CITEXT NOT NULL,
  "nonce" INT NOT NULL CONSTRAINT nonce_unsigned CHECK (nonce >= 0),
  "proof" JSONB NULL,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  PRIMARY KEY ("block_hash", "prover_version", "nonce"),
  CONSTRAINT fk_blockhash FOREIGN KEY (block_hash) REFERENCES proof_state.blocks(hash) ON DELETE CASCADE
);

-- 0008-create-batch-proofs-table.sql
-- Create a separate table for batch proofs
CREATE TABLE IF NOT EXISTS proof_state.batch_proofs(
  "batch_hash" CITEXT NOT NULL,
  "prover_version" CITEXT NOT NULL,
  "nonce" INT NOT NULL CONSTRAINT nonce_unsigned CHECK (nonce >= 0),
  "proof" JSONB NULL,
  "created_at" TIMESTAMP DEFAULT clock_timestamp() NOT NULL,
  PRIMARY KEY ("batch_hash", "prover_version", "nonce"),
  CONSTRAINT fk_batchhash FOREIGN KEY (batch_hash) REFERENCES proof_state.batches(hash) ON DELETE CASCADE
);

-- 0011-add-status-and-type-to-batches.sql
ALTER TABLE proof_state.batches ADD "status" CITEXT DEFAULT '' NOT NULL;
ALTER TABLE proof_state.batches ADD "type" CITEXT DEFAULT '' NOT NULL;

-- 0012-batch-number-index.sql
-- To speed up the query for recent batches
CREATE INDEX batch_number_index ON proof_state.blocks(number);

-- 0013-batch-prover-version.sql
-- Add the prover version to the batches table
-- This represents the currently active prover version being worked only
ALTER TABLE proof_state.batches ADD "prover_version" CITEXT DEFAULT '' NOT NULL
