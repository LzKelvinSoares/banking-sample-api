/* ===========================================================
   DROP (ordem inversa das dependências por causa das FKs)
   =========================================================== */
DROP TABLE IF EXISTS refresh_token_store;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;
GO

/* ===========================================================
   CUSTOMERS
   =========================================================== */
CREATE TABLE customers (
    customer_id   INT IDENTITY(2000,1) NOT NULL,
    customername  NVARCHAR(100) NOT NULL,
    date_of_birth DATE          NOT NULL,
    city          VARCHAR(100)  NOT NULL,
    zipcode       VARCHAR(10)   NOT NULL,
    status        BIT           NOT NULL CONSTRAINT DF_customers_status DEFAULT 1,
    CONSTRAINT PK_customers PRIMARY KEY (customer_id)
);
GO

SET IDENTITY_INSERT customers ON;

INSERT INTO customers (customer_id, customername, date_of_birth, city, zipcode, status) VALUES
    (2000, N'Steve',  '1978-12-15', 'Delhi',           '110075', 1),
    (2001, N'Arian',  '1988-05-21', 'Newburgh, NY',    '12550',  1),
    (2002, N'Hadley', '1988-04-30', 'Englewood, NJ',   '07631',  1),
    (2003, N'Ben',    '1988-01-04', 'Manchester, NH',  '03102',  0),
    (2004, N'Nina',   '1988-05-14', 'Clarkston, MI',   '48348',  1),
    (2005, N'Osman',  '1988-11-08', 'Hyattsville, MD', '20782',  0);

SET IDENTITY_INSERT customers OFF;
GO

/* ===========================================================
   ACCOUNTS
   =========================================================== */
CREATE TABLE accounts (
    account_id   INT IDENTITY(95470,1) NOT NULL,
    customer_id  INT            NOT NULL,
    opening_date DATETIME2(0)   NOT NULL CONSTRAINT DF_accounts_opening_date DEFAULT GETDATE(),
    account_type VARCHAR(10)    NOT NULL,
    amount       DECIMAL(10,2)  NOT NULL,
    status       BIT            NOT NULL CONSTRAINT DF_accounts_status DEFAULT 1,
    CONSTRAINT PK_accounts PRIMARY KEY (account_id),
    CONSTRAINT accounts_FK FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
);
GO

CREATE INDEX accounts_FK_idx ON accounts (customer_id);
GO

SET IDENTITY_INSERT accounts ON;

INSERT INTO accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES
    (95470, 2000, '2020-08-22 10:20:06', 'saving',   6823.23, 1),
    (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, 1),
    (95472, 2001, '2020-08-09 10:35:22', 'saving',   7000,    1),
    (95473, 2001, '2020-08-09 10:38:22', 'saving',   5861.86, 1);

SET IDENTITY_INSERT accounts OFF;
GO

/* ===========================================================
   TRANSACTIONS
   =========================================================== */
CREATE TABLE transactions (
    transaction_id   INT IDENTITY(1,1) NOT NULL,
    account_id       INT            NOT NULL,
    amount           DECIMAL(10,2)  NOT NULL,
    transaction_type VARCHAR(10)    NOT NULL,
    transaction_date DATETIME2(0)   NOT NULL CONSTRAINT DF_transactions_date DEFAULT GETDATE(),
    CONSTRAINT PK_transactions PRIMARY KEY (transaction_id),
    CONSTRAINT transactions_FK FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);
GO

CREATE INDEX transactions_FK_idx ON transactions (account_id);
GO

/* ===========================================================
   USERS
   =========================================================== */
CREATE TABLE users (
    username    VARCHAR(20)  NOT NULL,
    password    VARCHAR(20)  NOT NULL,
    role        VARCHAR(20)  NOT NULL,
    customer_id INT          NULL,
    created_on  DATETIME2(0) NOT NULL CONSTRAINT DF_users_created_on DEFAULT GETDATE(),
    CONSTRAINT PK_users PRIMARY KEY (username)
);
GO

INSERT INTO users (username, password, role, customer_id, created_on) VALUES
    ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
    ('2001',  'abc123', 'user',  2001, '2020-08-09 10:27:22'),
    ('2000',  'abc123', 'user',  2000, '2020-08-09 10:27:22');
GO

/* ===========================================================
   REFRESH_TOKEN_STORE
   =========================================================== */
CREATE TABLE refresh_token_store (
    refresh_token VARCHAR(300) NOT NULL,
    created_on    DATETIME2(0) NOT NULL CONSTRAINT DF_refresh_token_created_on DEFAULT GETDATE(),
    CONSTRAINT PK_refresh_token_store PRIMARY KEY (refresh_token)
);
GO