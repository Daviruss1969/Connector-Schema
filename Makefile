EXECUTABLE_NAME=csch.exe

GC = go build
GFLAGS = -o $(EXECUTABLE_NAME)
LIBMOD = ConnectorSchema/lib
TESTDIR = lib/test

all: $(EXECUTABLE_NAME)

$(EXECUTABLE_NAME): main.go lib/*
	$(GC) $(GFLAGS) .

clean:
	rm $(EXECUTABLE_NAME)
	rm $(TESTDIR)/c.out
	rm $(TESTDIR)/cover.html

test: all
	go test -v -cover $(LIBMOD)

coverage: all
	go test -v -coverprofile $(TESTDIR)/c.out $(LIBMOD)
	go tool cover -html $(TESTDIR)/c.out -o $(TESTDIR)/cover.html