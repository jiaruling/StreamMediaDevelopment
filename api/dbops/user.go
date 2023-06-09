package dbops

import (
	"database/sql"
	"log"

	"github.com/jiaruling/StreamMediaDevelopment/api/defs"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare(`insert into user(username, pwd) values(?, ?)`)
	if err != nil {
		log.Printf("AddUser %s", err.Error())
		return err
	}
	defer stmtIns.Close()
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare(`select pwd from user where username=?`)
	if err != nil {
		log.Printf("GetUser %s", err.Error())
		return "", err
	}
	defer stmtOut.Close()
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM user WHERE username=? AND pwd=?`)
	if err != nil {
		log.Printf("DeleteUser %s", err)
		return err
	}
	defer stmtDel.Close()
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM user WHERE username = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var id int
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}

	defer stmtOut.Close()

	return res, nil
}
