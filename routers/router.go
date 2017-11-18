// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"pinter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),

		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),

		beego.NSNamespace("/onboarding",
			beego.NSInclude(
				&controllers.OnboardingController{},
			),
		),

		beego.NSNamespace("/soal",
			beego.NSInclude(
				&controllers.SoalController{},
			),
		),

		beego.NSNamespace("/download",
			beego.NSInclude(
				&controllers.DownloadController{},
			),
		),

		beego.NSNamespace("/jawaban",
			beego.NSInclude(
				&controllers.JawabanController{},
			),
		),

		beego.NSNamespace("/subcategory",
			beego.NSInclude(
				&controllers.SubcategoryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
