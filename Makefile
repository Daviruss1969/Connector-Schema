EXECUTABLE_NAME=PAValidator

GC = go build
GFLAGS = -o $(EXECUTABLE_NAME)

all: $(EXECUTABLE_NAME)

$(EXECUTABLE_NAME): main.go lib/*
	$(GC) $(GFLAGS) .

clean:
	rm $(EXECUTABLE_NAME)
	rm test/c.out
	rm test/cover.html

test: all
	go test -v -cover PAValidator/lib

coverage: all
	go test -v -coverprofile test/c.out PAValidator/lib
	go tool cover -html test/c.out -o test/cover.html