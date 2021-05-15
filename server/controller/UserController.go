package controller

import (
	"kokow/go/common"
	"kokow/go/dto"
	"kokow/go/model"
	"kokow/go/response"
	"kokow/go/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	// })
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(model.User))},
	})
}
func Login(ctx *gin.Context) {
	//获取参数
	// phone := ctx.PostForm("phone")
	// password := ctx.PostForm("password")
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	phone := requestUser.Phone
	password := requestUser.Password
	//数据验证
	if len(phone) != 11 {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "手机号码不正确"})
		response.Fail(ctx, nil, "手机号码不正确")
		return
	}
	if len(password) < 6 {
		response.Fail(ctx, nil, "密码太短")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 403, "msg": "密码太短"})
		return
	}

	//判断手机号是否存在
	var user model.User
	common.DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 403, "msg": "用户不存在"})
		response.Fail(ctx, nil, "用户不存在")
		return
	}
	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// ctx.JSON(http.StatusBadRequest, gin.H{"code": 403, "msg": "密码错误"})
		response.Fail(ctx, nil, "密码错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "token 生成失败")
		log.Printf("token generate error:%v,", err)
		return
	}
	//发放token
	response.Success(ctx, gin.H{"token": token}, "登录成功")

	//返回结果
}
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	// var requestMap = make(map[string]string)
	// json.NewDecoder(ctx.Request.Body).Decode(&requestMap)
	var requestUser = model.User{}
	ctx.Bind(&requestUser)
	name := requestUser.Name
	phone := requestUser.Phone
	password := requestUser.Password
	// name := ctx.PostForm("name")
	// phone := ctx.PostForm("phone")
	// password := ctx.PostForm("password")
	//数据验证
	if len(phone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码不正确")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 423, nil, "密码太短")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, phone, password)
	//判断手机号是否存在
	if isPhoneExist(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 424, nil, "手机已存在")
		return
	}
	//创建用户
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}
	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: string(hasePassword),
	}
	DB.Create(&newUser)
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "token 生成失败")
		log.Printf("token generate error:%v,", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "注册成功")
	// response.Success(ctx, nil, "注册成功")
}
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ? ", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
