CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "questions" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "answers" varchar[] NOT NULL,
  "correct" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "group_id" bigserial NOT NULL
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tests" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "test_questions" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigserial NOT NULL,
  "test_id" bigserial NOT NULL
);

ALTER TABLE "questions" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "test_questions" ADD FOREIGN KEY ("question_id") REFERENCES "questions" ("id");

ALTER TABLE "test_questions" ADD FOREIGN KEY ("test_id") REFERENCES "tests" ("id");

CREATE INDEX ON "questions" ("name");

CREATE INDEX ON "questions" ("id");

CREATE INDEX ON "tests" ("name");