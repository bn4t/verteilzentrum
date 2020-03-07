package main

import "github.com/emersion/go-smtp"

type Backend struct{}

// we don't support authenticated logins
func (bkd *Backend) Login(_ *smtp.ConnectionState, _, _ string) (smtp.Session, error) {
	return nil, smtp.ErrAuthUnsupported
}

func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return &Session{}, nil
}
