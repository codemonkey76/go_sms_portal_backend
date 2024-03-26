CREATE TABLE IF NOT EXISTS "contact_list" (
  "id" serial PRIMARY KEY,

  -- Foreign keys
  "contact_id" bigint NOT NULL,
  "list_id" bigint NOT NULL,

  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id"),
  FOREIGN KEY ("list_id") REFERENCES "lists" ("id"),
  UNIQUE ("contact_id", "list_id")
);

