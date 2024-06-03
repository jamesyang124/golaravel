BINARY_NAME=go-laravel

# @ not show command
# - conitinue next command if failed

build:
	@go mod vendor
	@echo "Building Celeritas..."
	@go build -o ./${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	@./${BINARY_NAME} &
	@echo "Celeritas started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ./${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Celeritas..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped Celeritas!"

restart: stop start