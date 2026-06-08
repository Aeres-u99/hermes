BINARY_NAME=hermes
CMD_PATH=./cmd/hermes

build: clean
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
