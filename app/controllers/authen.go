package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"github.com/sutthirak/labyrinth/data"
)

type AuthenticationController struct {
	*revel.Controller
}

func (c AuthenticationController) LogIn() revel.Result {

	total := data.CountUser()
	if total == 0 {
		return c.Redirect("/setup")
	}

	return c.Render()
}

func (c AuthenticationController) LogOut() revel.Result {
	c.Session["user"] = ""
	return c.Redirect("/login")
}

func (c AuthenticationController) Auth(email string,password string) revel.Result {

	//check log in
	c.Validation.Required(email).Message("Please enter your email")
	c.Validation.Required(password).Message("Please enter your password")

	if email !="" && password !="" && !data.Auth(email,password) {
		c.Validation.Error("email or password was wrong")
	}

	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/login")
	}else{
		c.Session["user"] = email
		return c.Redirect("/admin")
	}
}

func (c AuthenticationController) SetUp() revel.Result{

	//if already had user redirect to Login
	total := data.CountUser()
	if total >= 1 {
		return c.Redirect("/login")
	}

	return c.Render()
}

func (c AuthenticationController) SetUpSave(email string,
										password string,
										confirmPassword string,
										firstName string,
										lastName string) revel.Result {

	//if already had user redirect to Login
	total := data.CountUser()
	if total >= 1 {
		return c.Redirect("/login")
	}

	//if already set up. redirect to login
	c.Validation.Email(email).Message("Please enter a correct email")
	c.Validation.Required(password).Message("Please enter a password")
	c.Validation.MaxSize(password,16).Message("Password must be at most 16 characters long")
	c.Validation.MinSize(password,8).Message("Password must be at least 8 characters long")
	c.Validation.Required(confirmPassword).Message("Please confirm your password")
	if password != confirmPassword {
		c.Validation.Error("Password and confirm password must be the same")
	}
	c.Validation.Required(firstName).Message("Please enter your name")
	c.Validation.Required(lastName).Message("Please enter your last name")
	
	if c.Validation.HasErrors() {

		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/setup")
	}

	user := data.User{}
	user.Email = email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = string(hashedPassword)
	user.FirstName = firstName
	user.LastName = lastName

	data.SaveUser(&user)
	c.Flash.Success("Set up account success. Please log in to Labyrinth")
	return c.Redirect("/login")
}