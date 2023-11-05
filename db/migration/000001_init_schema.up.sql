CREATE TABLE  IF NOT EXISTS accounts (
    id bigserial PRIMARY KEY,
    owner_name varchar NOT NULL,
    balance bigint NOT NULL,
    display_picture varchar,
    currency varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT now()
);

CREATE INDEX ON accounts (owner_name);

CREATE TABLE  entries (
  id bigserial PRIMARY KEY,
  amount bigint NOT NULL,
  account_id bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts (id);

CREATE INDEX ON entries (account_id);

CREATE TABLE transfers (
  id bigserial PRIMARY KEY,
  from_account_id bigint NOT NULL ,
  to_account_id bigint NOT NULL,
  amount bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT now()
);

ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);

ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);

CREATE INDEX ON transfers (from_account_id);
CREATE INDEX ON transfers (to_account_id);
CREATE INDEX ON transfers (from_account_id, to_account_id);

COMMENT ON COLUMN transfers.amount IS 'Must not be null';
