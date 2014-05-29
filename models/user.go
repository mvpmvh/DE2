/*contains all the modesl for the web application */

package models

type User struct {
	Id        *int32 `form:"-"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Email     string `form:"email"`
	FirstName string `form:"firstName"`
	LastName  string `form:"lastName"`
}
