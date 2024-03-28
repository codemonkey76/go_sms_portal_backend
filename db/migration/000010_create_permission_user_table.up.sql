CREATE TABLE IF NOT EXISTS "permission_user" (
  "id" bigserial PRIMARY KEY,

  -- Foreign keys
  "permission_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  UNIQUE ("permission_id", "user_id")
);

