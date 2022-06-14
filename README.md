# afa7789 GoFiber WebSite
## Golang Website Template

A website created that uses go server with the package [go fiber](https://gofiber.io/).

It's a small company website that have blogpost and a contact page.
It is manly used as forefront to freelance projects and contracts.

### Running the project

__Clone__

`git clone https://github.com/afa7789/gofiber-website.git && cd gofiber-website`

__Run the database__

This has to be done after the database is [setupped](#mysql-setup).

`docker start mysqldb_fiber_site`

__Run the server__

`go run .`

__Using make__
`make run`

### Mysql Setup

__Creating & running the mysql in docker.__
```sh
# create docker
docker create -v /var/lib/mysql --name mysqldata mysql
docker run --name mysqldb_fiber_site --volumes-from mysqldata -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql:latest
# restart
docker start mysqldb_fiber_site
# log on it and run the other codes bellow
docker exec -it mysqldb_fiber_site mysql -u root -ppassword
```

__Creating the user to access the database.__
```sql
CREATE USER 'site'@'%' IDENTIFIED BY 'PASSWORD';
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
    slug VARCHAR(255),
    image VARCHAR(255),
    related_posts VARCHAR(255),
    synopsis TINYTEXT,
    content LONGTEXT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP
);
```

if using dbeaver, it is possible that you will need to change the permission in driver properties: allowPublicKeyRetrieval to true. If the permission doesn't exist, just add a newer one.

__Exporting DB on CLI__

```sh
docker exec -it mysqldb_fiber_site mysqldump -u root -ppassword gofiber_website > dump.sql
```

## Features
- Contact Mailing & Failed and Thanks redirections
- Blog Post Edit and Create
- Blog Section, post view missing post and related ones
- Blog updating to last posts in front page.
- SLUG handling for better (SEO).

## Future Features
- Logs and prints are done to a file
- Setup blog part as subdomain: https://github.com/gofiber/fiber/issues/750 use subdomain on blog
- ToDoList Page
- Organize html css images to use smaller ones to save loading time.
- Links Page 
- Links Page as subdomain
- Github README profile, reder page

![Visitor Badge](https://visitor-badge.laobi.icu/badge?page_id=afa7789.gofiber-website)
