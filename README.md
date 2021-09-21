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
docker container ls
docker cp /home/mikolaj/Projects/Go-API/init.sql <container_id>:/home/init.sql
docker exec -it <container_id> /bin/bash
psql postgres://root:P%40ssw0rd@localhost
\l # get all databases
CREATE DATABASE go_movies;
\c go_movies # connect to database
\dt # show all databases in current db
\i /home/init.sql # execute sql script
\q # quit

\c go_movies
insert into movies_genres (movie_id, genre_id, created_at, updated_at) values (1, 1, '2021-05-19', '2021-05-19');
insert into movies_genres (movie_id, genre_id, created_at, updated_at) values (1, 2, '2021-05-19', '2021-05-19');
```
