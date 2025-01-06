package config

type Database struct {
	MysqlHost         string `json:"mysqlhost" yaml:"mysqlhost"`
	MysqlPort         int    `json:"mysqlport" yaml:"mysqlport"`
	MysqlUser         string `json:"mysqluser" yaml:"mysqluser"`
	MysqlPassword     string `json:"mysqlpassword" yaml:"mysqlpassword"`
	MysqlDatabasename string `json:"mysqldatabasename" yaml:"mysqldatabasename"`
}
