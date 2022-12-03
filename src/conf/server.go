package conf

type mysqlConf struct {
	Host   string
	Port   int32
	DBName string
}

type ServerConf struct {
	mysql *mysqlConf
}
