package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - the struct for our comment
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewSercice - returns a new comment services
func NewSercice(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - retrieves comments by their id from database
func (service *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := service.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /articles/name/)
func (service *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := service.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - Adds a new comment to the database
func (service *Service) PostComment(comment Comment) (Comment, error) {
	if result := service.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates a comments by id with new comment info
func (service *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := service.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := service.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

// DeleteComment - deletes a comment from the database by id
func (service *Service) DeleteComment(ID uint) error {
	if result := service.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - retrieves all comments from the database
func (service *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := service.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}
