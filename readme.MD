## How to run using docker
1. Install docker
2. Run `docker-compose up` in the root directory
3. Hit `http://localhost:3000`
4. Enjoy!

## Set up Postgres in docker
1. Run `docker-compose exec api sh`
2. Run `migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose up`

