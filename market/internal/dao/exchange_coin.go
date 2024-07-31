package dao

import (
	"common/msdb"
	"common/msdb/gorms"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"market/internal/model"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func (e *ExchangeCoinDao) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := e.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return data, err
}

func (e *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	if err != nil {
		logx.Error(err)
	}
	return
}

func NewExchangeCoinDao(db *msdb.MsDB) *ExchangeCoinDao {
	return &ExchangeCoinDao{
		conn: gorms.New(db.Conn),
	}
}
