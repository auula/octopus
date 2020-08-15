package controller

import (
	"github.com/gin-gonic/gin"
)

const (
	staticOld   = "/semantic-ui/"
	staticSrc   = "./resource/static/semantic-ui"
	assetOld    = "/assets/"
	assetSrc    = "./resource/static/assets"
	templateSrc = "./resource/template/*"
)

// View Page
type view struct {
	r *gin.Engine
}

// NewView return ViewPageGroup
func NewView(gin *gin.Engine) *view {
	view := &view{r: gin}
	view.init()
	view.mappingView()
	return view
}

func (v *view) init() {
	//加载静态资源，例如网页的css、js
	//router.Static("/static", "./static")
	//加载静态资源，一般是上传的资源，例如用户上传的图片
	//router.StaticFS("/assets/img", http.Dir("/dashboard/assets/img/"))
	//加载单个静态文件
	//router.StaticFile("/favicon.ico", "./static/favicon.ico")
	v.r.Static(staticOld, staticSrc)
	v.r.Static(assetOld, assetSrc)
	v.r.LoadHTMLGlob(templateSrc)
}

func (v *view) mappingView() {
	v.r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.tmpl", nil)
	})
	v.r.GET("/hosts", func(c *gin.Context) {
		c.HTML(200, "../dashboard/nodes.html", nil)
	})
	v.r.GET("/config", func(c *gin.Context) {
		c.HTML(200, "../dashboard/config.html", nil)
	})
	v.r.GET("/system", func(c *gin.Context) {
		c.HTML(200, "../dashboard/system.html", nil)
	})
	v.r.NoRoute(func(c *gin.Context) {
		c.String(200, "https://github.com/higker/octopus")
	})
}
