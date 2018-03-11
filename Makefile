# 
# Scripts
# 

main: clean

run: clean setup
	@- go run zoey.go

test: setup
	@- go test ./...

setup:
	@- go fmt ./...

clean:
	@- rm -f main
	@- rm -f *.o
	@- rm -f *.out
