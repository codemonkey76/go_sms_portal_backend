CREATE TABLE IF NOT EXISTS "sessions" (
  "id" varchar PRIMARY KEY,
  "ip_address" varchar  NULL,
  "user_agent" varchar  NULL,
  "payload" text NOT NULL,
  "last_activity" bigint NOT NULL,
  
  -- Foreign keys
 "user_id" bigint NOT NULL,
 
  -- Timestamps
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp,
  
  -- Constraints
  FOREIGN KEY ("user_id") REFERENCES "users" ("id") 
);
