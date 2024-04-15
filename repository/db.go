package repository

import (
	"database/sql"
	"empty-api-struct/api_error"
	"empty-api-struct/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const (
	//connectionString = "root:example@tcp(127.0.0.1:3307)/test_db?parseTime=true"
	connectionString = "docker:docker@tcp(golang_db:3306)/test_db?parseTime=true"
)

func SetupDB() (*gorm.DB, error) {
	_, gormDB, err := setupSqlDBAndGormDB()
	if err != nil {
		return nil, err
	}

	gormDB.Debug()
	return gormDB, nil
}

func setupSqlDBAndGormDB() (*sql.DB, *gorm.DB, error) {
	mysqlDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, nil, api_error.New(http.StatusInternalServerError, "failed to connect to db").WithInternal(err)
	}
	if err := mysqlDB.Ping(); err != nil {
		return nil, nil, api_error.New(http.StatusInternalServerError, "failed to ping to db").WithInternal(err)
	}

	mysqlConfig := mysql.Config{
		Conn: mysqlDB,
	}
	sqlConnection := mysql.New(mysqlConfig)
	gormConfig := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	}
	db, err := gorm.Open(sqlConnection, gormConfig)
	if err != nil {
		return nil, nil, api_error.New(http.StatusInternalServerError, "failed to open gorm db").WithInternal(err)
	}

	if err := db.AutoMigrate(&models.Person{}); err != nil {
		return nil, nil, api_error.New(http.StatusInternalServerError, "failed to automigrate db").WithInternal(err)
	}

	return mysqlDB, db, nil
}
