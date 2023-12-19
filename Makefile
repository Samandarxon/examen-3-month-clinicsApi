
run:
	go run cmd/main.go

gen-swag:
	swag init -g ./api/api.go -o ./api/docs

migration-up:
	migrate -path ./db/sql/ -database "postgresql://samandarxon:1234@localhost:5432/market_system?sslmode=disable" -verbose up

migration-down:
	migrate -path ./db/sql/ -database "postgresql://samandarxon:1234@localhost:5432/market_system?sslmode=disable" -verbose down

all: gen-swag 
	sleep 5 && air
gitA:
	git init && git add . && git commit -m "clinics" && git branch -M main && git remote add origin git@github.com:Samandarxon/examen-3-month-clinicsApi.git

gitPush:
git push -u origin main

air:
	air
zipA:
	zip -ru clinics.zip ./