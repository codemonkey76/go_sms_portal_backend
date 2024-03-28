CREATE TABLE IF NOT EXISTS "role_user" (
  "id" bigserial PRIMARY KEY,

  -- Foreign keys
  "role_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  UNIQUE ("role_id", "user_id")
);

