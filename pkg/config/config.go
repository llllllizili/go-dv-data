package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"path"
	"runtime"
	"strings"
)

var GlobalConfig *Config

// Config is application global config
type Config struct {
	AppName      string       `mapstructure:"app-name"`     //应用名称
	LogPath      string       `mapstructure:"logging-path"` //应用名称
	TapdConfig   TapdConfig   `mapstructure:"tapd"`         // tapd信息
	DBConfig     DBConfig     `mapstructure:"database"`     // database信息
	GitlabConfig GitlabConfig `mapstructure:"gitlab"`       // database信息
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Init     bool   `mapstructure:"init"`
}

type TapdConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Created  string `mapstructure:"created"`
}
type GitlabConfig struct {
	Endpoint    string `mapstructure:"endpoint"`
	Token       string `mapstructure:"token"`
	MRProjectId []int  `mapstructure:"mrprojectid"`
	PPL         []PPL  `mapstructure:"ppl"`
}
type PPL struct {
	ID   int    `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

// Load is a loader to load config file.
func Load(configFilePath string) *Config {
	resolveRealPath(configFilePath)
	// 初始化配置文件
	if err := initConfig(); err != nil {
		panic(err)
	}

	return GlobalConfig
}

func initConfig() error {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APPLICATION")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析到struct
	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}
	log.Println("The application configuration file is loaded successfully!")
	return nil
}

// 如果未传递配置文件路径将使用约定的环境配置文件
func resolveRealPath(filepath string) {
	if filepath != "" {
		viper.SetConfigFile(filepath)
	} else {
		var abPath string
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			abPath = path.Dir(filename)
		} // 设置默认的config
		viper.AddConfigPath(abPath + "../../../conf")
		viper.SetConfigName("config")
	}
}

// AppOptions 用来接收应用启动时指定的参数
type AppOptions struct {
	ConfigFilePath string // 配置文件路径
}

func ResolveAppOptions(opt *AppOptions) {
	var configFilePath string

	flag.StringVar(&configFilePath,
		"c", "",
		"-c 选项用于指定要使用的配置文件")
	flag.Parse()

	opt.ConfigFilePath = configFilePath
}

func GetGlobleConfig() *Config {
	initOpt := &AppOptions{}
	ResolveAppOptions(initOpt)

	// 加载配置文件
	c := Load(initOpt.ConfigFilePath)

	return c
}
