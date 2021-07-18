package actions

func (as *ActionSuite) Test_Users_New() {
	res := as.HTML("/register").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	res := as.HTML("/register").Post(map[string]string{
		"Name":                 "Mark Smith",
		"Email":                "mark@example.com",
		"Password":             "password",
		"PasswordConfirmation": "password",
	})
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
