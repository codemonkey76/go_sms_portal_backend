CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "email_verified_at" timestamp,
  "password" varchar NOT NULL,
  "active" boolean NOT NULL DEFAULT false,
  "avatar_url" varchar,
  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp
);
