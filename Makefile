SHELL = /bin/bash

# use bash strict mode
.SHELLFLAGS := -eu -o pipefail -c

.ONESHELL:
.DELETE_ON_ERROR:

# check if recipeprefix is supported
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

.SUFFIXES:      # delete the default suffixe
.SUFFIXES: .go  # add .go as suffix

PREFIX?=/usr/local
_INSTDIR=$(DESTDIR)$(PREFIX)
BINDIR?=$(_INSTDIR)/bin
GO?=go
GOFLAGS?=
RM?=rm -f # Exists in GNUMake but not in NetBSD make and others.


all: build

build:
> $(GO) build $(GOFLAGS) -o verteilzentrum .

run: build
> ./verteilzentrum

install: build
> install -m755 verteilzentrum $(BINDIR)/verteilzentrum
> if [ ! -f "/etc/verteilzentrum/config.toml" ]; then
>   install -m600 configs/config.example.toml /etc/verteilzentrum/config.toml
> fi

clean:
> $(RM) verteilzentrum


RMDIR_IF_EMPTY:=sh -c '\
if test -d $$0 && ! ls -1qA $$0 | grep -q . ; then \
	rmdir $$0; \
fi'


uninstall:
> $(RM) $(BINDIR)/verteilzentrum
> $(RMDIR_IF_EMPTY) /etc/verteilzentrum

.DEFAULT_GOAL = all
.PHONY: all build install uninstall clean

