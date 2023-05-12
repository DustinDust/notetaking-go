CREATE TYPE "note_format" AS ENUM (
  'html',
  'md',
  'plain_text'
);

CREATE TYPE "item_status" AS ENUM (
  'deleted',
  'archive',
  'active'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" text NOT NULL,
  "hash_password" text NOT NULL,
  "hash_refreshtoken" text NOT NULL,
  "status" item_status NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "user_info" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer UNIQUE NOT NULL,
  "firstname" text NOT NULL,
  "lastname" text,
  "profile_pic" text,
  "dateOfBirth" text,
  "bio" text
);

CREATE TABLE "projects" (
  "id" SERIAL PRIMARY KEY,
  "name" text NOT NULL,
  "description" text,
  "status" item_status NOT NULL,
  "owner_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "notes" (
  "id" SERIAL PRIMARY KEY,
  "title" text NOT NULL,
  "description" text,
  "format" note_format NOT NULL,
  "status" item_status NOT NULL,
  "user_id" int NOT NULL,
  "project_id" int NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "deleted_at" timestamp
);

ALTER TABLE "projects" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("id") REFERENCES "projects" ("id");

ALTER TABLE "notes" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("id")  REFERENCES "user_info" ("user_id");