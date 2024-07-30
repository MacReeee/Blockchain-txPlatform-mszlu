package dao

import (
	"common/msdb"
	"common/msdb/gorms"
	"context"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) Save(ctx context.Context, mem *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).
		Where("mobile_phone=?", phone).Limit(1).
		Take(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}
