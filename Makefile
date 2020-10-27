build:
	go build -o player-service cmd/httpserver/main.go

run:
	./player-service

clean:
	rm player-service

build-aws:
	go build -o hello-aws cmd/aws/main.go

run-aws:
	%USERPROFILE%\Go\bin\build-lambda-zip.exe -o hello-aws.zip hello-aws
