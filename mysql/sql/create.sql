CREATE DATABASE IF NOT EXISTS heysay;
USE heysay;

CREATE TABLE IF NOT EXISTS users(
  id         VARCHAR(26)  PRIMARY KEY,
  name       VARCHAR(40)  NOT NULL,
  email      VARCHAR(256) NOT NULL,
  password   VARCHAR(100) NOT NULL,
  created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME     NULL,
);
