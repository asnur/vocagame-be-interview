CREATE TABLE exchange_rates (
    id SERIAL PRIMARY KEY,
    from_currency_id INTEGER NOT NULL REFERENCES currencies(id) ON DELETE CASCADE,
    to_currency_id INTEGER NOT NULL REFERENCES currencies(id) ON DELETE CASCADE,
    rate DECIMAL(20, 10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

INSERT INTO exchange_rates (from_currency_id, to_currency_id, rate) VALUES
  (1, 2, 1.1),               -- USD → EUR (1 USD = 1.1 EUR)
  (2, 1, 0.9090909091),      -- EUR → USD (1 EUR = 0.909 USD)

  (1, 3, 110.0),             -- USD → JPY (1 USD = 110 JPY)
  (3, 1, 0.0090909091),      -- JPY → USD (1 JPY = 0.009 USD)

  (1, 4, 15300.0),           -- USD → IDR (1 USD = 15,300 IDR)
  (4, 1, 0.0000653594771),   -- IDR → USD (1 IDR = 0.0000653 USD)

  (2, 3, 121.0),             -- EUR → JPY (1 EUR = 121 JPY)
  (3, 2, 0.0082644628),      -- JPY → EUR (1 JPY = 0.008264 EUR)

  (2, 4, 16830.0),           -- EUR → IDR (1 EUR = 16,830 IDR)
  (4, 2, 0.0000594184),      -- IDR → EUR (1 IDR = 0.0000594 EUR)

  (3, 4, 139.0909091),       -- JPY → IDR (1 JPY = 139.09 IDR)
  (4, 3, 0.0071942446);      -- IDR → JPY (1 IDR = 0.007194 JPY)