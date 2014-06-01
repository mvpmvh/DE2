package main

import (
	"github.com/mvpmvh/beego"
	"github.com/mvpmvh/i18n"
	"os"
	_ "source.discoveryeducation.com/users/mhatch/repos/de2/routers"
	"source.discoveryeducation.com/users/mhatch/repos/de2/taglib"
	"strings"
)

func main() {
	initLogger()
	initLocales()
	initTemplateFuncs()
	beego.Run()
}

func initLogger() {
	beego.SetLevel(beego.LevelInfo)
	os.Mkdir("./log", os.ModePerm)
	beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
}

func initLocales() {
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")

	for i, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, names[i], "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func initTemplateFuncs() {
	beego.AddFuncMap("params", taglib.Params)
	beego.AddFuncMap("append", taglib.Append)
	beego.AddFuncMap("i18n", i18n.Tr)
}
