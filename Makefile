.PHONY: clean, test, run

run: main

clean:
	rm -f main

main:
	go build -o main

test: clean
	go test -v
