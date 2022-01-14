create database if not exists blades;
use blades;
drop table if exists skates;
create table skates
(
    id            int auto_increment
        primary key,
    name          varchar(128)             not null,
    price         float                    not null,
    modified_date datetime                 not null,
    added_date    datetime                 not null,
    is_active     varchar(4) default 'yes' not null,
    units         int        default 0     not null
);
INSERT INTO blades.skates (id, name, price, modified_date, added_date, is_active, units)
VALUES (1, 'Bauer Supreme Ultrasonic Skates', 599.99, NOW(), NOW(), 'yes', 12);
INSERT INTO blades.skates (id, name, price, modified_date, added_date, is_active, units)
VALUES (2, 'Bauer Vapor 2X Pro Skates', 549.99, NOW(), NOW(), 'yes', 22);
INSERT INTO blades.skates (id, name, price, modified_date, added_date, is_active, units)
VALUES (3, 'CCM Jetspeed FT2 Skates', 549.99, NOW(), NOW(), 'yes', 23);
INSERT INTO blades.skates (id, name, price, modified_date, added_date, is_active, units)
VALUES (4, 'TRUE Custom Skates', 499.99, NOW(), NOW(), 'yes', 54);
