.PHONY: clean, test, run

run: main
	./main --currencies exalted scan

clean:
	rm -f main

main:
	go build -o main

test: clean
	go test -v