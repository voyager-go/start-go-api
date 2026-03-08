package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	Conf    *Yaml
	ConfEnv string
)

type Yaml struct {
	Server `yaml:"server"`
	Log    `yaml:"log"`
	Mysql  `yaml:"mysql"`
	Redis  `yaml:"redis"`
}

type Server struct {
	Mode            string `yaml:"mode"`
	DefaultPageSize int    `yaml:"defaultPageSize"`
	MaxPageSize     int    `yaml:"maxPageSize"`
	TokenExpire     int64  `yaml:"tokenExpire"`
	TokenKey        string `yaml:"tokenKey"`
	TokenIssuer     string `yaml:"tokenIssuer"`
	JwtSecret       string `yaml:"jwtSecret"`
}

type Log struct {
	Debug    string `yaml:"debug"`
	FileName string `yaml:"fileName"`
	DirPath  string `yaml:"dirPath"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}
type Redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	DbNum       int    `yaml:"dbNum"`
	LoginPrefix string `yaml:"loginPrefix"`
}

// InitConfig 初始化配置信息
func InitConfig() {
	configFile := fmt.Sprintf("config.%s.yaml", ConfEnv)
	yamlConf, err := os.ReadFile(configFile)
	if err != nil {
		exampleFile := "config.example.yaml"
		if _, e := os.Stat(exampleFile); e == nil {
			panic(fmt.Errorf("读取配置文件失败: %s。请复制 %s 为 %s 并修改配置后重试", err, exampleFile, configFile))
		}
		panic(fmt.Errorf("读取配置文件失败: %s。请先创建 %s（可参考仓库中的 config.example.yaml）", err, configFile))
	}
	// 根据当前环境的值来替换配置文件中的环境变量(配合docker)
	yamlConf = []byte(os.ExpandEnv(string(yamlConf)))
	c := &Yaml{}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败: %s", err))
	}
	Conf = c
}
