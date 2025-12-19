package handler

import (
	"net/http"
	"router-gostashlg-template/entities"
	"router-gostashlg-template/entities/app"
	"router-gostashlg-template/entities/common/logger"
	"router-gostashlg-template/usecase"

	"github.com/gin-gonic/gin"
)

func GetEmployee(ctx *gin.Context) {
	payload := entities.EmployeeFilter{}
	var er error
	if ctx.ContentType() == "application/json" {
		er = ctx.BindJSON(&payload)
	} else {
		er = ctx.Bind(&payload)
	}
	if er != nil {
		ctx.String(http.StatusBadRequest, er.Error())
	}

	ucase := usecase.NewEmployeeUsecase()
	employee, er := ucase.GetEmployee(payload.Id)
	if er != nil {
		if er == app.ErrDuplicateEntry {
			ctx.String(http.StatusConflict, "Data karyawan sudah tersedia!")
		} else {
			logger.PrintError(er.Error())
			ctx.String(http.StatusInternalServerError, "internal service error")
		}
	} else {
		ctx.JSON(http.StatusOK, employee)
	}
}
