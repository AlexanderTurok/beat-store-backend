package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createArtist(c *gin.Context) {
	accountId, err := getAccountId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.Artist.Create(accountId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getArtistByToken(c *gin.Context) {

}

func (h *Handler) getArtistById(c *gin.Context) {

}

func (h *Handler) getAllArtists(c *gin.Context) {

}

func (h *Handler) deleteArtist(c *gin.Context) {

}
