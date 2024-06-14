CURRENT_DIR=$(shell pwd)
DBURL=postgres://n10:n10@localhost:5432/n10?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

mig-up:
	migrate -path migrations -database '$(DBURL)' -verbose up

mig-down:
	migrate -path migrations -database '$(DBURL)' -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table
