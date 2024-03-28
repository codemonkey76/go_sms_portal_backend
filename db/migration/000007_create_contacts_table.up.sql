CREATE TABLE IF NOT EXISTS "contacts" (
  "id" bigserial PRIMARY KEY,
  "phone" varchar NOT NULL,
  "first_name" varchar,
  "last_name" varchar,
  "company" varchar,

  -- Foreign keys
  "customer_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("customer_id") REFERENCES "customers" ("id")
);

