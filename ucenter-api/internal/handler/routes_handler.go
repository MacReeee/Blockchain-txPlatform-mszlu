package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	registerGroup := r.Group()
	register := NewRegisterHandler(serverCtx)
	registerGroup.Post("/uc/register/phone", register.Register)
	registerGroup.Post("/uc/mobile/code", register.SendCode)

	loginGroup := r.Group()
	login := NewLoginHandler(serverCtx)
	loginGroup.Post("/uc/login", login.Login)
	loginGroup.Post("/uc/check/login", login.CheckLogin)
}
