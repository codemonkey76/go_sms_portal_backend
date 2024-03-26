CREATE TABLE IF NOT EXISTS "permissions" (
  "id" serial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);
