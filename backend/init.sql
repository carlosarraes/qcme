CREATE SCHEMA IF NOT EXISTS data;

CREATE TABLE data.qrcode (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    linkedin VARCHAR(512),
    github VARCHAR(512)
);
