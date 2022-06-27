DROP DATABASE IF EXISTS my_db;

CREATE DATABASE my_db;

USE my_db;

DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Items;
DROP TABLE IF EXISTS Transactions;
DROP TABLE IF EXISTS Vouchers;
DROP TABLE IF EXISTS Api;

CREATE TABLE Users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    phone VARCHAR(10) NOT NULL,
    name VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    points INT DEFAULT 0,
    last_login DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
INSERT INTO Users (phone, name, password, points) VALUES
("81234567", "Jasper", "$2a$04$cNkuX2mhC4YtEcvkogWPR.S7QpWwv5Txyhh9i7gZnDB7FrSRFohAK", 200);
-- added more entry for users
INSERT INTO Users (phone, name, password, points) VALUES
("82234567", "Quek Qi", "$2a$04$A8nk4nE32ViV64LgzxPw4.BBRERpe0Z4VQSzcTl0CxjmgLT3zrdDW", 120);
INSERT INTO Users (phone, name, password, points) VALUES
("91234567", "Dylan Kiew", "$2a$04$9Q.Bx4c1kSorj8tmQ/lLmum3eG0ZX/Xkr8zfAjsiFHBe59BBNBtai", 100);
INSERT INTO Users (phone, name, password, points) VALUES
("92234567", "William Neo", "$2a$04$73GZOU7d6I.7dI6DxV8jSOhe73z51ONH0dcDpzXmnvqkQdVkDmG8O", 150);

CREATE TABLE Items (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(10) NOT NULL
);
INSERT INTO Items (name) VALUES ("Paper");
INSERT INTO Items (name) VALUES ("Plastic");
INSERT INTO Items (name) VALUES ("Glass");
INSERT INTO Items (name) VALUES ("Metal");

CREATE TABLE Transactions (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    item_id int NOT NULL,
    weight int NOT NULL,
    trans_date DATETIME DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (1, 2, 500);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (1, 3, 760);
-- added more entry for transactions
INSERT INTO Transactions (user_id, item_id, weight) VALUES (1, 2, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (1, 3, 460);
-- user 2
INSERT INTO Transactions (user_id, item_id, weight) VALUES (2, 1, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (2, 4, 460);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (2, 1, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (2, 4, 460);
-- user 3
INSERT INTO Transactions (user_id, item_id, weight) VALUES (3, 4, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (3, 3, 460);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (3, 4, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (3, 3, 460);
-- user 4
INSERT INTO Transactions (user_id, item_id, weight) VALUES (4, 1, 200);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (4, 1, 460);
INSERT INTO Transactions (user_id, item_id, weight) VALUES (4, 1, 300);

CREATE TABLE Vouchers (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    voucher_amt int NOT NULL,
    voucher_id VARCHAR(100) NOT NULL,
    redeem int DEFAULT 0,
    time_updated DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- added entry for Vouchers
INSERT INTO Vouchers (user_id, voucher_amt, voucher_id, redeem) VALUES (1, 20, "0fd5062eccc4b049f0ba75ca31db3a4c53253bb9fe3addf3332sfgd81307aa2", 1);
INSERT INTO Vouchers (user_id, voucher_amt, voucher_id, redeem) VALUES (1, 20, "0fd5062eccc4b049f0ba75ca31db3a4c53253bb9fe3addf3332sfgd81307bbb", 1);

CREATE TABLE Api (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(10) NOT NULL,
    api_key VARCHAR(100) NOT NULL
);
INSERT INTO Api (username, api_key) VALUES
("admin", "0fd5062eccc4b049f0ba75ca31db3a4cb12088bb9fe3addf33e9e2b481307aa2");