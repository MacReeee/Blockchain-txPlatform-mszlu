package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginHandler struct {
	svcCtx *svc.ServiceContext
}

func NewLoginHandler(svcCtx *svc.ServiceContext) *LoginHandler {
	return &LoginHandler{
		svcCtx: svcCtx,
	}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req types.LoginReq
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	NewRes := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, NewRes.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	//获取ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewLoginLogic(r.Context(), h.svcCtx)
	resp, err := l.Login(&req)
	res := NewRes.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, res)
}
