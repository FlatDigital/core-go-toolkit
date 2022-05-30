package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_PoolStats(t *testing.T) {
	// given
	ass := assert.New(t)

	readTimeout := time.Second

	config := ServiceConfig{
		ConnReadTimeout: &readTimeout,
	}
	service, sqlMock := newMockService(config)

	statsOutput := sql.DBStats{}

	// when
	sqlMock.PatchStats(statsOutput)
	stats := service.PoolStats()

	// then
	ass.NotNil(stats)
}
