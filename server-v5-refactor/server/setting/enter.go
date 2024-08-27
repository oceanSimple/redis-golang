package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"server-v5-refactor-server/static/output"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./setting")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(output.Error() + "Read config file error: " + err.Error())
	} else {
		fmt.Println(output.Info() + "Read config file success")
	}
}

func GetDefault(sth string) any {
	return viper.Get(sth)
}
