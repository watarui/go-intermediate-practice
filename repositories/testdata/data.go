package testdata

import "github.com/watarui/go-intermediate-practice/models"

var ArticleTestData = []models.Article{
	{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  2,
	},
	{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  4,
	},
}

var CommentTestData = []models.Comment{
	{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
