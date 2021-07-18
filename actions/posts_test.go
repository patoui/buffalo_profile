package actions

func (as *ActionSuite) Test_Posts_List() {
	as.LoadFixture("one published post, one not published post")
	res := as.HTML("/blog").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "This is published post title #1")
}
