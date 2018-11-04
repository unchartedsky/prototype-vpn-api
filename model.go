package main

import (
	"database/sql"
	"errors"
)

type user struct {
	ID       int    `json:"id"`
	USERID   string `json:"userid"`
	Password string `json:"password"`
}
type service struct {
	ID     int    `json:"id"`
	NAME   string `json:"name"`
	USERID string `json:"userid"`
}

func (u *user) login(db *sql.DB) error {
	rows, err :=
		db.Query("SELECT userid, password FROM users WHERE userid=$1 AND password=$2", u.USERID, u.Password)
	if err != nil {
		return err
	}
	if !rows.Next() {
		return errors.New("Invalid User Information")
	}
	return nil
}
func (u *user) signup(db *sql.DB) error {
	_, err :=
		db.Exec("INSERT INTO users (userid,password) VALUES($1,$2)", u.USERID, u.Password)
	if err != nil {
		return err
	}
	return nil
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
