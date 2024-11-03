-- create database
CREATE DATABASE "go_ecom_db";

-- create admin table
CREATE TABLE "admin" (
    "id" SERIAL PRIMARY KEY,
    "full_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL,
    "is_blocked" BOOLEAN DEFAULT FALSE,
    "is_verified" BOOLEAN DEFAULT FALSE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);