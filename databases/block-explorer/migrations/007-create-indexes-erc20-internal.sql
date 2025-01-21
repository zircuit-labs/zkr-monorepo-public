CREATE INDEX idx_erc20_token_addr ON block_explorer."ERC20Tokens" (addr);
CREATE INDEX idx_erc20_transfer_from_addr_lower ON block_explorer."ERC20Transfers" (lower(from_addr));
CREATE INDEX idx_erc20_transfer_l2_tx_hash ON block_explorer."ERC20Transfers" (l2_tx_hash);
CREATE INDEX idx_erc20_transfer_l2_tx_hash_lower ON block_explorer."ERC20Transfers" (lower(l2_tx_hash));
CREATE INDEX idx_erc20_transfer_timestamp_desc ON block_explorer."ERC20Transfers" (timestamp DESC);
CREATE INDEX idx_erc20_transfer_to_addr_lower ON block_explorer."ERC20Transfers" (lower(to_addr));
CREATE INDEX idx_erc20_transfer_to_from_timestamp_desc ON block_explorer."ERC20Transfers" (lower(to_addr), lower(from_addr), timestamp DESC);
CREATE INDEX idx_erc20_transfer_token_addr ON block_explorer."ERC20Transfers" (token_addr);
CREATE INDEX idx_internal_transactions_receiver_lower ON block_explorer."InternalTransactions" (lower(receiver));
CREATE INDEX idx_internal_transactions_sender_lower ON block_explorer."InternalTransactions" (lower(sender));
CREATE INDEX idx_internal_transactions_sender_receiver_lower_timestamp_desc ON block_explorer."InternalTransactions" (lower(sender), lower(receiver), timestamp DESC);
CREATE INDEX idx_transactions_to_from_origin_lower ON block_explorer."Transactions" (lower(to_addr), lower(from_addr), lower(l1_tx_origin));

ANALYZE block_explorer."ERC20Tokens";
ANALYZE block_explorer."ERC20Transfers";
ANALYZE block_explorer."InternalTransactions";

-- If migrating an existing database you might want to also vacuum the tables:
-- VACUUM ANALYZE block_explorer."ERC20Tokens";
-- VACUUM ANALYZE block_explorer."ERC20Transfers";
-- VACUUM ANALYZE block_explorer."InternalTransactions";
