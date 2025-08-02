package configs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	file, err := os.Open(path + "/.env")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir .env: %w", err)
	}
	defer file.Close()

	cfg := &conf{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Ignorar linhas vazias e comentários
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // ou tratar erro
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Setar na env (opcional)
		os.Setenv(key, value)

		// Popular struct manualmente
		switch key {
		case "DB_DRIVER":
			cfg.DBDriver = value
		case "DB_HOST":
			cfg.DBHost = value
		case "DB_PORT":
			cfg.DBPort = value
		case "DB_USER":
			cfg.DBUser = value
		case "DB_PASSWORD":
			cfg.DBPassword = value
		case "DB_NAME":
			cfg.DBName = value
		case "WEB_SERVER_PORT":
			cfg.WebServerPort = value
		case "GRPC_SERVER_PORT":
			cfg.GRPCServerPort = value
		case "GRAPHQL_SERVER_PORT":
			cfg.GraphQLServerPort = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erro na leitura do arquivo: %w", err)
	}

	return cfg, nil
}

// func LoadConfig(path string) (*conf, error) {
// 	_, err := os.Stat(".env")
// 	if err != nil {
// 		fmt.Println("Erro: .env não encontrado")
// 	} else {
// 		fmt.Println(".env encontrado!")
// 	}

// 	var cfg *conf
// 	viper.SetConfigName("app_config")
// 	viper.SetConfigType("env")
// 	viper.AddConfigPath(path)
// 	viper.SetConfigFile(".env")
// 	viper.AutomaticEnv()
// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = viper.Unmarshal(&cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cfg, err
// }
