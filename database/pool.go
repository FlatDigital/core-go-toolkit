package database

import "database/sql"

func (service *service) PoolStats() sql.DBStats {
	return service.db.Stats()
}
