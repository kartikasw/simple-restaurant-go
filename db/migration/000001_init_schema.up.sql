CREATE TABLE "variants" (
  "id" bigserial PRIMARY KEY,
  "variant_name" varchar NOT NULL,
  "quantity" INT NOT NULL DEFAULT 0,
  "product_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "variants" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");