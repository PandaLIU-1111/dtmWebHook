package Config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// 定义conf类型
// 类型里的属性，全是配置文件里的属性
type NotifyConfig struct {
	Type   string `yaml:"type"`
	Token  string `yaml:"token"`
	Secret string `yaml:"secret"`
}

type Configs struct {
	Configs []NotifyConfig `yaml:"configs"`
}

func InitConfig() *Configs {
	// Init config
	var c Configs
	//读取yaml配置文件
	configs := c.getConf()
	return configs
}

// 读取Yaml配置文件,
// 并转换成conf对象
func (c *Configs) getConf() *Configs {
	//应该是 绝对地址
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
	yamlFile, err := os.ReadFile(dir + "/config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
