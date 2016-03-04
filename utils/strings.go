package utils

import (
	"github.com/astaxie/beego"
	"strings"
)

func NoArticle(name string) string {
	articles := strings.Split(beego.AppConfig.String("ignoredArticles"), " ")
	for _, a := range articles {
		n := strings.TrimPrefix(name, a+" ")
		if n != name {
			return n
		}
	}
	return name
}