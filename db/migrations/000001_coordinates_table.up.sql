CREATE TABLE IF NOT EXISTS coordinates (
    id varchar(36) primary key,
    location varchar(255) not null,
    latitude varchar(255) not null,
    longitude varchar(255) not null
);