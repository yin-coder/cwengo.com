package admin

import (
	"cwengo.com/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/Tpl/T.category.tpl"
	this.Data["isCategory"] = true
	this.Data["Username"] = this.GetSession("username")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Post() {
	newCategory := this.Input().Get("NewCategory")
	if newCategory == "" {
		this.Data["CategoryNull"] = "分类名不能为空"
		return
	}
	err := models.AddCategory(newCategory)
	if err != nil {
		beego.Error(err)
	}
	this.Ctx.Redirect(302, "/admin/category")
}
