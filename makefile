migrate-up:
	@echo "Migrating up database..."
	@migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose up $(version)

migrate-down:
	@echo "Migrating down database..."
	@migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose down $(version)

migrate-force:
	@echo "Migrating force database..."
	@migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose force $(version)

migrate-status:
	@echo "Migrating status database..."
	@migrate -path migrations -database "postgres://postgres:@localhost:5432/mekari_employee?sslmode=disable" -verbose version

migrate-revision:
	@echo "Migrating revision database..."
	@migrate create -ext sql -dir migrations -seq $(name)