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
$ make run
```

Check the server is running by pinging it:
```sh
$ curl localhost:8090/hello
Hello world!
```

Done!

# To get Docker running 🐳
### Make sure to change the example.env to .env after you have docker running. 
Install [Docker Desktop](https://docs.docker.com/get-docker/). Docker compose is already included in the Docker Desktop installation. Check that docker is working by running the command below. You should get a message saying: "Hello from Docker!
This message shows that your installation appears to be working correctly."
```
docker run hello-world
```
Go to the backend folder
```
cd backend
```
Run this command to pull the latest image of mongo
```
docker pull mongo:latest
```
After changes to docker-compose.yml/ starting up docker run:
```
docker-compose up 
```
Run this command to see container running/enter bash
```
docker exec -it mongo-container bash
```
Login into mongoDB
```
mongo -u <username from .env> -p <password from .env> 
```
Make new admin user
```
db.createUser(  
  {
    user: <username>,
    pwd: <pwd>,
    roles: [ { role: "root", db: "admin" } ]
  }
)
```
Login as admin user
```
docker-compose exec mongoDb mongo -u <username> -p <password> --authenticationDatabase <db you made when making an admin acc>
```
Check ports and volume:
```
docker ps
docker volume ls
```
Make new db
```
use <new-database-name>
```
Check all your db
```
db
```
Add things to db
```
db.<db name>.insert({})
```
- Monitering link: https://cloud.mongodb.com/freemonitoring/cluster/AAELUOCX3UZ6CRTLTDAGF62V7ND2KWBD

- Useful link: https://start.jcolemorrison.com/docker-fresh-mysql-or-mongodb-instances-in-projects/

