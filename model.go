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
}

func (s *service) createService(db *sql.DB) error {
	_, err :=
		db.Exec("INSERT INTO services (name,userid) VALUES($1,$2)", s.NAME, s.USERID)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) deleteService(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getServices(userid string, db *sql.DB) ([]string, error) {
	var serviceNames []string
	rows, err :=
		db.Query("SELECT name FROM services WHERE userid=$1", userid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		serviceNames = append(serviceNames, name)
	}
	return serviceNames, nil
}
