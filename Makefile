# 
# Scripts
# 

# CC := gcc-5 -std=gnu11 -Wall

run: clean
	go run main.go

clean:
	@- rm -f main
	@- rm -f *.o
