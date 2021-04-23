package utils

import (
	"course/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// 读取配置文件
func UnmarshalConf(conf *config.TotalConf) error {
	env, ok := os.LookupEnv("CUR_ENV")
	if !ok {
		return fmt.Errorf("Not Found ENV: CUR_ENV")
	}
	var confByt []byte
	var err error
	if env == "test" {
		confByt, err = ioutil.ReadFile("./conf/myConf_test.yml")
	} else {
		confByt, err = ioutil.ReadFile("./file/myConf_pro.yml")
	}
	if err != nil {
		return err
	}
	return yaml.Unmarshal(confByt, conf)
}