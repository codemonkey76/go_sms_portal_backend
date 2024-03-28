CREATE TABLE IF NOT EXISTS "customers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "sender_id" varchar NOT NULL,
  "active" boolean NOT NULL DEFAULT false,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);
