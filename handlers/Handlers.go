package handlers

import (
	"fmt"
	"net/http"

	"github.com/GhostPowerShell/RegisterLoginStudy/data"

	"github.com/GhostPowerShell/RegisterLoginStudy/helpers"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/login.html")
	fmt.Fprintf(response, body)
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")

	redirectTarget := "/"

	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		// Database check for user data!
		_userIsValid := data.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, response)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/register"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)

}

// for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("templates/register.html")
	fmt.Fprintf(response, body)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	email := r.FormValue("email")
	passwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	_username, _email, _passwd, _confirmPwd := false, false, false, false

	_username = !helpers.IsEmpty(username)
	_email = !helpers.IsEmpty(email)
	_passwd = !helpers.IsEmpty(passwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _username && _email && _passwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", username)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", passwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	username := GetUserName(request)
	if !helpers.IsEmpty(username) {
		var indexBody, _ = helpers.LoadFile("templates/index.html")
		fmt.Fprintf(response, indexBody, username)
	} else {
		http.Redirect(response, request, "/", 302)
	}

}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}

func SetCookie(username string, response http.ResponseWriter) {
	value := map[string]string{
		"name": username,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func GetUserName(request *http.Request) (username string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			username = cookieValue["name"]
		}
	}
	return username
}
