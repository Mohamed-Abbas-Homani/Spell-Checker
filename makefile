build:
	go build -o ./bin/spellCheck

run: build
	./bin/spellCheck