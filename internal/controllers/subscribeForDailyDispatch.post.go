package controllers

import (
	"net/http"
	"subscription-api/config"
	pb_ds "subscription-api/pkg/grpc/dispatch_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type subscribeParams struct {
	Email string `json:"email"`
}

const USD_UAH_DISPATCH_ID = "f669a90d-d4aa-4285-bbce-6b14c6ff9065"

func SubscribeForDailyDispatch(ctx Context, ds pb_ds.DispatchServiceClient) {
	var params subscribeParams
	if err := ctx.BindJSON(&params); err != nil {
		config.Log().Debug("failed to bind subscribe params from json: " + err.Error())
		ctx.String(http.StatusBadRequest, "invalid data provided")
		return
	}

	_, err := ds.SubscribeFor(ctx.Context(), &pb_ds.SubscribeForRequest{
		Email:      params.Email,
		DispatchId: USD_UAH_DISPATCH_ID,
	})
	if status.Code(err) == codes.AlreadyExists {
		ctx.Status(http.StatusConflict)
		return
	}
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
