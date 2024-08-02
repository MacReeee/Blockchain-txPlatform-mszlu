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
	memberRepo repo.MemberRepo
}

func (d MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	mem, err := d.memberRepo.FindByPhone(ctx, phone)
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
	mem.Avatar = "https://avatars.githubusercontent.com/u/54505023?s=400&u=965ad764e3376feb406d986a4c679aa62b9dafcb&v=4"
	mem.TransactionStatus = 1
	err := d.memberRepo.Save(ctx, mem)
	if err != nil {
		logx.Error(err)
		return errors.New("数据库异常")
	}
	return nil
}

func (d *MemberDomain) UpdateLoginCount(ctx context.Context, id int64, step int) {
	err := d.memberRepo.UpdateLoginCount(ctx, id, step)
	if err != nil {
		logx.Error(err)
	}
}

func (d *MemberDomain) FindMemberById(ctx context.Context, memberId int64) (*model.Member, error) {
	id, err := d.memberRepo.FindMemberById(ctx, memberId)
	if err == nil && id == nil {
		return nil, errors.New("用户不存在")
	}
	return id, err
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		memberRepo: dao.NewMemberDao(db),
	}
}
