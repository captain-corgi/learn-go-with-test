build:
	go build -o player-service cmd/httpserver/main.go

run:
	./player-service

clean:
	rm player-service