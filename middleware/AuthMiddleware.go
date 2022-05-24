package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stereon/aivin.com/common"
	"github.com/stereon/aivin.com/model"
	"github.com/stereon/aivin.com/response"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println(tokenString)
		//验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer"){
			response.Response(ctx,http.StatusUnauthorized,401,nil,"权限不足")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims,err := common.ParseToken(tokenString)
		fmt.Println(token,claims,err)
		if err != nil || !token.Valid {
			//ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			response.Response(ctx,http.StatusUnauthorized,401,nil,"权限不足")
			ctx.Abort()
			return
		}

		//获取验证后的claims 中的userid
		userId := claims.UserId
		db := common.InitDb()
		var user model.User
		sql := "SELECT Fusername, Ftelphone, Fpassword FROM t_user_info WHERE ftelphone = ?"
		db.Raw(sql, userId).Scan(&user)
		fmt.Println(user)
		if userId == "" {
			response.Response(ctx,http.StatusUnauthorized,401,nil,"权限不足")
			ctx.Abort()
			return
		}
		// 用户存在，将user信息写入上下文
		ctx.Set("user",user)
		ctx.Next()

	}
	
}