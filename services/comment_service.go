package services

import (
	"github.com/watarui/go-intermediate-practice/apperrors"
	"github.com/watarui/go-intermediate-practice/models"
	"github.com/watarui/go-intermediate-practice/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to insert record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
