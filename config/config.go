package config

import "github.com/spf13/viper"

// DBConfig to store the DB configuration.
type DBConfig struct {
	UserName string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PW"`
	DBName   string `mapstructure:"DB_NAME"`
	DBIP     string `mapstructure:"DB_IP"`
	DBPort   string `mapstructure:"DB_PORT"`
}

// LoadDBConfig loads the DB parameters for connecting to the DB.
func LoadDBConfig(path string, filename string) (DBConfig, error) {
	var config DBConfig

	// specify the location of config file to viper
	viper.AddConfigPath(path)

	// tell viper to look for a config file with specific name
	viper.SetConfigName(filename)

	// specify the type/format of file for viper to look for
	// in this case, it is env format
	viper.SetConfigType("env")

	// read values from env vars to viper
	viper.AutomaticEnv()

	// start reading in the config/env vars values
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	// unmarshal decodes json into go struct
	// arg is the address of go struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	// fmt.Println(config, "config")

	return config, nil
}
