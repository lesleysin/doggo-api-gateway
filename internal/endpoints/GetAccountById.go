package endpoints

import (
	"github.com/DogGoOrg/doggo-api-gateway/internal/helpers"
	"github.com/gin-gonic/gin"
)

func GetAccountById(ctx *gin.Context) {
	conn, err := grpcController.ConnGrpc("ACCOUNT_HOST")

	if err != nil {
		helpers.Error5xx(ctx, err)
		return
	}

	defer conn.Close()

	// client := Account.NewAccountClient(conn)

	// res, err := client.Login(context.TODO(), &Account.LoginRequest{})
}
