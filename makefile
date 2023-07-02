migration_up: migrate -path src/database/migrations/ -database "postgresql://postgres:todo@1234@localhost:5438/todo-go?sslmode=disable" -verbose up

migration_down: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down

migration_fix: migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force VERSION