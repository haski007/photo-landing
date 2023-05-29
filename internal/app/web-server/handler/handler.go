package handler

import (
	"github.com/haski007/photo-landing/internal/app/web-server/handler/representation"
	"github.com/haski007/photo-landing/internal/app/web-server/usecase"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ContactUseCase *usecase.ContactUseCase
}

func NewHandler(usecase *usecase.ContactUseCase) *Handler {
	return &Handler{
		ContactUseCase: usecase,
	}
}

func (h *Handler) HandleContactForm(c *gin.Context) {
	var req representation.ContactRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		logrus.WithError(err).Println("bind json error")
		return
	}

	logrus.WithFields(map[string]interface{}{
		"name":    req.Name,
		"email":   req.Email,
		"message": req.Message,
	}).Println("got post request")

	err := h.ContactUseCase.SendContactForm(req.Name, req.Email, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		logrus.WithError(err).Println("send contact form via grpc error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Form successfully submitted",
	})
}
