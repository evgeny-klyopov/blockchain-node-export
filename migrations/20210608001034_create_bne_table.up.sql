CREATE SEQUENCE bne_transactions_id_seq START 1;
DROP TABLE IF EXISTS "bne_transactions";
CREATE TABLE "bne_transactions" (
    "id" int4 NOT NULL DEFAULT nextval('bne_transactions_id_seq'::regclass),
    "address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "category" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "amount" numeric(10,8) NOT NULL,
    "fee" numeric(10,8) NOT NULL,
    "confirmations" int4 NOT NULL,
    "blocktime" int4 NOT NULL,
    "block_datetime" timestamp(6) NOT NULL,
    "tx_id" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "time" int4 NOT NULL,
    "datetime" timestamp(6) NOT NULL,
    "block_index" int4 NOT NULL,
    "block_hash" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
);
ALTER TABLE "bne_transactions" OWNER TO "postgres";

CREATE UNIQUE INDEX "bne_transactions_address_category_amount_fee_blocktime_tx_i_idx" ON "bne_transactions" USING btree (
  "address" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "category" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "amount" "pg_catalog"."numeric_ops" ASC NULLS LAST,
  "fee" "pg_catalog"."numeric_ops" ASC NULLS LAST,
  "blocktime" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "tx_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "time" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "block_index" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "block_hash" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

ALTER TABLE "bne_transactions" ADD CONSTRAINT "bne_transactions_pkey" PRIMARY KEY ("id");