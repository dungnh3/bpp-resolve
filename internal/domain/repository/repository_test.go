package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type TestRepositorySuite struct {
	suite.Suite
	repo  IRepository
	mock  sqlmock.Sqlmock
	sqlDB *sql.DB
	db    *gorm.DB
}

func TestRepository(t *testing.T) {
	suite.Run(t, &TestRepositorySuite{})
}

func (suite *TestRepositorySuite) SetupTest() {
	var err error

	suite.sqlDB, suite.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(suite.T(), err)

	suite.db, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      suite.sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(suite.T(), err)

	suite.repo = NewRepository(suite.db, nil)
}
