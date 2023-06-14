package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDbLinkConnection(t *testing.T) {
	connName := "test"
	host := "127.0.0.1"
	user := "test"
	pass := "pass123"
	db := "dbtest"
	port := uint(1123)
	newDbLinkConn, err := NewDbLinkConnection(connName, host, port, user, pass, db)
	assert.NotNil(t, newDbLinkConn)
	assert.NoError(t, err)
}

func TestNewDbLinkConnection_ConnNotValid(t *testing.T) {
	connName := "test"
	host := ""
	user := ""
	pass := "pass123"
	db := "dbtest"
	port := uint(1123)
	newDbLinkConn, err := NewDbLinkConnection(connName, host, port, user, pass, db)
	assert.Nil(t, newDbLinkConn)
	assert.Error(t, err)
}

func TestNewDbLinkConnection_OpenConnection(t *testing.T) {
	connName := "test"
	host := "127.0.0.1"
	user := "test"
	pass := "pass123"
	db := "dbtest"
	port := uint(1123)
	newDbLinkConn, _ := NewDbLinkConnection(connName, host, port, user, pass, db)
	conn := newDbLinkConn.OpenConnection()
	connExpected := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1123 dbname=dbtest user=test password=pass123')"
	assert.Equal(t, connExpected, conn)
}

func TestNewDbLinkConnection_CloseConnection(t *testing.T) {
	connName := "test"
	host := "127.0.0.1"
	user := "test"
	pass := "pass123"
	db := "dbtest"
	port := uint(1123)
	newDbLinkConn, _ := NewDbLinkConnection(connName, host, port, user, pass, db)
	closeConn := newDbLinkConn.CloseConnection()
	connExpected := "SELECT dblink_disconnect('test')"
	assert.Equal(t, connExpected, closeConn)
}
