package logic

import (
	"common/tools"
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter/internal/domain"

	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/svc"
)

const RegisterRedisKey = "REGISTER::"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	//1 检查人机验证是否通过
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip)
	if !isVerify {
		return &register.RegRes{}, errors.New("人机验证不通过")
	}
	//2 校验验证码
	redisVal := ""
	err := l.svcCtx.Cache.GetCtx(context.Background(), RegisterRedisKey+in.Phone, &redisVal)
	if err != nil {
		return nil, errors.New("redis验证码获取失败")
	}
	if in.Code != redisVal {
		return nil, errors.New("验证码错误")
	}
	//3 验证通过，执行注册逻辑，验证手机号是否注册
	logx.Info("人机校验通过....")
	mem, err := l.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if err != nil {
		return nil, err
	}
	if mem != nil {
		return nil, errors.New("此手机号已被注册")
	}
	//logx.Info("第三部分完成")

	//4. 生成member模型，存入数据库
	err = l.MemberDomain.Register(
		context.Background(),
		in.Phone,
		in.Password,
		in.Username,
		in.Country,
		in.SuperPartner,
		in.Promotion)
	if err != nil {
		return nil, errors.New("注册失败")
	}
	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(in *register.CodeReq) (*register.NoRes, error) {
	code := tools.Rand4Num()
	logx.Infof("验证码为: %s", code)
	//通过短信平台发送验证码
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterRedisKey+in.Phone, code, 5*time.Minute)
	return &register.NoRes{}, err
}
