package config

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	TokenVK string
	ChatsID map[int64]int64
}

var instance *Config
var once sync.Once

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		CreateConfig()
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config := GetInstance()
	chats_ids := viper.Get("chats_id").([]interface{})
	config.ChatsID = make(map[int64]int64)
	for k, v := range chats_ids {
		config.ChatsID[int64(v.(int))] = int64(k)
	}
	config.TokenVK = viper.Get("vk_api_token").(string)
}
func CreateConfig() {
	var config_content = []byte(`vk_api_token : token
chats_id : 
- 1
- 2
- 3`)
	f, _ := os.Create("config.yaml")
	w := bufio.NewWriter(f)
	w.WriteString(string(config_content) + "\n")
	w.Flush()
}

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}
