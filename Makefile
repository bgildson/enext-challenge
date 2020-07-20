test:
	go test -v ./...

parse:
	go run ./cmd/parser/main.go -log ./games.log -out ./games.json

report-general:
	go run cmd/report/main.go -games-json-path=./games.json -general=true

report-by-game:
	go run cmd/report/main.go -games-json-path=./games.json -general=false

api:
	go run ./cmd/api/main.go -games-json-path=./games.json -port=8080

.PHONY: test parse report-general report-by-game api
