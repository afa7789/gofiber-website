![Theme Image](resources/banner.png)
# afa7789 GoFiber WebSite

A Golang Website Template for developers made from the scracth using go, html , css and a little bit of js with ajax requests. It's a small company website that have blogpost and a contact page, and a profile github page rendered from the README.md profile page. It can be used as forefront to freelance projects and contracts for hackers ( developers who wants to craft projects for other persons)

This website heavily uses the go server in this package [go fiber](https://gofiber.io/). And I trully recommend it.

### Settuping the project

__Clone__

`git clone https://github.com/afa7789/gofiber-website.git && cd gofiber-website`
#### Mysql Setup

__Creating & running the mysql in docker.__
```sh
# create docker
docker create -v /var/lib/mysql --name mysqldata mysql
docker run --name mysqldb_fiber_site --volumes-from mysqldata -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 -d mysql:latest
# restart
docker start mysqldb_fiber_site
# log on it and run the other codes bellow
sudo docker exec -it mysqldb_fiber_site mysql -u root -ppassword
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

### Running the project

__Run the database__

This has to be done after the database is [setupped](#mysql-setup).

`docker start mysqldb_fiber_site`

__Run the server__

`go run .`

### Make commands

__lint__
run the linter to check if the code is all good with the golang paterns.
`make lint`

__serve__
Serve the website to write the outputs to the log files.
`make serve`

__build__
Build the code to be used as a binary.
`make build`

__Running it__
runs the docker database and the project.
`make run`

## Features
- Contact Mailing & Failed and Thanks redirections
- Blog Post Edit and Create
- Blog Section, post view missing post and related ones
- Blog updating to last posts in front page.
- SLUG handling for better (SEO).
- Github README profile, reader page.

## Future Features
- Logs and prints are done to a file ( can be improved I am doing it with make serve as unix redirecting the output)
- Setup blog part as subdomain: https://github.com/gofiber/fiber/issues/750 use subdomain on blog
- ToDoList Page
- Organize html css images to use smaller ones to save loading time.
- Links Page 
- Links Page as subdomain

![Visitor Badge](https://visitor-badge.laobi.icu/badge?page_id=afa7789.gofiber-website)
