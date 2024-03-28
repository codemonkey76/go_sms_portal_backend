CREATE TABLE IF NOT EXISTS "customer_user" (
  "id" bigserial PRIMARY KEY,

  -- Foreign keys
  "customer_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("customer_id") REFERENCES "customers" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  UNIQUE ("customer_id", "user_id")
);

