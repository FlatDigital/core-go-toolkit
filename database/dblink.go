package database

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type DbLink struct {
	ConnectionName string `validate:"required"`
	Host           string `validate:"required"`
	Port           uint   `validate:"required"`
	User           string `validate:"required"`
	Password       string `validate:"required"`
	DbName         string `validate:"required"`
}

func NewDbLinkConnection(connectionName string, host string, port uint, user string, password string, dbName string) (*DbLink, error) {
	dbLinkConn := &DbLink{
		ConnectionName: connectionName,
		Host:           host,
		Port:           port,
		User:           user,
		Password:       password,
		DbName:         dbName,
	}
	err := dbLinkConn.validate(dbLinkConn)
	if err != nil {
		return nil, err
	}
	return dbLinkConn, nil
}

func (service DbLink) OpenConnection() string {
	return fmt.Sprintf("SELECT * FROM dblink_connect('%s', 'host=%s port=%d dbname=%s user=%s password=%s')",
		service.ConnectionName,
		service.Host,
		service.Port,
		service.DbName,
		service.User,
		service.Password,
	)
}

func (service DbLink) CloseConnection() string {
	return fmt.Sprintf("SELECT dblink_disconnect('%s')", service.ConnectionName)
}

func (service DbLink) validate(dbLinkConn *DbLink) error {
	validate := validator.New()
	err := validate.Struct(dbLinkConn)
	if err != nil {
		return err
	}
	return nil
}
