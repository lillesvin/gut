BINARY_BASENAME = gut
DEST_DIR = $(HOME)/bin

all: linux windows macosx

linux: main.go
	GOOS=linux ARCH=amd64 go build -ldflags=$(LDFLAGS) -o $(BINARY_BASENAME)_linux-amd64 .

windows: main.go
	GOOS=windows ARCH=amd64 go build -ldflags=$(LDFLAGS) -o $(BINARY_BASENAME)_windows-amd64 .

macosx: main.go
	GOOS=darwin ARCH=amd64 go build -ldflags=$(LDFLAGS) -o $(BINARY_BASENAME)_macosx-amd64 .

install:
	install -m0755 -T $(BINARY_BASENAME)_linux-amd64 $(DEST_DIR)/$(BINARY_BASENAME)

clean:
	go clean
	rm -f $(BINARY_BASENAME)_*-amd64

.PHONY: all install clean
