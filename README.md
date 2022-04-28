# afa7789 GoFiber WebSite
## Golang Website Template

A website created that uses go server with the package [go fiber](https://gofiber.io/).

It's a small company website that have blogpost and a contact page.
It is manly used as forefront to freelance projects and contracts.

### Mysql Setup

running the mysql in docker.
```sh
# run docker
docker run --name mysql-docker-local -e MYSQL_ROOT_PASSWORD=Password -d mysql:latest
# log on it and run the other codes bellow
docker exec -it mysql-docker-local mysql -u root -pPassword
```

Creating the user
```sql
CREATE USER 'site'@'%' IDENTIFIED BY 'PASSWORD' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'site'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
exit
```

Add or drop table of posts
```
DROP TABLE IF EXISTS WHITELIST;
CREATE table WHITELIST(
    id serial PRIMARY KEY,
    address VARCHAR (42) NOT NULL UNIQUE
    VARCHAR(max)
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

docker run --name mysql-docker-local -eMYSQL_ROOT_PASSWORD=Password -d mysql:latest
