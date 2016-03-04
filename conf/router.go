package conf

import (
	"github.com/deluan/gosonic/api"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/deluan/gosonic/controllers"
)

func init() {
	mapEndpoints()
	mapControllers()
	mapFilters()
}

func mapEndpoints() {
	ns := beego.NewNamespace("/rest",
		beego.NSRouter("/ping.view", &api.PingController{}, "*:Get"),
		beego.NSRouter("/getLicense.view", &api.GetLicenseController{}, "*:Get"),
		beego.NSRouter("/getMusicFolders.view", &api.GetMusicFoldersController{}, "*:Get"),
		beego.NSRouter("/getIndexes.view", &api.GetIndexesController{}, "*:Get"),
		beego.NSRouter("/getMusicDirectory.view", &api.GetMusicDirectoryController{}, "*:Get"),
		beego.NSRouter("/getCoverArt.view", &api.GetCoverArtController{}, "*:Get"),
		beego.NSRouter("/stream.view", &api.StreamController{}, "*:Get"),
		beego.NSRouter("/download.view", &api.StreamController{}, "*:Get"),
		beego.NSRouter("/getUser.view", &api.UsersController{}, "*:GetUser"),
		beego.NSRouter("/getAlbumList.view", &api.GetAlbumListController{}, "*:Get"),
	)
	beego.AddNamespace(ns)

}

func mapControllers() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sync", &controllers.SyncController{})

	beego.ErrorController(&controllers.MainController{})
}

func mapFilters() {
	var ValidateRequest = func(ctx *context.Context) {
		c := &api.BaseAPIController{}
		c.Ctx = ctx
		api.Validate(c)
	}

	beego.InsertFilter("/rest/*", beego.BeforeRouter, ValidateRequest)
}