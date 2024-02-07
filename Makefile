build:
	go build -o ./bin/projectx

run: build
	./bin/projectx

# test:
# 	go test -v ./... 

clean: 
	rm -f ./bin/projectx

docker:
	docker run -it -v /Users/adi/repo/blocker/:/root/blocker --name blocker-go --network host golang:1.20

vm:
	docker exec -it blocker-go bash

test:
	go run main.go
