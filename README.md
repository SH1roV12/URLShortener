## URL Shortener in Go + Gin + MySQL + NanoID + Goose
A simple service for linking to links written in
Gin, a MySql database, and a library for creating IDs for short links - nanoid

## Features
- Accepts short URLs and direct URL codes.
- Store matching URLs in the database.
- All responses in JSON format
- Redirect short links to the original.
- Clean architecture + DDD

## How it works
1. Upon entry, the original link in JSON format is accepted.
Example: {"url":"www.youtube.com"}
2. The database stores the values ​​of short links and original links.
3. To go to the original site, go to the video at http://localhost:8080/generated id

## Launch
1. Clone the repository
```
git clone https://github.com/SH1roV12/urlshortener.git

```
2.
```
docker compose --build
```

The server will start at http://localhost:8080

You can optionally change the server parameters by changing the environment variables in docker-compose.yaml:
```
SERVICE_PORT=8080 - the server port
LENGTH_GEN=8 - generate short link lengths (id)
ALPHABET_GEN=(for example)ABCDabcd1234 - the alphabet that the id will consist of (it is not recommended to use special characters such as !?$, etc.)
DB_NAME=link_shortener - the database name (must match the name in the MYSQL_DATABASE variable)
DB_PORT=3306 - The default MySQL port is 3306
DB_HOST=database - The database host; in the case of Docker Compose, the service (database) name
DB_PASS=1212!!AaLXX - The database password (must match the password in the MYSQL_ROOT_PASSWORD variable)
DB_USER=root - The username that will be used to connect
```
