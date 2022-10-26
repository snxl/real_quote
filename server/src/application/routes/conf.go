package routes

import "github.com/gin-gonic/gin"

func Routes(eng *gin.Engine) {
	eng.SetTrustedProxies(nil)

	quotesRoutes(eng)
}
