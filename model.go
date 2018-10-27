package main

import (
	"database/sql"
	"errors"
)

type user struct {
	ID       int    `json:"id"`
	USERID   string `json:"id"`
	Password string `json:"password"`
}
type service struct {
	ID     int    `json:"id"`
	NAME   string `json:"name"`
	USERID string `json:"user_id"`
}

func (u *user) login(db *sql.DB, id string, password string) error {
	return errors.New("Not implemented")
}
func (u *user) signup(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (s *service) getService(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) createService(db *sql.DB) error {
	return errors.New("Not implemented")
}
func (u *user) deleteService(db *sql.DB) error {
	return errors.New("Not implemented")
}
