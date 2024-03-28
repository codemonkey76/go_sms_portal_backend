CREATE TABLE IF NOT EXISTS "roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);

