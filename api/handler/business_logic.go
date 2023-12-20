package handler

import (
	"fmt"
	"net/http"

	"github.com/Samandarxon/examen_3-month/clinics/models"
	"github.com/gin-gonic/gin"
)

// ClinetOverallReport godoc
// @ID				ClinetOverallReport
// @Router		/report_client [GET]
// @Summary		ClinetOverallReport
// @Description	ClinetOverallReport
// @Tags			OverallReport
// @Accept		json
// @Produce		json
// @Param date_from query string false "date_from"
// @Param date_to query string false "date_to"
// @Success		201		{object}	Response{data=models.Client}	"ClinetOverallReportBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) ClinetOverallReport(c *gin.Context) {
	from := c.Query("date_from")
	to := c.Query("date_to")
	fmt.Println("DATE>>>:", to, from)
	clent_resp, err := h.strg.Report().GetListReport(c, models.GetListClientReportRequest{
		DateFrom: from,
		DateTo:   to,
	})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}
	handleResponse(c, http.StatusOK, clent_resp)
}

// SaleOverallReport godoc
// @ID				SaleOverallReport
// @Router		/report_sale [GET]
// @Summary		SaleOverallReport
// @Description	SaleOverallReport
// @Tags			OverallReport
// @Accept		json
// @Produce		json
// @Success		201		{object}	Response{data=models.Client}	"SaleOverallReportBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) SaleOverallReport(c *gin.Context) {

	sale_resp, err := h.strg.Report().GetListSaleBranch(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}
	handleResponse(c, http.StatusOK, sale_resp)
}
