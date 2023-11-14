# fandom-cookbook-server
The backend for the [Fandom Cookbook](https://github.com/Lucky4rever/fandom-cookbook) project

## Docker and MongoDB
1. To build the image run:
```
make db_build
```
or:
```
docker build -f Dockerfile -t fandom-cookbook/mongodb .
```
2. To start the database run:
```
make docker_run
```
or:
```
docker run -d -p 27017:27017 --name fandom-cookbook_mongodb fandom-cookbook/mongodb
```
4. To restore the db run:
```
docker exec -it fandom-cookbook_mongodb bash
```
```
mongorestore dbbackup/
```
## Server
To run the backend:
```
go run main.go
```
