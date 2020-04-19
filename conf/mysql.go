package conf

var mysqlConf = map[string]map[string]string{
	"default": {
		"host":     "",
		"password": "",
		"port":     "",
		"user":     "",
		"dbname":   "",
	},
}

func GetAllMysqlConf() map[string]map[string]string {
	return mysqlConf
}

func GetMysqlConf(key string) map[string]string {
	if key == "" {
		key = "default"
	}
	return mysqlConf[key]
}
