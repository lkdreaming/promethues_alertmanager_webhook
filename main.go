package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"promethues_alertmanager_webhook/config"
	"promethues_alertmanager_webhook/router"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := config.BootStrapInit()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	router.Router(&r.RouterGroup)
	r.Use(gin.Recovery())
	r.Run(fmt.Sprintf(":%d", config.AppConf.Server.Port)) //TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
}
