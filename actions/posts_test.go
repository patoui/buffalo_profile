package actions

func (as *ActionSuite) Test_Posts_List() {
	as.LoadFixture("one published post, one not published post")
	res := as.HTML("/blog").Get()

	as.Equal(200, res.Code)

	rb := res.Body.String()
	as.Contains(rb, "This is published post title #1")
	as.NotContains(rb, "This is published post title #2")
}
