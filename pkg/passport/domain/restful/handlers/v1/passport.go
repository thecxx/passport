package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thecxx/passport/pkg/passport/service"
)

func RegisterHandler(ctx *gin.Context) {

}

func LoginHandler(ctx *gin.Context) {

	// param: login_method
	loginMethod := ctx.PostForm("login_method")

	// param: type
	accountType := ctx.PostForm("type")

	// param: account
	accountName := ctx.PostForm("account")

	account, err := service.Account.Login(context.Background(), accountType, accountName, loginMethod, map[string]string{})
	if err != nil {
		return
	}

	ctx.JSON(0, account)
	return
}
