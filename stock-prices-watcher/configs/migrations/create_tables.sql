CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE stocks CASCADE;
CREATE TABLE stocks (
    id UUID DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL UNIQUE,
    CONSTRAINT "pk_stocks_id" PRIMARY KEY ("id")
);

DROP TABLE stock_prices;
CREATE TABLE stock_prices (
    id UUID DEFAULT uuid_generate_v4(),
    stocks_id UUID NOT NULL,
    open REAL,
    high REAL,
    low REAL,
    close REAL,
    volume REAL,
    time_series TIMESTAMP,
    CONSTRAINT "pk_stock_prices_id" PRIMARY KEY ("id"),
    CONSTRAINT "fk_stocks_id" FOREIGN KEY ("stocks_id") REFERENCES stocks ("id")
    UNIQUE(time_series, stocks_id)
);
