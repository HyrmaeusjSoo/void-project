package mysql

import (
	"errors"
	"void-project/internal/model"
	"void-project/internal/repository/driver"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{db: driver.MySQL}
}

// 消息列表
func (m *MessageRepository) GetList() ([]model.Message, error) {
	var list []model.Message
	if err := m.db.Preload(clause.Associations, omitUserSensitivity).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// 查询消息
func (m *MessageRepository) GetById(id uint) *model.Message {
	msg := &model.Message{}
	m.db.First(msg, id)
	return msg
}

// 新增消息
func (m *MessageRepository) Create(msg *model.Message) error {
	tx := m.db.Create(msg)
	if tx.RowsAffected == 0 {
		return errors.New("新增失败")
	}
	return nil
}

// 删除消息
func (m *MessageRepository) Delete(id uint) error {
	if tx := m.db.Delete(&model.Message{}, id); tx.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return nil
}
