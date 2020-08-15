package controller

import "github.com/gin-gonic/gin"


// Api  Api standard
type api struct {
	r *gin.Engine
}
// NewApi return Api
func NewApi(r *gin.Engine) *api {
	return &Api{r:r}
}