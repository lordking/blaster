package common

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

// InitConfig 读取并初始化配置文件。返回配置文件路径或错误。
// 如果appName != "" 和 cfgFile == ""，在/etc/[appName]/、$HOME/.[appName]/、./config/ `.`目录下需找。
// 如果appName == "" 和 cfgFile != ""，指定配置文件的具体路径。
// 如果appName != "" 和 cfgFile != ""，优先使用cfgFile
func InitConfig(appName, cfgFile string) (string, error) {

	//寻找配置文件
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

		if appName == "" {
			return "", fmt.Errorf("Not found appName or cfgFile")
		}

		viper.SetConfigName("config")
		viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
		viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", appName))
		viper.AddConfigPath("./config/")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
	}

	//导入配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	return viper.ConfigFileUsed(), nil
}

// ReadConfigKey 从读取的配置中寻找符合key的配置，并赋值到配置对象(config)定义的变量中。
// 如果config有环境变量定义，那么以环境变量设置为准
func ReadConfigKey(key string, obj interface{}) error {

	var err error

	if key == "" {
		err = viper.Unmarshal(obj)
	} else {
		err = viper.UnmarshalKey(key, obj)
	}
	if err != nil {
		return err
	}

	err = ReadEnv(obj)
	return err
}

//ReadEnv 读取struct内env的标签内容，转为读取环境变量
func ReadEnv(obj interface{}) error {

	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	} else {
		fmt.Errorf("This object is nil")
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		key := field.Tag.Get("env")
		value := os.Getenv(key)

		if key != "" && value != "" {
			val.Field(i).SetString(value)
		}

	}

	return nil
}
