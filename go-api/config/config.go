package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		cfg,
	}
	if err := c.initConfig(); err != nil {
		return err
	}

	c.initLog()
	//监控配置文件并热加载
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 指定了,则解析配置文件
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设定配置文件的格式之类
	viper.SetConfigType("yaml")
	// 读取环境变量,并且只会读取指定前缀的环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

func (c *Config) watchConfig() {
	// 监听配置文件变更, 如果有变更,可以使进程不重启的情况下,更新最新配置.
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
