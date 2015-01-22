package app

import (
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/app/controllers"
)

func checkUser(c *revel.Controller) revel.Result{

	if(c.Session["user"] == ""){
		//check set up status
		status := revel.Config.BoolDefault("setup",false)
		if !status {
			return c.Redirect("/setup")
		}

		return c.Redirect("/login")
	}
	return nil
}

func init() {

	//register interceptor function.
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.Admin{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.SceneController{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.ObjectController{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.AnimationController{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.SpriteController{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.MaterialController{})
	revel.InterceptFunc(checkUser, revel.BEFORE,&controllers.ResourceController{})

	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

}
