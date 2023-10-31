package main

import (
	"fmt"
	"os"
	"sync"
)

type UserConfig struct {
	UserConfigDir string
}

var configInstance *UserConfig
var once sync.Once

func GetConfig() *UserConfig {
	once.Do(func() {
		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			fmt.Println("An error occurred:", err)
		}
		configInstance = &UserConfig{
			UserConfigDir: userConfigDir,
		}
	})
	return configInstance
}

func main() {
	config1 := GetConfig()
	fmt.Printf("User1 config directory: %s\n", config1.UserConfigDir)

	config2 := GetConfig()
	fmt.Printf("User2 config directory: %s\n", config2.UserConfigDir)
}
