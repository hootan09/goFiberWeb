# sudo apt install golang-golang-x-tools
doc:
	godoc -http=:8080 
test:
	go test -v
build: clean test
	go build .
run:
	go run .
dev:
	air -c .air.toml
install:
	go install .
clean:
	rm -rf ./build
swag:
	swag init --parseDependency --parseInternal

.PHONY: doc test build run dev install clean swag