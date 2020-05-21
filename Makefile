NAME=torch
BINDIR=builds
GOBUILD=CGO_ENABLED=0 go build --ldflags="-s -w" -v -x -a
GOFILES=cmd/torch/*

PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64 \
	linux-armv5 \
	linux-armv6 \
	linux-armv7 \
	linux-armv8 \
	freebsd-amd64 \

WINDOWS_ARCH_LIST = \
	windows-amd64

all: linux-amd64 darwin-amd64 windows-amd64 # Most used

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-386:
	GOARCH=386 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-armv5:
	GOARCH=arm GOOS=linux GOARM=5 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-armv6:
	GOARCH=arm GOOS=linux GOARM=6 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-armv7:
	GOARCH=arm GOOS=linux GOARM=7 $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

linux-armv8:
	GOARCH=arm64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

freebsd-amd64:
	GOARCH=amd64 GOOS=freebsd $(GOBUILD) -o $(BINDIR)/$(NAME)-$@ $(GOFILES)

windows-amd64:
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o $(BINDIR)/$(NAME)-$@.exe $(GOFILES)

# gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))
# zip_releases=$(addsuffix .zip, $(WINDOWS_ARCH_LIST))

# $(gz_releases): %.gz : %
# 	chmod +x $(BINDIR)/$(NAME)-$(basename $@)
# 	gzip -f -S -$(VERSION).gz $(BINDIR)/$(NAME)-$(basename $@)

# $(zip_releases): %.zip : %
# 	zip -m -j $(BINDIR)/$(NAME)-$(basename $@)-$(VERSION).zip $(BINDIR)/$(NAME)-$(basename $@).exe

all-arch: $(PLATFORM_LIST) $(WINDOWS_ARCH_LIST)

clean:
	rm $(BINDIR)/*