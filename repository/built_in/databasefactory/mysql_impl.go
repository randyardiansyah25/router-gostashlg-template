package databasefactory

import (
	"database/sql"
	"fmt"
	"os"
	"router-gostashlg-template/entities/app"
	"time"

	"github.com/randyardiansyah25/libpkg/security/aes"
	"github.com/randyardiansyah25/libpkg/util/env"

	_ "github.com/go-sql-driver/mysql"
)

func newMysqlImpl() Database {
	return &mysqlImpl{}
}

type mysqlImpl struct {
	conn   *sql.DB
	prefix string
}

func (m *mysqlImpl) Connect() (er error) {

	//Mysql Connection String Format : [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//Example : root:123456@tcp(127.0.0.1:3306)/employees?charset=utf8
	pass := env.GetString(m.prefix + "mysql.password")
	if pass != "" {
		pvKey := []byte(app.PrivateKey)
		pass, er = aes.Decrypt(pvKey, pvKey, pass)
		if er != nil {
			return er
		}
	}
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv(m.prefix+"mysql.username"),
		pass,
		os.Getenv(m.prefix+"mysql.address"),
		os.Getenv(m.prefix+"mysql.port"),
		os.Getenv(m.prefix+"mysql.name"),
	)
	maxPool := env.GetInt(m.prefix+"mysql.maxpoolsize", 10)
	maxIdleConn := env.GetInt(m.prefix+"mysql.maxidleconn", 10)
	maxLifeTime := env.GetInt(m.prefix+"mysql.maxconnlifetime", 40)
	if m.conn, er = sql.Open(DRIVER_MYSQL, connectionString); er != nil {
		return er
	}

	m.conn.SetMaxOpenConns(maxPool)
	m.conn.SetMaxIdleConns(maxIdleConn)
	m.conn.SetConnMaxLifetime(time.Minute * time.Duration(maxLifeTime))

	return nil
}

func (m *mysqlImpl) Ping() error {
	return m.conn.Ping()
}

func (m *mysqlImpl) GetConnection() interface{} {
	return m.conn
}

func (m *mysqlImpl) GetDriverName() string {
	return DRIVER_MYSQL
}

func (m *mysqlImpl) SetEnvironmentVariablePrefix(prefix string) {
	m.prefix = prefix
}

func (m *mysqlImpl) Close() {
	_ = m.conn.Close()
}
