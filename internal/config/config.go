package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Config ...
type Config struct {
	Http struct {
		Port               string        `json:"port"`
		ReadTimeout        time.Duration `json:"readTimeout"`
		WriteTimeout       time.Duration `json:"writeTimeout"`
		MaxHeaderMegabytes int           `json:"maxHeaderBytes"`
	}

	JWT struct {
		SigningKey      string        `json:"signingkey"`
		ExpiredDuration time.Duration `json:"expiredduration"`
	} `json:"jwt_token"`

	Database struct {
		DBMS     string `json:"DBMS"`
		Username string `json:"username"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		Dbname   string `json:"Dbname"`
	}
}

// Init ...
func Init(path string) (*Config, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var conf Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &conf)

	return &conf, nil
}
