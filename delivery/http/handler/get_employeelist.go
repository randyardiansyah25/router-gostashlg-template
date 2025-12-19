package handler

import (
	"net/http"
	"router-gostashlg-template/entities/app"
	"router-gostashlg-template/entities/common/logger"
	"router-gostashlg-template/usecase"

	"github.com/gin-gonic/gin"
)

func GetEmployeListHandler(ctx *gin.Context) {
	ucase := usecase.NewEmployeeUsecase()
	data, er := ucase.GetEmployeeList()
	if er != nil {
		if er == app.ErrNoRecord {
			ctx.String(http.StatusNoContent, "record not found.")
		} else {
			logger.PrintError(er.Error())
			ctx.String(http.StatusInternalServerError, "internal service error")
		}
	} else {
		ctx.JSON(http.StatusOK, data)
	}
}
