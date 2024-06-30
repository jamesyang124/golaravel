BINARY_NAME=golaravel

# @ not show command
# - conitinue next command if failed
# build: cleanApp clean -> run cleanApp then clean then build steps

build: #cleanApp
	@go mod tidy
	@go mod vendor
	@echo "Building Celeritas..."
	@go build -o ./${BINARY_NAME} .
	@echo "Celeritas built!"

run: build
	@echo "Starting Celeritas..."
	@./${BINARY_NAME}
#	@./${BINARY_NAME} &
#	@echo "Celeritas started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ./${BINARY_NAME}
	@echo "Cleaned!"

cleanApp:
	@echo "Cleaning application scaffold files..."
#	@-rm -r "./handler"
	@-rm -r "./migrations"
#	@-rm -r "./views"
	@-rm -r "./data"
#	@-rm -r "./public"
	@-rm -r "./tmp"
	@-rm -r "./logs"
	@-rm -r "./middleware"
#	@-rm -r "./.env"
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

startCompose:
	docker-compose up -d

stopCompose:
	docker-compose stop
	docker-compose rm -f
#	docker rm $(docker ps -a | awk '{print $1}' | tail -n +2)
	docker network rm golaravel_default