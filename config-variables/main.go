package main

import (
	"fmt"
	"config-variables/config"
)

func main() {
	fmt.Println("App", myconfig.AppName)
	fmt.Println("Db Url", myconfig.DbURL)
	fmt.Println("Port", myconfig.ServerPort)
}
