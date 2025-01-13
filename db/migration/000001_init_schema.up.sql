CREATE TABLE "clients" (
    "id" SERIAL PRIMARY KEY,
    "client_name" VARCHAR(100) NOT NULL,
    "client_id" VARCHAR(50) NOT NULL UNIQUE,
    "client_password" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);