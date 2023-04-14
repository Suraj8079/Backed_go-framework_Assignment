package handler

import (
	"encoding/json"
	"net/http"

	"employee/internal/svc"
	"employee/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Employee struct {
	Name   string
	Email  string
	DOB    string
	Mobile string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func PostDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		resp := Response{
			Code: http.StatusOK,
		}
		user := Employee{} //initialize empty user
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			resp.Code = http.StatusBadRequest
			resp.Message = err.Error()

		} else {
			resp.Message = "Record Saved Successfully Name: " + user.Name + " Email: " + user.Email + " DOB: " + user.DOB + " Mobile Number:" + user.Mobile
		}

		httpx.Ok(w)
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}