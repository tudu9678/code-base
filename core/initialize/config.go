package initialize

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"myapp/core/common"

	g "gorm.io/gorm"
	gLog "gorm.io/gorm/logger"
)

type Config struct {
	Port     string `json:"PORT"`
	ENV      string `json:"ENV"`
	LogLevel string `json:"LOG_LEVEL"`
}

type Postgres struct {
	Username    string   `json:"POSTGRES_USERNAME"`
	Password    string   `json:"POSTGRES_PASSWORD"`
	Instance    string   `json:"POSTGRES_INSTANCE"`
	DB          string   `json:"POSTGRES_DB_NAME"`
	Port        string   `json:"POSTGRES_PORT"`
	Proxy       string   `json:"POSTGRES_PROXY"`
	Drive       string   `json:"POSTGRES_DRIVE"`
	MaxIdle     string   `json:"POSTGRES_MAX_IDLE_CONNS"`
	MaxOpen     string   `json:"POSTGRES_MAX_OPEN_CONNS"`
	MaxLifetime string   `json:"POSTGRES_MAX_LIFETIME"`
	GormConfig  g.Config //gorm config for open connection
}

type PostgresReplica struct {
	Username    string   `json:"POSTGRES_REPLICA_USERNAME"`
	Password    string   `json:"POSTGRES_REPLICA_PASSWORD"`
	Instance    string   `json:"POSTGRES_REPLICA_INSTANCE"`
	DB          string   `json:"POSTGRES_REPLICA_DB_NAME"`
	Port        string   `json:"POSTGRES_REPLICA_PORT"`
	Proxy       string   `json:"POSTGRES_REPLICA_PROXY"`
	Drive       string   `json:"POSTGRES_REPLICA_DRIVE"`
	MaxIdle     string   `json:"POSTGRES_REPLICA_MAX_IDLE_CONNS"`
	MaxOpen     string   `json:"POSTGRES_REPLICA_MAX_OPEN_CONNS"`
	MaxLifetime string   `json:"POSTGRES_REPLICA_MAX_LIFETIME"`
	GormConfig  g.Config //gorm config for open connection
}

func GormConfig(logMode string) g.Config {
	m := gLog.Silent
	switch strings.ToLower(logMode) {
	case common.DBLogModeError:
		m = gLog.Error
	case common.DBLogModeWarn:
		m = gLog.Warn
	case common.DBLogModeInfo:
		m = gLog.Info
	default:
		m = gLog.Silent
	}

	return g.Config{
		Logger:      gLog.Default.LogMode(m),
		PrepareStmt: true,
	}
}

type JwtConfig struct {
	Issuer        string `json:"JWT_ISSUER"`
	TokenLifeTime string `json:"JWT_TOKEN_LIFE_TIME"`
	PrivateKey    string `json:"JWT_PRIVATE_KEY"`
	PublicKey     string `json:"JWT_PUBLIC_KEY"`
}

type Redis struct {
	Endpoint string `json:"REDIS_ENDPOINT"`
	Password string `json:"REDIS_PASSWORD"`
}

type RedisSentinel struct {
	Endpoint   string `json:"REDIS_SENTINEL_ENDPOINT"`
	MasterName string `json:"REDIS_SENTINEL_MASTER_NAME"`
	Password   string `json:"REDIS_SENTINEL_PASSWORD"`
	DB         string `json:"REDIS_SENTINEL_DB"`
	Privatekey string `json:"REDIS_SENTINEL_TLS_PRIVATE_KEY_FILE"`
	CA         string `json:"REDIS_SENTINEL_TLS_CA_FILE"`
	Pem        string `json:"REDIS_SENTINEL_TLS_PEM_FILE"`
}

func (r *RedisSentinel) GetDB() int {
	i, err := strconv.Atoi(r.DB)
	if err != nil {
		log.Fatalf("get redis-db:%v", err)
	}
	return i
}

type ElasticSearch struct {
	URL      string `json:"ELASTIC_SEARCH_ENDPOINT"`
	Username string `json:"ELASTIC_SEARCH_USERNAME"`
	Password string `json:"ELASTIC_SEARCH_PASSWORD"`
}

func (es *ElasticSearch) GetURLs() []string {
	return strings.Split(es.URL, ",")
}

func LoadConfiguration(exts ...interface{}) error {
	for _, ex := range exts {
		rt := reflect.TypeOf(ex)
		if rt.Kind() != reflect.Ptr {
			return fmt.Errorf("invalid type config")
		}

		values := reflect.ValueOf(ex).Elem()
		for i := 0; i < values.NumField(); i++ {
			if rt.Elem().Field(i).IsExported() {
				env := values.Type().Field(i).Tag.Get("json")
				if len(env) > 0 {
					values.Field(i).SetString(os.Getenv(env))
				}
			}
		}
	}
	return nil
}

func LoadConfigurationWithTagPrefix(pre string, ex interface{}) error {
	rt := reflect.TypeOf(ex)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("invalid type config")
	}

	values := reflect.ValueOf(ex).Elem()
	for i := 0; i < values.NumField(); i++ {
		if rt.Elem().Field(i).IsExported() {
			env := pre + values.Type().Field(i).Tag.Get("json")
			if len(env) > 0 {
				values.Field(i).SetString(os.Getenv(env))
			}
		}
	}

	return nil
}
