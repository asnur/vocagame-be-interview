CREATE TYPE transaction_type AS ENUM (
    'deposit',
    'withdrawal',
    'transfer',
    'payment'
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    trx_id VARCHAR(50) NOT NULL UNIQUE,
    wallet_id INTEGER NOT NULL REFERENCES wallets(id) ON DELETE CASCADE,
    currency_id INTEGER NOT NULL REFERENCES currencies(id) ON DELETE CASCADE,
    type transaction_type NOT NULL,
    amount DECIMAL(20, 2) NOT NULL,
    description TEXT,
    refrence_wallet_id INTEGER REFERENCES wallets(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_transactions_trx_id ON transactions(trx_id);