package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LBAddr          string `json:"lb_addr"`
	OssAddr         string `json:"oss_addr"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}

func GetAccessKeyId() string {
	return configuration.AccessKeyId
}

func GetAccessKeySecret() string {
	return configuration.AccessKeySecret
}
