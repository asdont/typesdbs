# Database field types

---

docker-compose.yml

or

```shell
docker run -d --name oracle \
  -p 1522:1521 \
  -v $(pwd)/data/oracle:/data \
  container-registry.oracle.com/database/free:latest
```

```shell
docker run -d -p 27017:27017 --name mongodb \
  -e MONGO_INITDB_ROOT_USERNAME=test \
  -e MONGO_INITDB_ROOT_PASSWORD=test \
  mongo
```

```shell
docker run -d --name clickhouse \
  -p 8123:8123 -p 9000:9000 \
  -v $(pwd)/data/clickhouse:/var/lib/clickhouse \
  yandex/clickhouse-server:latest
```

```shell
docker run -d --name postgres \
  -p 5433:5432 \
  -v $(pwd)/data/postgres:/var/lib/postgresql/data \
  -e POSTGRES_USER=test \
  -e POSTGRES_PASSWORD=test \
  -e POSTGRES_DB=testdb \
  postgres:latest
```

```shell
docker run -d --name mysql \
  -p 3307:3306 \
  -v $(pwd)/data/mysql:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=test \
  -e MYSQL_USER=test \
  -e MYSQL_PASSWORD=test \
  -e MYSQL_DATABASE=testdb \
  mysql:latest
```

```shell
docker run -d --name mssql \
  -p 1434:1433 \
  -v $(pwd)/data/mssql:/var/opt/mssql \
  -e ACCEPT_EULA=Y \
  -e MSSQL_SA_PASSWORD='Test12345!' \
  -e MSSQL_PID=Developer \
  mcr.microsoft.com/mssql/server:2022-latest
```
