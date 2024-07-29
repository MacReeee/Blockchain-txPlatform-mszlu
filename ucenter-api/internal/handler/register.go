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

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
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
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)
	res := NewRes.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, res)
}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	res := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, res)
}
