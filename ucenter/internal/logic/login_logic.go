package logic

import (
	"common/tools"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/types/login"
	"time"
	"ucenter/internal/domain"

	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/svc"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
	//1 检查人机验证是否通过
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip)
	if !isVerify {
		return &login.LoginRes{}, errors.New("人机验证不通过")
	}
	//2. 校验密码
	member, err := l.MemberDomain.FindByPhone(context.Background(), in.GetUsername())
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登陆失败")
	}
	if member == nil {
		return nil, errors.New("用户未注册")
	}
	password, salt := member.Password, member.Salt
	verify := tools.Verify(in.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("密码不正确")
	}
	//3. 登陆成功，生成token，提供给前端
	key := l.svcCtx.Config.JWT.AccessSecret
	expireTime := l.svcCtx.Config.JWT.AccessExpire

	token, err := l.getJwtToken(key, time.Now().Unix(), expireTime, member.Id)
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	//4.返回登陆所需信息
	loginCount := member.LoginCount + 1
	go func() {
		l.MemberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
	}()
	return &login.LoginRes{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount),
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
