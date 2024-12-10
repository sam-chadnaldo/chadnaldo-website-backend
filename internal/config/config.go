package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct{
	Env 			string 			`yaml:"env" env-default:"local"`
	PgSql 			PgSql 			`env-required:"true"`
	// AppSecret		string			`env:"APP_SECRET,required"`
	Http			HttpConfig		`yaml:"http" env-required:"true"`
	CoinGate 		CoinGateConfig	`yaml:"coingate" env-required:"true"`

}

type CoinGateConfig struct{
	BaseURL 		string			`yaml:"base_url" env-required:"true"`
	TestURL 		string 			`yaml:"test_url" env-required:"true"`
	ApiToken		string 			`env:"COINGATE_TOKEN,required"`
	Timeout			time.Duration	`yaml:"timeout" env-required:"true"`
	RetriesCount	int				`yaml:"retries_count" env-required:"true"`
}

type HttpConfig struct {
	Port	string			`yaml:"port"`
}

type PgSql struct {
	Host     string `env:"POSTGRES_HOST,required"`
	User     string `env:"POSTGRES_USER,required"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	DbName   string `env:"POSTGRES_DB,required"`
	Port     int    `env:"POSTGRES_PORT" env-default:"5432"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-default:"disable"`
}


func MustLoad()	*Config {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	path := fetchConfigFlag()

	if path == ""{
		panic("config path is empty")
	}

	if _,err := os.Stat(path); os.IsNotExist(err){
		panic("config path dose not exist: " + path)
	}

	var config Config

	if err:=cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	if err:=cleanenv.ReadEnv(&config); err != nil {
		panic("failed to read env: " + err.Error())
	}

	return &config
}

func fetchConfigFlag() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file") 
	flag.Parse()

	if res == ""{
		res = os.Getenv("CONFIG_PATH")	
	}

	return res
}