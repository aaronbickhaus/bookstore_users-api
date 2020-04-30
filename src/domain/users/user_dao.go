package users

import (
	"github.com/aaronbickhaus/bookstore_users-api/src/datasource/mysql/users_db"
	"github.com/aaronbickhaus/bookstore_users-api/src/logger"
	"github.com/aaronbickhaus/bookstore_users-api/src/utils/errors"
	"github.com/aaronbickhaus/bookstore_users-api/src/utils/mysql_utils"
)

const (
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status from user WHERE status=?"
	queryDeleteUser = "DELETE FROM user WHERE id=?"
	queryUpdateUser = "UPDATE user SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryInsertUser = "INSERT into user(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created, status FROM user WHERE id=?;"
)
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		logger.Error("error when trying to prepare SaveUser", err)
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)

	if saveErr != nil {
		logger.Error("error when trying to execute SaveUser", err)
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to insert SaveUser", err)
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Get()  *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare GetUser", err)
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to scan GetUser", err)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		logger.Error("error when trying to prepare DeleteUser", err)
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	if err !=nil {
		logger.Error("error when trying to execute DeleteUser", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}
func (user *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("error when trying to prepare UpdateUser", err)
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id, user.Status, user.Password)

	if err !=nil {
		logger.Error("error when trying to execute UpdateUser", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare SearchUser", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to execute SearchUser", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()

	 results := make([]User, 0)
	 for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan rows SearchUser", err)
			return  nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	 }
	 if len(results) == 0 {
	 	return nil, errors.NewNotFoundError("no users matching status found")
	 }
	return results, nil
}