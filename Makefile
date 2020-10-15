build:
	go build -o learn-go-with-test cmd/tdd/main.go

run:
	./learn-go-with-test

clean:
	rm learn-go-with-test