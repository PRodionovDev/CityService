package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type SiteService struct {

}

func NewSiteService() SiteService {
    return SiteService{}
}

func (s *SiteService) Index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
}
