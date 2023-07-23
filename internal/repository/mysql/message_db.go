package mysql

import (
	"errors"
	"void-project/internal/model"
	"void-project/internal/model/base"
	"void-project/internal/repository"
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

// 消息列表 - 清理多余信息后
func (m *MessageRepository) GetListClean(uId, targetId uint, pager base.Pager) ([]model.Message, int, error) {
	var list []model.Message
	total, err := repository.Paginate(m.db.
		Where("from_id=? AND target_id=?", uId, targetId).
		Or("from_id=? AND target_id=?", targetId, uId).
		Joins("LEFT JOIN `user` AS f ON f.id=message.from_id").
		Joins("LEFT JOIN `user` AS t ON t.id=message.target_id").
		Select("message.*, f.`name` AS from_name, t.`name` AS target_name").
		Order("send_time DESC"), &list, pager)
	if err != nil {
		return nil, 0, err
	}
	return list, int(total), nil
}

// 查询消息
func (m *MessageRepository) GetById(id uint) *model.Message {
	msg := &model.Message{}
	m.db.First(msg, id)
	return msg
}

// 新增消息
func (m *MessageRepository) Create(msg *model.Message) error {
	tx := m.db.Omit(clause.Associations).Create(msg)
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
