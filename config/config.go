package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Conf *Yaml

type Yaml struct {
	Server `yaml:"server"`
	Log    `yaml:"log"`
	Mysql  `yaml:"mysql"`
	Redis  `yaml:"redis"`
}

type Server struct {
	Port            string `yaml:"port"`
	Mode            string `yaml:"mode"`
	DefaultPageSize uint64 `yaml:"defaultPageSize"`
	MaxPageSize     uint64 `yaml:"maxPageSize"`
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

// init 初始化配置信息
func init() {
	var defaultConfigFile = fmt.Sprintf("config.%s.yaml", os.Getenv("SERVER_ENV"))
	configFile := flag.String("c", defaultConfigFile, "help config path")
	flag.Parse()
	yamlConf, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
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
