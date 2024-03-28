CREATE TABLE IF NOT EXISTS "permission_role" (
  "id" bigserial PRIMARY KEY,

  -- Foreign keys
  "permission_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id"),
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id"),
  UNIQUE ("permission_id", "role_id")
);

