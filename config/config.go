package config

import (
	"fmt"
	"strconv"
	"github.com/spf13/viper"
)

type databaseConfig struct {
	host        string
	port        int
	username    string
	password    string
	name        string
	maxPoolSize int
}

func getIntOrPanic(key string) int {
        v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(err)
	}
        return v
}

// Reads config variables from application.yml which is
// present at the root of this project
func LoadConfig() {
	//viper.SetDefault("LOG_LEVEL", "debug")
	//viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// fills databaseConfig struct with values
func newDatabaseConfig() *databaseConfig {
	return &databaseConfig {
		host:        viper.GetString("DB_HOST"),
		port:        getIntOrPanic("DB_PORT"),
		name:        viper.GetString("DB_NAME"),
		username:    viper.GetString("DB_USER"),
		password:    viper.GetString("DB_PASSWORD"),
		maxPoolSize: getIntOrPanic("DB_POOL"),
	}
}

func DatabaseConfig() *databaseConfig {
	return newDatabaseConfig()
}

func (dc *databaseConfig) ConnectionString() string {
        return fmt.Sprintf("user=%s dbname=%s password='%s' host=%s sslmode=disable", dc.username, dc.name, dc.password, dc.host)
}

func (dc *databaseConfig) DbName() string {
	return dc.name
}
