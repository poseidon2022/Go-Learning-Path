package controllers

import (
	"blog_filtration_feature/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FiltrationController struct {
	FiltrationUseCase 				domain.BlogFiltrationUseCase
}

func (fc *FiltrationController) FilterBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var filterReq domain.BlogRequest
		if err := c.BindJSON(&filterReq); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
			return
		}

		filtrationResponse, err := fc.FiltrationUseCase.FilterBlog(&filterReq)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"filtered_blogs" : filtrationResponse})
	} 
}
