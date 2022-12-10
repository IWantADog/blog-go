package conf

type mysqlConf struct {
	Uri string
}

type redisConf struct {
	Uri string
}

type ServerConf struct {
	Mysql *mysqlConf
	Redis *redisConf
}
