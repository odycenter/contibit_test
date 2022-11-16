package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"contibit_test/pkg"

	"github.com/go-redis/redis"
)

var cf pkg.Configs
var THREADS int
var NAME string
var VERSION string

var MYSQL_HOST string
var MYSQL_USER string
var MYSQL_PWD string
var MYSQL_DB string
var MYSQL_PORT int
var DB_CONNECT string

var REDIS_HOST string
var REDIS_PWD string
var REDIS_DB int
var REDIS_PORT int
var REDIS_KEY string
var CTX context.Context
var RWMUTEX sync.RWMutex
var LAST_API_TIMESTAMP int64
var COUNTER int64

var COIN_FLOW chan string = make(chan string, 400)
var LAST_TRANSFER_TIME chan int64 = make(chan int64, 400)
var TRANSFER_NUMBER chan int = make(chan int, 400)
var TRANSFER_CHAN chan pkg.TransferChannel = make(chan pkg.TransferChannel, 400)

func init() {
	ReadConfig(&cf)
	THREADS = cf.Threads
	NAME = cf.Name
	VERSION = cf.Version

	MYSQL_HOST = cf.Mysql.MysqlHost
	MYSQL_USER = cf.Mysql.MysqlUser
	MYSQL_PWD = cf.Mysql.MysqlPwd
	MYSQL_DB = cf.Mysql.MysqlDb
	MYSQL_PORT = cf.Mysql.MysqlPort
	DB_CONNECT = MYSQL_USER + ":" + MYSQL_PWD + "@tcp(" + MYSQL_HOST + ":" + strconv.Itoa(MYSQL_PORT) + ")/" + MYSQL_DB

	REDIS_HOST = cf.Redis.RedisHost
	REDIS_PWD = cf.Redis.RedisPwd
	REDIS_DB = cf.Redis.RedisDb
	REDIS_PORT = cf.Redis.RedisPort
	REDIS_KEY = cf.Redis.RedisKey
}

func ReadConfig(cfg *pkg.Configs) {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}
	configFileName, _ = filepath.Abs(configFileName)
	log.Printf("Loading config: %v", configFileName)

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}

func RedisClient(add string, port int, passwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     add + ":" + strconv.Itoa(port),
		Password: passwd, // no password set
		DB:       db,     // use default DB
	})
	return client
}
