CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "authors" ("name");

ALTER TABLE "books" ADD COLUMN "author_id" bigserial;
ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");
