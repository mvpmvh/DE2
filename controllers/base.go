package controllers

import (
	"github.com/mvpmvh/beego"
	"github.com/mvpmvh/i18n"
	"strings"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	// Redirect to make URL clean.
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

// setLangVer sets site language version.
func (this *baseController) setLangVer() bool {
	isNeedRedir := false
	hasCookie := false

	// 1. Check URL arguments.
	lang := this.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "en-US"
		isNeedRedir = false
	}

	this.Lang = lang
	this.Data["Lang"] = lang
	this.Data["CurLang"] = i18n.GetDescriptionByLang(lang)

	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", lang, 1<<31-1, "/")
	}

	langDescriptions := i18n.ListLangDescriptions()
	otherLangs := make([]string, 0, len(langDescriptions)-1)

	for _, langDesc := range langDescriptions {
		if this.Data["CurLang"] != langDesc {
			otherLangs = append(otherLangs, langDesc)
		}
	}

	this.Data["OtherLangs"] = otherLangs

	return isNeedRedir
}
