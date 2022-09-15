DROP DATABASE IF EXISTS gofiber_website;
CREATE DATABASE gofiber_website;
USE gofiber_website;
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    slug VARCHAR(255),
    image VARCHAR(255),
    related_posts VARCHAR(255),
    synopsis TINYTEXT,
    content LONGTEXT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);
DROP TABLE IF EXISTS link;
CREATE TABLE links (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    href VARCHAR(255),
    image LONGTEXT,
    index_order INT UNIQUE,
    description TINYTEXT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER `link_trigger`
BEFORE INSERT ON `links` 
FOR EACH ROW SET NEW.index_order=(
    SELECT `AUTO_INCREMENT` FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA`=DATABASE() AND `TABLE_NAME`='links'
);


DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    subject VARCHAR(255),
    name VARCHAR(255),
    email VARCHAR(255),
    text LONGTEXT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);


