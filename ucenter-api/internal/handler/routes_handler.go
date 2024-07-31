package handler

import (
	"ucenter-api/internal/midd"
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

	assetGroup := r.Group()
	assetGroup.Use(midd.Auth(serverCtx.Config.JWT.AccessSecret))
	asset := NewAssetHandler(serverCtx)
	assetGroup.Post("/uc/asset/wallet/:coinName", asset.FindWalletBySymbol)
	assetGroup.Post("/uc/asset/wallet", asset.FindWallet)
	assetGroup.Post("/uc/asset/wallet/reset-address", asset.ResetAddress)
	assetGroup.Post("/uc/asset/transaction/all", asset.FindTransaction)

	approveGroup := r.Group()
	approve := NewApproveHandler(serverCtx)
	approveGroup.Use(midd.Auth(serverCtx.Config.JWT.AccessSecret))
	approveGroup.Post("/uc/approve/security/setting", approve.SecuritySetting)

	withdrawGroup := r.Group()
	withdraw := NewWithdrawHandler(serverCtx)
	withdrawGroup.Use(midd.Auth(serverCtx.Config.JWT.AccessSecret))
	withdrawGroup.Post("/uc/withdraw/support/coin/info", withdraw.QueryWithdrawCoin)
	withdrawGroup.Post("/uc/mobile/withdraw/code", withdraw.SendCode)
	withdrawGroup.Post("/uc/withdraw/apply/code", withdraw.WithdrawCode)
	withdrawGroup.Post("/uc/withdraw/record", withdraw.Record)
}
