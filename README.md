GoTicket

#### Requirements:
```sh
go get -u github.com/gofiber/fiber/v2
go get -u github.com/mattn/go-sqlite3
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/gofiber/template
go get -u github.com/google/uuid
go get -u github.com/joho/godotenv
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

#### Running
```sh
go mod tidy
go run .
# Development
make dev
```