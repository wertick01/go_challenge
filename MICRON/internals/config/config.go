package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Cfg struct { //наша структура для хранения конфигов, я полагаюсь на Viper для матчинга имен
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New() //создаем экземпляр нашего ридера для Env
	v.SetEnvPrefix("SERV") //префикс, все переменные нашего сервера должны теперь стартовать с SERV_ для того, чтобы не смешиваться с системными
	v.SetDefault("PORT", "8080")//ставим умолчальные настройки
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "")
	v.SetDefault("DBHOST", "")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "postgres")
	v.AutomaticEnv()//собираем наши переменные с системных

	var cfg Cfg

	err := v.Unmarshal(&cfg) //закидываем переменные в cfg после анмаршалинга
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string { //маленький метод для сборки строки соединения с БД
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}