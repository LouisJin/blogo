package g

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/config/yaml"
	"github.com/beego/beego/v2/core/logs"
	yaml2 "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Site   SiteConfig   `yaml:"site"`
	Policy PolicyConfig `yaml:"policy"`
	Admin  AdminConfig  `yaml:"admin"`
}

type SiteConfig struct {
	Name  string `yaml:"name"`
	Desc  string `yaml:"desc"`
	Url   string `yaml:"url"`
	Beian string `yaml:"beian"`
}

type PolicyConfig struct {
	CommentInterval  int `yaml:"comment_interval"`
	ThumbsupInterval int `yaml:"thumbsup_interval"`
}

type AdminConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

const configPath = "conf/config.yaml"
const UserToken = "UserToken"

var GlobalConfig = new(Config)
var parse config.Configer

/**
* 获取全局基本配置
 */
func init() {
	config := new(yaml.Config)
	var err error
	parse, err = config.Parse(configPath)
	if err != nil {
		logs.Error("读取配置文件失败！", err)
		return
	}
	err = parse.Unmarshaler("", GlobalConfig)
	if err != nil {
		logs.Error("转换配置文件失败！", err)
		return
	}
}

func SetGlobalConfig(globalConfig *Config) bool {
	out, err := yaml2.Marshal(globalConfig)
	if err != nil {
		logs.Error("转换配置文件失败！", err)
		return false
	}
	err = ioutil.WriteFile(configPath, out, 0644)
	if err != nil {
		logs.Error("保存配置文件失败！", err)
		return false
	}
	GlobalConfig = globalConfig
	return true
}
