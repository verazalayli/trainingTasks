package middle

import "log"

//**Задача:**
//Приложение читает:
//
//- порт из флага `-port` (по умолчанию 8080);
//
//- лог-уровень из env `LOG_LEVEL` (info/debug);
//
//- печатает итоговую конфигурацию и стартует HTTP-сервер.

type Config struct {
	Port     int
	LogLevel string
}

func loadConfig() Config {
	// TODO: считать флаги и env, вернуть Config
	return Config{}
}

func main() {
	cfg := loadConfig()
	log.Printf("starting on :%d, logLevel=%s", cfg.Port, cfg.LogLevel)
	// TODO: запустить сервер на cfg.Port
}
