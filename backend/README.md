# Installation
Before running the Go server locally, you need to download Go on your OS.

On Ubuntu you easilly install Go using:
```sh
$ sudo snap install go --classic
```

Windows and MacOS instructions can be found [here](https://golang.org/doc/install).

After installing Go, check its version:
```sh
$ go version
1.15.3 linux/amd64
```
# Run the server
First cd to the main directory:
```sh
$ cd news-sentiment/backend
```

Run the server:
```sh
$ go run main.go
```

Check the server is running by pinging it:
```sh
$ curl localhost:8090/hello
Hello world!
```

Done!

# To get Docker running 🐳
Install [Docker Desktop](https://docs.docker.com/get-docker/). Docker compose is already included in the Docker Desktop installation. Check that docker is working by running the command below. You should get a message saying: "Hello from Docker!
This message shows that your installation appears to be working correctly."
```
docker run hello-world
```
Run this command to pull the latest image of mongo
```
docker pull mongo:latest
```
Run this command to see container running
```
docker exec -it mongo-container bash
```
Login into mongoDB
```
mongo -u root -p rootpassword 
```
After changes to docker-compose.yml run
```
docker-compose up
```
Login as admin user
```
docker-compose exec mongoDb mongo -u admin -p admin --authenticationDatabase "admin"
```
Make new db
```
use <new-database-name>
```
Monitering link: https://cloud.mongodb.com/freemonitoring/cluster/AAELUOCX3UZ6CRTLTDAGF62V7ND2KWBD
Useful link: https://start.jcolemorrison.com/docker-fresh-mysql-or-mongodb-instances-in-projects/

