# go-file-manager
File Manager that let you upload and download file coded with Go Programming Language (Golang)

### PreInstall
1. For using Docker https://docs.docker.com/desktop/install/windows-install/

---
### Database (Not needed if you use Docker)
1. Create a database named `file_manager`
2. Run `go run main.go` to create the tables using gorm
3. It will automatically create the tables using gorm
---

### How to use with Docker
1. Clone this Repository using `git clone`
2. Run Docker Desktop
3. Check config/db.go for the commented out code and adjust accordingly
4. use CMD Command on this project root
    1. `docker-compose build`
    2. `docker-compose up`
5. access http://localhost:8080/swagger/index.html#/
6. The app is ready to use.
---
### How to use with Go Build
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go build -tags netgo -ldflags '-s -w' -o app`
    2. `./app`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---
### How to use with Go Run
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go run main.go`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.
---



For more information, please contact me LinkedIn: https://www.linkedin.com/in/MicBun