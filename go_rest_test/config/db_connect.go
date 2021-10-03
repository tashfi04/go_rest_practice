package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go_rest_test/helpers"
)

var Db *sql.DB

func Init() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}


	dbSource := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8",  configuration.Database.DBUser, configuration.Database.DBPassword, configuration.Server.Port, configuration.Database.DBName)

	Db, err = sql.Open("mysql", dbSource)

	helpers.Catch(err)

	fmt.Println("Connected to MySql Database")

	//defer Db.Close()

}
