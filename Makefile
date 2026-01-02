LDFLAGS = "-L/usr/lib -lusb-1.0"
GCFLAGS = -gcflags=all="-N -l"

.PHONY: clean
clean:
	rm -rf bin/mtpfs

.PHONY: build
build: clean
	CGO_ENABLED=1 go build -o bin/mtpfs .