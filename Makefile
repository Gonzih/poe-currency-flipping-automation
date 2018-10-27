.PHONY: clean, test, run

run: main
	./main --currencies exalted --league Delve ui

clean:
	rm -f main

main:
	go build -o main

test: clean
	go test -v
