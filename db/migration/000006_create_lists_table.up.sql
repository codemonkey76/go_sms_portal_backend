CREATE TABLE IF NOT EXISTS "lists" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,

  -- Foreign keys
  "customer_id" bigint NOT NULL,
  
  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("customer_id") REFERENCES "customers" ("id")
);

