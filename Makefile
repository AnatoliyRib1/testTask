before-push:
	go mod tidy &&\
	gofumpt -w -l . &&\
	go build ./...&&\
	golangci-lint run ./... &&\
	go test -v ./tests/...

run-service:
	go build -o ./client/bin/cli ./client/. &&\
	docker compose up -d --build

stop-service:
	rm -r ./client/bin &&\
	docker compose down