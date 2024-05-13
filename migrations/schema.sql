-- Version: 1.1
-- Description: Create table users

CREATE TABLE users (
	id UUID DEFAULT gen_random_uuid () NOT NULL,
	first_name VARCHAR(255) NULL,
	last_name VARCHAR(255) NULL,
	username VARCHAR(255) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	date_created TIMESTAMP,
	date_updated TIMESTAMP,
	PRIMARY KEY (id)
);

CREATE TABLE accounts (
	id UUID DEFAULT gen_random_uuid () NOT NULL,
	user_id UUID NOT NULL,
	name VARCHAR(255) NOT NULL,
	available_amount DECIMAL (10,
		2) NOT NULL,
	TYPE VARCHAR(100),
	slug VARCHAR(255),
	date_created TIMESTAMP,
	date_updated TIMESTAMP,
	PRIMARY KEY (id)
);

-- Version: 1.3
-- Description: Create table transactions
CREATE TABLE transactions (
	id            UUID default gen_random_uuid() NOT NULL,
	from_account_id UUID,
	to_account_id UUID,
	account_id UUID,
	user_id UUID,
	to_user_id UUID,
	amount DECIMAL (10,
		2),
	name VARCHAR(255),
	slug VARCHAR(255),
	TYPE VARCHAR(100),
    transaction_type VARCHAR(100) NULL DEFAULT '""',
	date_created TIMESTAMP,
	date_updated TIMESTAMP,
	PRIMARY KEY (id)
);
