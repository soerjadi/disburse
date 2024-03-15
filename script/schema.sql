CREATE TABLE IF NOT EXISTS "account" (
    "id" VARCHAR PRIMARY KEY NOT NULL,
    "name" VARCHAR NOT NULL,
    "number" VARCHAR NOT NULL,
    "origin_bank" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW()
);

CREATE INDEX account_number_index ON account USING btree(number);
CREATE INDEX account_name_index ON account USING btree(name);

CREATE TABLE IF NOT EXISTS "transaction" (
    "id" BIGSERIAL PRIMARY KEY NOT NULL,
    "trx_id" VARCHAR NULL,
    "destination_id" VARCHAR NOT NULL,
    "amount" BIGINT NOT NULL,
    "unique_number" BIGINT NOT NULL,
    "type" VARCHAR NOT NULL,  
    "status" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT NOW()
);