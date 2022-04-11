package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/database/converter"
	"github.com/lib/pq"
)

var (
	errDBCNil                  = errors.New("dbc_nil")
	errTransactionNil          = errors.New("trx_nil")
	errNoLestingLevel          = errors.New("no_lesting_level")
	errPanicOnBeginTransaction = errors.New("panic_on_begin_trx")
)

type (
	// Database service interface
	Database interface {
		PoolStats() sql.DBStats
		TestConnection(dbc *DBContext) error
		ExecuteWithQuery(dbc *DBContext, query *Query) (*DBResult, error)
		Execute(dbc *DBContext, query string, params ...interface{}) (*DBResult, error)
		ExecuteEnsuringOneAffectedRowWithQuery(dbc *DBContext, query *Query) error
		ExecuteEnsuringOneAffectedRow(dbc *DBContext, query string, params ...interface{}) error
		SelectWithQuery(dbc *DBContext, query *Query) (*DBResult, error)
		Select(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBResult, error)
		SelectUniqueValueWithQuery(dbc *DBContext, query *Query) (*DBRow, error)
		SelectUniqueValue(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBRow, error)
		SelectUniqueValueNonEmptyWithQuery(dbc *DBContext, query *Query) (*DBRow, error)
		SelectUniqueValueNonEmpty(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBRow, error)
		Begin(dbc *DBContext) (*DBContext, error)
		Commit(dbc *DBContext) error
		Rollback(dbc *DBContext) error
		Connection() (*DBContext, error)
		Close(dbc *DBContext) error
	}

	// ServiceConfig database service config
	ServiceConfig struct {
		MaxConnectionRetries int
		DatadogMetricPrefix  string

		DBHost           string
		DBName           string
		DBPassword       string
		DBUsername       string
		MaxIdleConns     int
		MaxOpenConns     int
		ConnMaxLifetime  time.Duration
		ConnReadTimeout  *time.Duration
		ConnWriteTimeout *time.Duration
		ConnTimeout      *time.Duration
	}

	// DBContext database transaction token
	DBContext struct {
		tx           converter.DBTxer
		nestingLevel int
		dbConn       converter.DBConner
		ctx          context.Context
	}

	// Service type
	service struct {
		db                   converter.DBer
		maxConnectionRetries int
		datadogMetricPrefix  string
	}

	logType string
)

const (
	// default values
	defaultMaxConnectionRetries = 3

	logError   logType = "error"
	logSuccess logType = "success"
)

// NewService returns a database service interface
func NewService(config ServiceConfig) (Database, error) {
	// connection for MySQL
	// connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.DBUsername,
	// 	config.DBPassword, config.DBHost, config.DBName)

	// connection for Postgress
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, 5432, config.DBUsername, config.DBPassword, config.DBName)
	// if config has ConnReadTimeout set, appends readTimeout param
	if config.ConnReadTimeout != nil {
		connectionString = fmt.Sprintf("%s&readTimeout=%s", connectionString, config.ConnReadTimeout.String())
	}
	// if config has ConnWriteTimeout set, appends writeTimeout param
	if config.ConnWriteTimeout != nil {
		connectionString = fmt.Sprintf("%s&writeTimeout=%s", connectionString, config.ConnWriteTimeout.String())
	}
	// if config has ConnTimeout set, appends timeout param
	if config.ConnTimeout != nil {
		connectionString = fmt.Sprintf("%s&timeout=%s", connectionString, config.ConnTimeout.String())
	}
	db, err := sql.Open("postgres", connectionString)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	connMaxLifetime := config.ConnMaxLifetime * time.Second
	db.SetConnMaxLifetime(connMaxLifetime)
	if err != nil {
		return nil, err
	}

	retries := defaultMaxConnectionRetries
	if config.MaxConnectionRetries > 0 {
		retries = config.MaxConnectionRetries
	}

	var metricPrefix string
	if config.DatadogMetricPrefix != "" {
		metricPrefix = config.DatadogMetricPrefix
	} else {
		metricPrefix = os.Getenv("APPLICATION")
	}

	sqllib := converter.SQLToDBer(db)

	// make the service
	service := &service{
		db:                   sqllib,
		maxConnectionRetries: retries,
		datadogMetricPrefix:  metricPrefix,
	}

	// done
	return service, nil
}

// Connection returns a new connection that can be used to execute queries always in the same connection
func (service *service) Connection() (*DBContext, error) {
	var dbctx *DBContext
	var err error

	// Obtain the ctx
	ctx := context.Background()
	for retry := 0; retry < service.maxConnectionRetries; retry++ {
		// Obtain the connection
		conn, err := service.db.Conn(ctx)
		if err != nil {
			service.logMetric(logError, "connection", "service.db.Conn(ctx)", err)
			continue
		}

		// Build the object and return
		dbctx = &DBContext{
			tx:           nil,  // we don't have a tx
			nestingLevel: 0,    // init with nestingLevel as 0
			dbConn:       conn, // set the connection
			ctx:          ctx,  // use the ctx obtained previously
		}

		err = service.TestConnection(dbctx)
		if err != nil {
			service.logMetric(logError, "connection", "service.TestConnection(dbctx)", err)
			dbctx = nil
			continue
		} else {
			service.logMetric(logSuccess, "connection", "service.TestConnection(dbctx)", nil)
			break
		}
	}

	// in case of leaving the loop by continue
	if dbctx == nil {
		return nil, errors.New("couldn't connect to database")
	}

	return dbctx, err
}

// TestConnection tests the given connection
func (service *service) TestConnection(dbc *DBContext) error {
	if dbc != nil && (dbc.tx != nil || dbc.dbConn != nil) {
		return service.db.PingContext(dbc.ctx)
	}
	return service.db.Ping()
}

// Close closes a given connection
func (service *service) Close(dbc *DBContext) error {
	// We have a valid dbc?
	if dbc == nil {
		return fmt.Errorf("you must send a DBContext in order to close")
	}

	// We can't close if we have an active transaction
	if dbc.tx != nil {
		return fmt.Errorf("you are closing a connection with an active transaction")
	}

	// Close the connection
	err := dbc.dbConn.Close()
	if err != nil {
		service.logMetric(logError, "close", "dbc.dbConn.Close()", err)
		return err
	}

	// Set dbConn to nil
	dbc.dbConn = nil

	// done
	return nil
}

// Begin starts a transaction in the database
func (service *service) Begin(inDbc *DBContext) (outDbc *DBContext, err error) {
	// recover from any panic.
	// Known issue in the following "if" block, dbc should be nill after the "or" operator
	// It might be a race condition with multiple routines setting dbc to nil
	defer func() {
		if r := recover(); r != nil {
			service.logMetric(logError, "begin", "recover", errPanicOnBeginTransaction)

			// TODO: Implement datadog
			// tags := new(godog.Tags).
			// 	ToArray()
			// godog.RecordSimpleMetric(fmt.Sprintf("application.%s.db.service.connection.begin.panic_recover",
			// 	service.datadogMetricPrefix), 1, tags...)
		}
	}()

	outDbc = inDbc
	// If we don't have a dbc, we create one calling Connection()
	// We also support the case in which both tx and dbConn are nil,
	// in this case we threat everything like a nil dbc.
	if outDbc == nil || (outDbc.tx == nil && outDbc.dbConn == nil) {
		// Create a new connection
		newDbc, err := service.Connection()
		if err != nil {
			service.logMetric(logError, "begin", "service.Connection()", err)
			return nil, err
		}

		// Copy data from the newDbc to the original dbc
		if outDbc != nil {
			outDbc.tx = newDbc.tx
			outDbc.nestingLevel = newDbc.nestingLevel
			outDbc.dbConn = newDbc.dbConn
			outDbc.ctx = newDbc.ctx
		} else {
			outDbc = newDbc
		}
	}

	// If the nesting level is 0, we start a real transaction
	if outDbc.nestingLevel == 0 {
		// We begin a real transaction
		tx, err := service.db.BeginTx(outDbc.ctx, nil)
		if err != nil {
			service.logMetric(logError, "begin", "service.db.BeginTx(outDbc.ctx, nil)", err)
			return nil, err
		}

		// Set into the dbc
		outDbc.tx = tx
	}

	// Increment the nesting level counter
	outDbc.nestingLevel++

	// done
	return outDbc, nil
}

// Commit commits the active transaction
func (service *service) Commit(dbc *DBContext) error {
	// We have a valid dbc?
	if dbc == nil {
		service.logMetric(logError, "commit", "validation", errDBCNil)
		return fmt.Errorf("you must send a DBContext in order to commit")
	}

	// Check the nesting level
	if dbc.nestingLevel <= 0 {
		service.logMetric(logError, "commit", "validation", errNoLestingLevel)
		return fmt.Errorf("the nestingLevel must be greater than 0 in order to call Commit()")
	}

	// Check tx
	if dbc.tx == nil {
		service.logMetric(logError, "commit", "validation", errTransactionNil)
		return fmt.Errorf("the dbc.tx was not set before call Commit()")
	}

	// We only trigger a Commit() if we are reaching 0
	if dbc.nestingLevel == 1 {
		err := dbc.tx.Commit()
		if err != nil {
			service.logMetric(logError, "commit", "dbc.tx.Commit()", err)
			return err
		}

		// Reset the tx because it's no longer valid
		dbc.tx = nil

		// 2018-07-02: If you don't call Close() after a Commit(),
		// the connection is never released. So, we force the close.
		err = service.Close(dbc)
		if err != nil {
			service.logMetric(logError, "commit", "service.Close(dbc)", err)
			return err
		}
	}

	// We decrease the nesting level
	dbc.nestingLevel--

	// done, we've done a "virtual" commit
	return nil
}

// Rollback rollbacks a transaction
func (service *service) Rollback(dbc *DBContext) error {
	// We have a valid dbc?
	if dbc == nil {
		service.logMetric(logError, "rollback", "validation", errDBCNil)
		return fmt.Errorf("you must send a DBContext in order to rollback")
	}

	// Check the nesting level
	if dbc.nestingLevel == 0 {
		// If we have a nestingLevel == 0, we return nil directly. Because maybe
		// another (more-deep) method has rollbacked the transaction.
		return nil
	}

	// Check tx
	if dbc.tx == nil {
		service.logMetric(logError, "rollback", "validation", errTransactionNil)
		return fmt.Errorf("the dbc.tx was not set before call Rollback()")
	}

	// We do the rollback
	err := dbc.tx.Rollback()
	if err != nil {
		service.logMetric(logError, "rollback", "dbc.tx.Rollback()", err)
		return err
	}

	// We always go to nestingLevel to 0
	dbc.nestingLevel = 0

	// Reset the tx because it's no longer valid
	dbc.tx = nil

	// 2018-07-02: If you don't call Close() after a Commit(),
	// the connection is never released. So, we force the close.
	err = service.Close(dbc)
	if err != nil {
		service.logMetric(logError, "rollback", "service.Close(dbc)", err)
		return err
	}

	// done
	return nil
}

// SelectWithQuery does a select in the database with a Query and process results returning a Map
func (service *service) SelectWithQuery(dbc *DBContext, query *Query) (*DBResult, error) {
	stmt, params, err := query.GetStatementAndParams()
	if err != nil {
		return nil, err
	}
	return service.Select(dbc, stmt, false, params...)
}

// Select does a select in the database and process results returning a Map
func (service *service) Select(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBResult, error) {
	// Add a "FOR UPDATE" at the end of the query if we have a true forUpdate flag.
	if forUpdate {
		query = regexp.MustCompile(`(?i)(FOR UPDATE|)(;|)$`).ReplaceAllString(strings.Trim(query, " "), " FOR UPDATE")
	}

	// Do the query and interpret results
	rows, err := service.doQuery(service.db, dbc, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Get columns
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Build the DbRows
	dbRowArray := make(DBRowArray, 0)
	for rows.Next() {
		// Builds columnPointers as an array with pointers
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create a DBColumns
		dbColumns := make(DBColumns)
		for i, colName := range cols {
			val := *columnPointers[i].(*interface{})
			dbColumns[colName] = DBColumn{
				name:  colName,
				field: val,
			}
		}

		// Append column into DBRowArray
		dbRowArray = append(dbRowArray, DBRow{columns: dbColumns})
	}

	// done
	return &DBResult{
		affectedRows: 0,
		rows:         DBRows{dbRowArray},
	}, nil
}

// SelectUniqueValueWithQuery does a select in the database with a Query and returns the first row
func (service *service) SelectUniqueValueWithQuery(dbc *DBContext, query *Query) (*DBRow, error) {
	stmt, params, err := query.GetStatementAndParams()
	if err != nil {
		return nil, err
	}
	return service.SelectUniqueValue(dbc, stmt, false, params...)
}

// SelectUniqueValue selects and returns the first row
func (service *service) SelectUniqueValue(dbc *DBContext, query string,
	forUpdate bool, params ...interface{}) (*DBRow, error) {
	dbResult, err := service.Select(dbc, query, forUpdate, params...)
	if err != nil {
		return nil, err
	}
	sizeRecords := len(dbResult.rows.DBRowArray)
	if sizeRecords == 0 {
		return nil, nil
	}

	if sizeRecords > 1 {
		return nil, fmt.Errorf("unexpected records size, expected 1 but was: %d", sizeRecords)
	}
	return &dbResult.rows.DBRowArray[0], nil
}

// SelectUniqueValueNonEmptyWithQuery does a select in the database with a
// Query and returns the first row and error if it don't exists
func (service *service) SelectUniqueValueNonEmptyWithQuery(dbc *DBContext, query *Query) (*DBRow, error) {
	stmt, params, err := query.GetStatementAndParams()
	if err != nil {
		return nil, err
	}
	return service.SelectUniqueValueNonEmpty(dbc, stmt, false, params...)
}

// SelectUniqueValueNonEmpty selects and returns the first row and error if it don't exists
func (service *service) SelectUniqueValueNonEmpty(dbc *DBContext, query string,
	forUpdate bool, params ...interface{}) (*DBRow, error) {
	dbRow, err := service.SelectUniqueValue(dbc, query, forUpdate, params...)
	if err != nil {
		return nil, err
	}
	if dbRow == nil {
		return nil, fmt.Errorf("unable to find record")
	}
	return dbRow, nil
}

// query executes a query inside a given transaction (if you have one)
func (service *service) doQuery(db converter.DBer, dbc *DBContext, query string,
	params ...interface{}) (converter.DBRowser, error) {
	var rows converter.DBRowser

	// We have a db transaction?
	// We only support the case in which dbc.tx or dbc.dbConn have anything... if not
	// we threat everything like if we have a nil dbc.
	if dbc != nil && (dbc.tx != nil || dbc.dbConn != nil) {
		// We have a transaction?
		if dbc.tx != nil {
			// Prepare the query
			stmt, err := dbc.tx.PrepareContext(dbc.ctx, query)
			if err != nil {
				service.logMetric(logError, "do_query", "dbc.tx.PrepareContext(dbc.ctx, query)", err)
				return nil, err
			}

			// Execute inside the transaction
			rows, err = stmt.QueryContext(dbc.ctx, params...)
			if err != nil {
				service.logMetric(logError, "do_query", "txstmt.QueryContext(dbc.ctx, params...)", err)
				return nil, err
			}
		} else if dbc.dbConn != nil {
			// Prepare the query
			stmt, err := db.PrepareContext(dbc.ctx, query)
			if err != nil {
				service.logMetric(logError, "do_query", "db.PrepareContext(dbc.ctx, query)", err)
				return nil, err
			}
			defer stmt.Close()

			// Execute using the context
			rows, err = stmt.QueryContext(dbc.ctx, params...)
			if err != nil {
				service.logMetric(logError, "do_query", "stmt.QueryContext(dbc.ctx, params...)", err)
				return nil, err
			}
		} else {
			// Not possible
			return nil, fmt.Errorf("you have sent a dbc without tx or dbConn")
		}
	} else {
		// We don't have a connection
		stmt, err := db.Prepare(query)
		if err != nil {
			service.logMetric(logError, "do_query", "db.Prepare(query)", err)
			return nil, err
		}
		defer stmt.Close()

		// Execute without context
		rows, err = stmt.Query(params...)
		if err != nil {
			service.logMetric(logError, "do_query", "stmt.Query(params...)", err)
			return nil, err
		}
	}

	// done
	return rows, nil
}

// ExecuteEnsuringOneAffectedRowWithQuery executes a query inside a given transaction (if you have one)
func (service *service) ExecuteEnsuringOneAffectedRowWithQuery(dbc *DBContext, query *Query) error {
	stmt, params, err := query.GetStatementAndParams()
	if err != nil {
		return err
	}
	return service.ExecuteEnsuringOneAffectedRow(dbc, stmt, params...)
}

// ExecuteEnsuringOneAffectedRow executes a query inside a given transaction (if you have one)
func (service *service) ExecuteEnsuringOneAffectedRow(dbc *DBContext, query string, params ...interface{}) error {
	dbr, err := service.Execute(dbc, query, params...)
	if err != nil {
		return err
	}
	if dbr.AffectedRows() != 1 {
		return fmt.Errorf("unable to insert or update: %d", dbr.AffectedRows())
	}
	return nil
}

// ExecuteWithQuery executes a query inside a given transaction (if you have one)
func (service *service) ExecuteWithQuery(dbc *DBContext, query *Query) (*DBResult, error) {
	stmt, params, err := query.GetStatementAndParams()
	if err != nil {
		return nil, err
	}
	return service.Execute(dbc, stmt, params...)
}

// Execute executes a query inside a given transaction (if you have one)
func (service *service) Execute(dbc *DBContext, query string, params ...interface{}) (*DBResult, error) {
	// Result
	var res sql.Result

	// We have a db transaction?
	// We only support the case in which dbc.tx or dbc.dbConn have anything... if not
	// we threat everything like if we have a nil dbc.
	if dbc != nil && (dbc.tx != nil || dbc.dbConn != nil) {
		// We have a transaction?
		if dbc.tx != nil {
			// Prepare the query
			stmt, err := dbc.tx.PrepareContext(dbc.ctx, query)
			if err != nil {
				service.logMetric(logError, "execute", "dbc.tx.PrepareContext(dbc.ctx, query)", err)
				return nil, err
			}

			// Execute inside the transaction
			res, err = stmt.ExecContext(dbc.ctx, params...)
			if err != nil {
				service.logMetric(logError, "execute", "txstmt.ExecContext(dbc.ctx, params...)", err)
				return nil, err
			}
		} else if dbc.dbConn != nil {
			// Prepare the query
			stmt, err := service.db.PrepareContext(dbc.ctx, query)
			if err != nil {
				service.logMetric(logError, "execute", "service.db.PrepareContext(dbc.ctx, query)", err)
				return nil, err
			}
			defer stmt.Close()

			// Execute using the context
			res, err = stmt.ExecContext(dbc.ctx, params...)
			if err != nil {
				service.logMetric(logError, "execute", "stmt.ExecContext(dbc.ctx, params...)", err)
				return nil, err
			}
		} else {
			// Not possible
			return nil, fmt.Errorf("you have sent a dbc without tx or dbConn")
		}
	} else {
		// We don't have a connection
		stmt, err := service.db.Prepare(query)
		if err != nil {
			service.logMetric(logError, "execute", "service.db.Prepare(query)", err)
			return nil, err
		}
		defer stmt.Close()

		// Execute without context
		res, err = stmt.Exec(params...)
		if err != nil {
			service.logMetric(logError, "execute", "stmt.Exec(params...)", err)
			return nil, err
		}
	}

	// Get affected rows
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	// done
	return &DBResult{
		affectedRows: affectedRows,
	}, nil
}

// logMetric logs an error metric if the error is not null
func (service *service) logMetric(logType logType, operation string, detail string, err error) {
	// tags := new(godog.Tags).
	// 	Add("operation", operation).
	// 	Add("detail", detail)

	if err != nil {
		postgresError, ok := err.(*pq.Error)
		if ok && postgresError != nil {
			// tags = tags.Add("error", fmt.Sprintf("%s", postgresError.Code))
		}
	}

	// godog.RecordSimpleMetric(fmt.Sprintf("application.%s.db.service.%s", service.datadogMetricPrefix,
	// 	string(logType)), 1, tags.ToArray()...)
}
