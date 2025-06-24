/* create database and tables */
SET GLOBAL tidb_multi_statement_mode='ON' 

CREATE DATABASE IF NOT EXISTS ztf_db;

USE ztf_db;

CREATE TABLE
  IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

CREATE TABLE
  IF NOT EXISTS merchants (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  );

/* seed data */
INSERT INTO
  users (id, username, email)
VALUES
  ('a1927cb1-1db0-4b18-91ed-578559ba7489', 'ntbinh', 'ntbinh243@gmail.com'),
  ('bb8a677a-c9f3-46ca-8299-54b23d2c4d23', 'dqtrieu', 'dqtrieu@gmail.com'),
  ('a7dacaa5-baa2-4458-b62e-27a5a073dfb1', 'hatra', 'hatra@gmail.com');

INSERT INTO
  merchants (id, username, email)
VALUES
  ('53c7e139-5c92-49e7-a4b2-667782e8fd9e', 'highland', 'merchant@highland.com'),
  ('1540bf4a-07d6-48b9-8047-726c9150cf1f', 'phuclong', 'merchant2@phuclong.com'),
  ('9b66ffae-02d1-48b0-94e3-937adf52f85a', 'thecoffeehouse', 'merchant3@thecoffeehouse.com');