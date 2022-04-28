# afa7789 GoFiber WebSite
## Golang Website Template

A website created that uses go server with the package [go fiber](https://gofiber.io/).

It's a small company website that have blogpost and a contact page.
It is manly used as forefront to freelance projects and contracts.

### Mysql Setup

__Creating & running the mysql in docker.__
```sh
# create docker
docker create -v /var/lib/mysql --name mysqldata mysql
docker run --name mysqldb_fiber_site --volumes-from mysqldata -e MYSQL_ROOT_PASSWORD=password -p 3307:3306 -d mysql:latest
# restart
docker start mysqldb_fiber_site
# log on it and run the other codes bellow
docker exec -it mysqldb_fiber_site mysql -u root -ppassword
```

__Creating the user to access the database.__
```sql
CREATE USER 'site'@'%' IDENTIFIED BY 'PASSWORD' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'site'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
exit
```

__Add or drop database and table of posts.__
```
DROP DATABASE [IF EXISTS] gofiber_website;
CREATE DATABASE gofiber_website;
USE gofiber_website;
DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    image VARCHAR(255),
    related_posts VARCHAR(255),
    synopsis TINYTEXT,
    content LONGTEXT
);
```

## Features
- Contact Mailing & Failed and Tahnks redirections
- Blog Post Edit and Create
- Blog Section ,  missing post and individual one
  
## Missing Features
- Log remove panic and log to files.
- Setup blog part as subdomain.
- ToDoList Page
- https://github.com/gofiber/fiber/issues/750 use subdomain on blog
