migrate:
	@go run scripts/migrate/main.go migrate

seed:
	@go run scripts/seed/main.go seeder

dev:
	@${GOPATH}/bin/air
