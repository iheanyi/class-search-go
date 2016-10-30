all: build

build: 
	go build
	echo "Done building!"

run:  build
	./class-search-go
