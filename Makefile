test:
	go test -v ./...

parse:
	go run ./cmd/parser/main.go -log ./games.log -out ./games.json
