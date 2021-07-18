package actions

import "net/http"

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "My name is Patrique Ouimet :D")
}
