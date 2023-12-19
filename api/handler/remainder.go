package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Samandarxon/examen_3-month/clinics/config"
	"github.com/Samandarxon/examen_3-month/clinics/models"
	"github.com/gin-gonic/gin"
)

// CreateRemainder godoc
// @ID				create_remainder
// @Router		/remainder [POST]
// @Summary		Create Remainder
// @Description	Create Remainder
// @Tags			Remainder
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateRemainder	true	"CreateRemainderRequestBody"
// @Success		201		{object}	Response{data=models.Remainder}	"RemainderBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateRemainder(c *gin.Context) {

	var remainder models.CreateRemainder
	err := c.ShouldBindJSON(&remainder)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(remainder)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.Remainder().Create(ctx, remainder)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdRemainder godoc
// @ID				get_by_id_remainder
// @Router		/remainder/{id} [GET]
// @Summary		GetById Remainder
// @Description	GetById Remainder
// @Tags			Remainder
// @Accept		json
// @Produce		json
// @Param	 id path string		true	"GetByIdRemainderRequestBody"
// @Success		200		{object}	Response{data=models.Remainder}	"RemainderBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdRemainder(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.Remainder().GetById(c, models.RemainderPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListRemainder godoc
// @ID				GetList_remainder
// @Router		/remainder [GET]
// @Summary		GetList Remainder
// @Description	GetList Remainder
// @Tags			Remainder
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param search query string false "offset"
// @Success		200		{object}	Response{data=models.GetListRemainderResponse} "RemainderBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListRemainder(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	search := c.Query("search")
	resp, err := h.strg.Remainder().GetList(c, models.GetListRemainderRequest{Offset: offset, Limit: limit, Search: search})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateRemainder godoc
// @ID update_baranch
// @Router 			/remainder/{id} [PUT]
// @Summary 		Update Remainder
// @Description Update Remainder
// @Tags 				Remainder
// @Accept 			json
// @Produce 		json
// @Param id path string true "RemainderPrimaryKey_ID"
// @Param object body models.UpdateRemainder true "UpdateRemainderBody"
// @Success  200 {object} Response{data=models.Remainder} "Updated Remainder"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateRemainder(c *gin.Context) {

	var (
		baranch = models.UpdateRemainder{}
		id      = c.Param("id")
	)
	err := c.ShouldBindJSON(&baranch)
	fmt.Println(err)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	baranch.Id = id
	fmt.Println(id)
	resp, err := h.strg.Remainder().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Remainder does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteRemainder godoc
// @ID delete_baranch
// @Router 			/remainder/{id} [DELETE]
// @Summary	 		Delete Remainder
// @Description Delete Remainder
// @Tags 				Remainder
// @Accept 			json
// @Produce 		json
// @Param id path string true "DeleteRemainderPath"
// @Success  200 {object} Response{data=string} "Deleted Remainder"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteRemainder(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.Remainder().Delete(c, models.RemainderPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "Remainder does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
