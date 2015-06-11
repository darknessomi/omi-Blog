package main

import (
	"github.com/astaxie/beego"
	"github.com/darknessomi/omi-Blog/g"
	_ "github.com/darknessomi/omi-Blog/routers"
)

func main() {
	g.InitEnv()
	beego.Run()
}
