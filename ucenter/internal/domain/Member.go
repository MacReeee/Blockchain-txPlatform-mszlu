package domain

import (
	"common/msdb"
	"common/tools"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func (d MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("数据库处理异常")
	}
	return mem, nil
}

func (d MemberDomain) Register(ctx context.Context,
	phone string,
	password string,
	username string,
	country string,
	partner string,
	promotion string) error {

	mem := model.NewMember()
	_ = tools.Default(mem)
	salt, pwd := tools.Encode(password, nil)
	logx.Info(len(pwd))
	mem.Username = username
	mem.Country = country
	mem.Password = pwd
	mem.MobilePhone = phone
	mem.FillSuperPartner(partner)
	mem.PromotionCode = promotion
	mem.MemberLevel = model.GENERAL
	mem.Salt = salt
	err := d.MemberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error(err)
		return errors.New("数据库异常")
	}
	return nil
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}
