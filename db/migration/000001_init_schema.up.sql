CREATE TABLE "clients" (
    "id" SERIAL PRIMARY KEY,
    "client_name" VARCHAR(100) NOT NULL,
    "client_id" VARCHAR(50) NOT NULL UNIQUE,
    "client_password" VARCHAR(200) NOT NULL,
    "created_at" TIMESTAMP
);

CREATE TABLE "tokens" (
    "id" UUID PRIMARY KEY,
    "client_id" INT NOT NULL,
    "token" VARCHAR(200) NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "expires_at" TIMESTAMP NOT NULL
);

ALTER TABLE "tokens" ADD CONSTRAINT "fk_client" FOREIGN KEY ("client_id") REFERENCES "clients" ("id");