package router

import "github.com/gin-gonic/gin"

func Router() {

	r := gin.New()

	r.Run(":8088")

}
