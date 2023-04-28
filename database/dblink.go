package database

import (
	"fmt"
)

type DbLink struct {
	ConnectionName string
	Host           string
	Port           uint
	User           string
	Password       string
	DbName         string
}

func NewDbLinkConnection(connectionName string, host string, port uint, user string, password string, dbName string) *DbLink {
	return &DbLink{
		ConnectionName: connectionName,
		Host:           host,
		Port:           port,
		User:           user,
		Password:       password,
		DbName:         dbName,
	}
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
