-- Setup user & database
CREATE ROLE BovespaWatch superuser;

GRANT ALL ON ALL TABLES IN SCHEMA public TO BovespaWatch;
ALTER ROLE BovespaWatch WITH LOGIN;

CREATE DATABASE BovespaWatch WITH OWNER BovespaWatch;

-- Setup tables
CREATE TABLE traders (
    id            SERIAL PRIMARY KEY,
    fullName      CHAR(30) NOT NULL,
    emailAddress  CHAR(30) NOT NULL,
    phoneNumber   CHAR(30) NOT NULL,
    monthlyQuota  INT NOT NULL DEFAULT 30,
    createdAt     TIMESTAMP DEFAULT (now() at TIME zone 'America/Recife')
);

CREATE TABLE watch_orders (
    id           SERIAL PRIMARY KEY,
    stockCode    CHAR(14) NOT NULL,
    targetPrice  DECIMAL NOT NULL,
    traderId     INT NOT NULL,
    active       BOOLEAN DEFAULT true NOT NULL,
    createdAt    TIMESTAMP DEFAULT (now() at TIME zone 'America/Recife')
);

CREATE TABLE messages (
    id         SERIAL PRIMARY KEY,
    traderId   INT NOT NULL,
    stockCode  CHAR(14) NOT NULL,
    fullText   CHAR(120) NOT NULL,
    createdAt  TIMESTAMP DEFAULT (now() at TIME zone 'America/Recife')
);

-- Seeds

INSERT INTO watch_orders (stockCode, targetPrice, traderId) VALUES
('CSNA3', 10.57, 1),
('GGBR4', 7.24, 1);

INSERT INTO traders (fullName, emailAddress, phoneNumber, monthlyQuota) VALUES
('Rodrigo Alves', 'rodrigoalvesv94@icloud.com', '+5581982599865', 200);
