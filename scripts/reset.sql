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
    description TINYTEXT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);