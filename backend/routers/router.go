// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/ws", &controllers.WebSocketController{}, "get:Get")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/visitor/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/owner/user",
			beego.NSInclude(&controllers.OwnerController{}),
		),
		beego.NSNamespace("/owner/company",
			beego.NSInclude(&controllers.CompanyController{}),
		),
		beego.NSNamespace("/geocoder/cords",
			beego.NSInclude(&controllers.GeoController{}),
		),
	)

	beego.AddNamespace(ns)
}
