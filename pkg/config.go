package pkg

type Configs struct {
	Threads int    `json:"threads"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Mysql   Mysql  `json:"mysql"`
	Redis   Redis  `json:"redis"`
}

type Mysql struct {
	MysqlHost string `json:"mysqlHost"`
	MysqlUser string `json:"mysqlUser"`
	MysqlPwd  string `json:"mysqlPwd"`
	MysqlDb   string `json:"mysqlDb"`
	MysqlPort int    `json:"mysqlPort"`
}

type Redis struct {
	RedisHost string `json:"redisHost"`
	RedisPwd  string `json:"redisPwd"`
	RedisDb   int    `json:"redisDb"`
	RedisPort int    `json:"redisPort"`
	RedisKey  string `json:"redisKey"`
}
