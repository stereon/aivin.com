package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stereon/aivin.com/common"
	"github.com/stereon/aivin.com/dto"
	"github.com/stereon/aivin.com/model"
	"github.com/stereon/aivin.com/response"
	"github.com/stereon/aivin.com/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context)  {
	db := common.InitDb()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	fmt.Println(name,telephone,password)
	//验证参数
	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,200,nil,"手机号必须要11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,200,nil,"密码必须大于6位")
		return
	}

	//判断手机号是否存在
	if IsPhoneExsit(db ,telephone){
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}

	//创建用户
	if len(name) ==0 {
		name = util.RandomString(10)
	}

	hashpassword ,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx,http.StatusInternalServerError,500,nil,"用户已经存在")
		return
	}

	user := model.User{Fusername: name, Ftelphone: telephone, Fpassword: string(hashpassword)}

	db.Table("t_user_info").Create(&user)
	response.Success(ctx,nil,"注册成功")
}

func IsPhoneExsit( db *gorm.DB, telephone string) bool {

	var user *model.User
	sql := "SELECT Fusername, Ftelphone, Fpassword FROM t_user_info WHERE ftelphone = ?"
	fmt.Println(sql)
	db.Raw(sql, telephone).Scan(&user)
	fmt.Printf("user:%#v\n",&user)
	fmt.Println(user)
	if user != nil {
		return true
	}
	return false
}

func Login(ctx *gin.Context)   {
	db := common.InitDb()
	var user *model.User
	//获取参数
	telephone := ctx.PostForm("telphone")
	password := ctx.PostForm("password")

	//数据验证
	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码必须大于6位")
		return
	}

	sql := "SELECT Fusername, Ftelphone, Fpassword FROM t_user_info WHERE ftelphone = ?"
	db.Raw(sql, telephone).Scan(&user)

	if user == nil {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}

	// 判断密码是否存在
	if err := bcrypt.CompareHashAndPassword([]byte(user.Fpassword),[]byte(password)) ; err != nil {
		response.Response(ctx,http.StatusBadRequest,400,nil,"密码错误")
		return
	}

	// 发放token
	token,err := common.ReleaseToken(user)
	fmt.Println(token)
	if err != nil {
		response.Response(ctx,http.StatusInternalServerError,500,nil,"密码错误")
		log.Printf("token generate error :%v\n", err)
		return
	}

	response.Success(ctx,gin.H{"token":token},"登录成功")
}

func Info(ctx *gin.Context)  {
	user , _ := ctx.Get("user")
	name := dto.UserTodo(user.(model.User))
	fmt.Printf("v的值为: %v, v的类型为: %T\n", name, name)
	fmt.Printf("v的值为: %v, v的类型为: %T\n", user, user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"code":200,"data":gin.H{"user":dto.UserTodo(user.(model.User))}},
	)
}