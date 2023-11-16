EXECUTABLE_NAME=PAValidator

GC = go build
GFLAGS = -o $(EXECUTABLE_NAME)

all: $(EXECUTABLE_NAME)

$(EXECUTABLE_NAME): main.go lib/*
	$(GC) $(GFLAGS) .

clean:
	rm $(EXECUTABLE_NAME)

test: all
	go test -v