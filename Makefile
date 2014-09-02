GOCMD   = go
GOBUILD = $(GOCMD) build
GOFMT   = $(GOCMD) fmt
GOCLEAN = $(GOCLEAN)
GOPATH  = ${PWD}:${GOPATHi}
export $(GOPATH)

SRCDIR   = ${PWD}/src
BINDIR   = ${PWD}/bin
PACKAGE  = dc2
SRCFILE  = $(SRCDIR)/$(PACKAGE)/$(PACKAGE).go
BINFILE  = $(BINDIR)/$(PACKAGE)

DESTDIR    = .
INSTALLDIR = $(DESTDIR)/usr

DESTBIN = $(INSTALLDIR)/bin

.PHONY: default build install clean

default: build

build:
	$(GOFMT) $(SRCDIR)/...
	$(GOBUILD) -v -o $(BINFILE) $(SRCFILE)

clean:
	rm -rf $(BINDIR)

install:
	mkdir -p $(DESTBIN)
	cp $(BINFILE) $(DESTBIN)
