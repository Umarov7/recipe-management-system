DB_URL=postgres://postgres:root@localhost:5432/recipe_management?sslmode=disable

tidy:
	go mod tidy
swag-gen:
	~/go/bin/swag init -g ./internal/api/router.go -o internal/api/docs
run :
	go run cmd/main.go

mig-up:
	migrate -path migrations -database ${DB_URL} -verbose up

mig-down:
	migrate -path migrations -database ${DB_URL} -verbose down

mig-force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

mig-cr:
	migrate create -ext sql -dir migrations -seq ${1}_${2}_${3}