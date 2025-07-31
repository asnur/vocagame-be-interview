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
  (1, 2, 1.1),               -- USD → EUR
  (2, 1, 0.9090909091),      -- EUR → USD

  (1, 3, 0.009),             -- USD → JPY
  (3, 1, 111.1111111111),    -- JPY → USD

  (1, 4, 0.0000645),         -- USD → IDR
  (4, 1, 15503.87596899),    -- IDR → USD

  (2, 3, 0.0081818),         -- EUR → JPY
  (3, 2, 122.25),            -- JPY → EUR

  (2, 4, 0.0000588),         -- EUR → IDR
  (4, 2, 17006.80272109),    -- IDR → EUR

  (3, 4, 0.000105),          -- JPY → IDR
  (4, 3, 9523.80952381);     -- IDR → JPY