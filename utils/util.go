package utils

import (
	"course/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// 读取配置文件
func UnmarshalConf(conf *config.TotalConf) error {
	env, ok := os.LookupEnv("CUR_ENV")
	var confByt []byte
	var err error
	if !ok {
		log.Println("Not Found ENV: CUR_ENV")
		return fmt.Errorf("NO CUR_ENV")
	}
	log.Println("ENV: " + env)
	if env == "test" {
		confByt, err = ioutil.ReadFile("./conf/myConf_test.yml")
	} else {
		confByt, err = ioutil.ReadFile("./conf/myConf_pro.yml")
	}
	if err != nil {
		return err
	}
	return yaml.Unmarshal(confByt, conf)
}