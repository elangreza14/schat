CREATE TABLE IF NOT EXISTS "users" (
  "id" CHAR(26) PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL,
  "email" VARCHAR(50) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ,
  "deleted_at" TIMESTAMPTZ
);

---- create above / drop below ----
DROP TABLE "users";
