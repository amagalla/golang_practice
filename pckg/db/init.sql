CREATE DATABASE IF NOT EXISTS sample_database;

USE sample_database;

CREATE TABLE IF NOT EXISTS profiles (
    id                  BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    first_name          VARCHAR(255) NOT NULL,
    last_name           VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);