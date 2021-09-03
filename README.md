# Go-API
#### Get modules
```sh
go mod init server
go get -u github.com/julienschmidt/httprouter
go get -u github.com/lib/pq@v1.10.
```
#### Run
Unix
```sh
go run app/*.go
```
Windows
```sh
go run ./app/.
```
### DB
```
docker exec -it <container_id> /bin/bash
psql postgres://root:P%40ssw0rd@localhost
\l # get all databases
CREATE DATABASE go_movies;
```