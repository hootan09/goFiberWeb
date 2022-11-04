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

TODOS:

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
- [#] JWT middleware & utils (generator)
- [ ] validators
- [ ] /ws Websocket
- [ ] workFlow (goReleaser)