## API Spec
1. https://app.swaggerhub.com/apis/najibjodiansyah/employee-restful_api/1.0.0

## How to run using docker
1. Install docker and docker compose
2. Run `docker-compose up --build -d`
3. Hit `http://localhost:3000/v1/employees`
4. Enjoy!

## Set up Postgres in docker
1. Run `docker-compose exec api sh`
2. Run `migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose up`

## set up locally
1. Install postgres
2. Create database `mekari_employee`
3. Run go mod download
4. Run go install --tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
5. Run make migrate-up
6. Run `go run main.go`

## test env
Create database `mekari_employee_test`
1. Run env MKR_ENV=test
2. Run go test ./...
3. Run go test -coverprofile=coverage.out ./...
4. Run go tool cover -html=coverage.out -o cover.html
5. Run go tool cover -func coverage.out
6. Run go test -v -coverpkg=./... -coverprofile=profile.cov ./...
