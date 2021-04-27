CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "question" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "answers" varchar[] NOT NULL,
  "correct" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "group_id" bigserial NOT NULL,
  "test_id" bigserial
);

CREATE TABLE "group" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "test" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "test_queston" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigserial,
  "test_id" bigserial
);

ALTER TABLE "question" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "question" ADD FOREIGN KEY ("test_id") REFERENCES "test" ("id");

ALTER TABLE "test_queston" ADD FOREIGN KEY ("question_id") REFERENCES "question" ("id");

ALTER TABLE "test_queston" ADD FOREIGN KEY ("test_id") REFERENCES "test" ("id");

CREATE INDEX ON "question" ("name");

CREATE INDEX ON "question" ("id");

CREATE INDEX ON "test" ("name");