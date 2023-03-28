package dbops

import "log"

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare(`insert into user(username, pwd) values(?, ?)`)
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare(`select pwd from user where username=?`)
	if err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM user WHERE username=? AND pwd=?`)
	if err != nil {
		log.Printf("DeleteUser %s", err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
