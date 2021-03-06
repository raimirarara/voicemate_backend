package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/repository"
)

type (
	// TagPersistence is a persistence to preserve tags.
	TagPersistence struct {
		Connection *gorm.DB
	}
)

// NewTagPersistence creates a new Tag persistence.
func NewTagPersistence() repository.TagRepository {
	return TagPersistence{
		Connection: getDBConnection(),
	}
}

// FindTagByID gets tag from DB.
func (tagPersistence TagPersistence) FindTagByID(id uint64) (model.Tag, error) {
	tag := model.Tag{}

	result := tagPersistence.Connection.New().
		First(&tag, id)

	if result.RecordNotFound() {
		return tag, nil
	}
	if result.Error != nil {
		return tag, result.Error
	}
	return tag, nil
}

// GetAll gets tags from DB.
func (tagPersistence TagPersistence) GetTags(name string) ([]model.Tag, error) {
	tags := []model.Tag{}

	query := tagPersistence.Connection.New().Table("tags")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	result := query.Find(&tags)

	if result.RecordNotFound() {
		return tags, nil
	}
	if result.Error != nil {
		return tags, result.Error
	}
	return tags, nil
}

// CreateTag
func (tagPersistence TagPersistence) CreateTag(name string) (model.Tag, error) {
	tag := model.Tag{
		Name: name,
	}

	result := tagPersistence.Connection.New().Create(&tag)
	if result.RecordNotFound() {
		return tag, nil
	}
	if result.Error != nil {
		return tag, result.Error
	}

	return tag, nil
}

// DeleteTag
func (tagPersistence TagPersistence) DeleteTag(id uint64) error {
	tag := model.Tag{
		ID: id,
	}

	result := tagPersistence.Connection.New().Delete(&tag)
	if result.RecordNotFound() {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
