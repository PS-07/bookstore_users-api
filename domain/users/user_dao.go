package users

import (
	"fmt"
	"strings"

	"github.com/PS-07/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/PS-07/bookstore_users-api/logger"
	"github.com/PS-07/bookstore_users-api/utils/errors"
	"github.com/PS-07/bookstore_users-api/utils/mysqlutils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

// Get func
func (user *User) Get() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	result := statement.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil
}

// Update func
func (user *User) Update() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	_, updateErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		logger.Error("error when trying to update user", updateErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// Delete func
func (user *User) Delete() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	_, deleteErr := statement.Exec(user.ID)
	if deleteErr != nil {
		logger.Error("error when trying to delete user", deleteErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// FindByStatus func
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	statement, err := usersdb.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

// FindByEmailAndPassword func
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	statement, err := usersdb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysqlutils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error trying to get user by email and password", getErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
