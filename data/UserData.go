package data

func UserIsValid(username, password string) bool {
	_username, _password, _isValid := "admin", "admin", false

	if username == _username && password == _password {
		_isValid = true
	} else {
		_isValid = false
	}

	return _isValid
}
