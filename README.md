GoV2Web

#### Requirements:
```sh
go get -u github.com/gofiber/fiber/v2
go get -u github.com/mattn/go-sqlite3
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/gofiber/template
go get -u github.com/google/uuid
go get -u github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt/v4
go get -u github.com/gofiber/jwt/v3
```

#### Swagger
```sh
# go get -u github.com/swaggo/swag/cmd/swag   //old
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseDependency --parseInternal
go get -u github.com/gofiber/swagger
```

### Air Config (Live Reload Development)
```sh
#go install github.com/cosmtrek/air@latest
air init
air -c .air.toml
```

### GCC for windows
```sh
# for building sqlite driver in windows
https://jmeubank.github.io/tdm-gcc/
go env -w CGO_ENABLED=1
go install github.com/mattn/go-sqlite3
```

### Disable GoProxy
```sh
#Linx
export GOPROXY=direct
# windows(cmd)
set GOPROXY=direct
```

#### Running
```sh
go mod tidy
go run .
# Development
make dev
# update Docs </apidoc>
make swag
```


#### TODOS: (some sample code in tmp folder)

- [X] Swagger automatic documention **$make swag**
- [X] sqlite3 database
- [X] Api & web endpoint routes
- [X] Gorm
- [X] models
- [X] makefile
- [X] .env config
- [X] api & web group section in main.go
- [X] air livereload
- [X] Template engine (.html)
- [X] embed statics files
- [X] session for web auth with cookie expiration
- [X] JWT middleware & utils (generator)
- [X] File Uploader
- [X] CSV Downloader (get all users)
- [ ] validators [go-playground validator](https://github.com/go-playground/validator)
- [ ] /ws Websocket
- [ ] workFlow (goReleaser)
- [ ] not embed template in dev mode
- [X] bug in get all users (email not shown)
- [ ] speedup compile time (sqlite in windows)
- [ ] swagger fix annotations
- [ ] make two stage DockerFile (builder & production)
- [ ] database Constans with one connection for all route