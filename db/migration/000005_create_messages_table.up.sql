CREATE TABLE IF NOT EXISTS "messages" (
  "id" bigserial PRIMARY KEY,
  "body" varchar(1000) NOT NULL,
  "segments" int NOT NULL DEFAULT 1,
  "from" varchar(255),
  "to" varchar(255),
  "status" varchar(255),
  "sid" varchar(255),
  "archived" boolean NOT NULL DEFAULT false,
  "sender_id" bigint NOT NULL,

  -- Foreign keys
  "customer_id" bigint NOT NULL,

  -- Timestamps
  "sent_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,

  -- Constraints
  FOREIGN KEY ("customer_id") REFERENCES "customers" ("id")
);

