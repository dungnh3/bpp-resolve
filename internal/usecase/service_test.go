package usecase

import (
	"context"
	"fmt"
	"github.com/dungnh3/bpp-resolve/internal/domain/repository"
	"github.com/dungnh3/bpp-resolve/internal/dto"
	"github.com/dungnh3/bpp-resolve/pkg/database"
	"github.com/dungnh3/bpp-resolve/pkg/log"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"

	migrateV4 "github.com/golang-migrate/migrate/v4"
	// import mysql
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	// import file
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// import go_bin_data
	_ "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

var dbCfg = &database.MySQLConfig{
	Config: database.Config{
		Host:     "127.0.0.1",
		Database: "test",
		Port:     3307,
		Username: "root",
		Password: "secret",
		Options:  "parseTime=true",
	},
}

var logCfg = log.DefaultConfig()

type TestServiceSuite struct {
	suite.Suite
	db       *gorm.DB
	svc      *UseCase
	resource *dockertest.Resource
	pool     *dockertest.Pool
}

func TestService(t *testing.T) {
	suite.Run(t, &TestServiceSuite{})
}

func (suite *TestServiceSuite) SetupSuite() {
	var err error
	suite.pool, err = dockertest.NewPool("")
	require.NoError(suite.T(), err)

	opts := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "latest",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + dbCfg.Password,
			"MYSQL_DATABASE=" + dbCfg.Database,
		},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: dbCfg.Host, HostPort: fmt.Sprintf("%v", dbCfg.Port)},
			},
		},
		Cmd: []string{"--default-authentication-plugin=mysql_native_password"},
	}
	suite.resource, err = suite.pool.RunWithOptions(&opts)
	require.NoError(suite.T(), err)

	err = suite.pool.Retry(suite.connect)
	require.NoError(suite.T(), err)

	err = suite.migrateUp()
	require.NoError(suite.T(), err)

	logger := logCfg.MustBuildLogR()
	repo := repository.NewRepository(suite.db, logger)
	suite.svc = NewUseCase(logger, repo)
}

func (suite *TestServiceSuite) TearDownSuite() {
	suite.dropContainer()
}

func (suite *TestServiceSuite) SetupTest() {
	_, err := suite.svc.InitializeWager(context.Background(), &dto.CreateWagerDto{
		TotalWagerValue:   100,
		Odds:              40,
		SellingPercentage: 50,
		SellingPrice:      70,
	})
	require.NoError(suite.T(), err)
}

func (suite *TestServiceSuite) connect() error {
	var err error
	suite.db, err = gorm.Open(mysql.Open(dbCfg.DSN()), &gorm.Config{})
	return err
}

func (suite *TestServiceSuite) migrateUp() error {
	m, err := migrateV4.New("file:../../migrations", dbCfg.String())
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrateV4.ErrNoChange {
		return err
	}
	return nil
}

func (suite *TestServiceSuite) dropContainer() {
	err := suite.pool.Purge(suite.resource)
	require.NoError(suite.T(), err)
}
