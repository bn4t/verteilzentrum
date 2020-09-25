SHELL = /bin/bash

# use bash strict mode
.SHELLFLAGS := -eu -o pipefail -c

.ONESHELL:
.DELETE_ON_ERROR:

.SUFFIXES:      # delete the default suffixes
.SUFFIXES: .go  # add .go as suffix

PREFIX?=/usr/local
_INSTDIR=$(DESTDIR)$(PREFIX)
BINDIR?=$(_INSTDIR)/bin
GO?=go
GOFLAGS?=
RM?=rm -f # Exists in GNUMake but not in NetBSD make and others.


build:
	$(GO) build $(GOFLAGS) -o verteilzentrum .

run: build
	./verteilzentrum

install: build
	id "verteilzentrum" &>/dev/null || useradd -MUr verteilzentrum
	install -m755 -gverteilzentrum -overteilzentrum verteilzentrum $(BINDIR)/verteilzentrum
	setcap 'cap_net_bind_service=+ep' $(BINDIR)/verteilzentrum # allow verteilzentrum to bind to priviledged ports as non-root user
	mkdir -p /etc/verteilzentrum
	mkdir -p /var/lib/verteilzentrum
	chown -R verteilzentrum:verteilzentrum /etc/verteilzentrum
	chown -R verteilzentrum:verteilzentrum /var/lib/verteilzentrum
	if [ ! -f "/etc/verteilzentrum/config.toml" ]; then
		install -m600 -gverteilzentrum -overteilzentrum configs/config.example.toml /etc/verteilzentrum/config.toml
	fi

install-systemd:
	install -m644 -groot -oroot init/verteilzentrum.service /etc/systemd/system/verteilzentrum.service
	systemctl daemon-reload

clean:
	$(RM) verteilzentrum


RMDIR_IF_EMPTY:=sh -c '\
if test -d $$0 && ! ls -1qA $$0 | grep -q . ; then \
	rmdir $$0; \
fi'


uninstall:
	$(RM) $(BINDIR)/verteilzentrum
	$(RMDIR_IF_EMPTY) /etc/verteilzentrum
	$(RMDIR_IF_EMPTY) /var/lib/verteilzentrum

uninstall-systemd:
	systemctl stop verteilzentrum
	$(RM) /etc/systemd/system/verteilzentrum.service
	systemctl daemon-reload

.DEFAULT_GOAL = build
.PHONY: all build install uninstall clean install-systemd uninstall-systemd

