target = server

all: build

deps:
	@go get -u github.com/labstack/echo
	@go get -u github.com/valyala/quicktemplate/qtc
	@go get -u github.com/gorilla/securecookie

build:
	@echo -n "Building..."
	@mkdir -p bin
	@go generate 2>/dev/null
	@go build -o bin/$(target) *.go
	@echo "Done"

clean:
	@echo "Cleaning..."
	@rm bin/$(target)
	@rmdir bin
	@rm templates/*.go
