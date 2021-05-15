package controller

import (
	"kokow/go/model"
	"kokow/go/repository"
	"kokow/go/response"
	"kokow/go/vo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	// DB *gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	// db := common.GetDB()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "验证错误,分类名称必填")
		return
	}
	// ctx.Bind(&requestCategory)
	// if requestCategory.Name == "" {
	// 	response.Fail(ctx, nil, "分类名称必填")
	// }
	// category := model.Category{Name: requestCategory.Name}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		// response.Fail(ctx, nil, "创建失败")
		panic(err)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "创建成功")

}
func (c CategoryController) Edit(ctx *gin.Context) {
	var requestCategory vo.CategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, nil, "验证错误,分类名称必填")
		return
	}
	//获取url参数强转int
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	// var updateCategory model.Category
	// if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
	// 	response.Fail(ctx, nil, "分类不存在")
	// }
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
	}
	//更新分类
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, nil, "更新失败")
	}
	response.Success(ctx, gin.H{"category": category}, "更新成功")
	// c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	// response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}
func (c CategoryController) Remove(ctx *gin.Context) {
	//获取url参数强转int
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}
	response.Success(ctx, nil, "删除成功")

}
func (c CategoryController) View(ctx *gin.Context) {
	//获取url参数强转int
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	viewCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
	}

	// var category model.Category
	// if c.DB.First(&category, categoryId).RecordNotFound() {
	// 	response.Fail(ctx, nil, "分类不存在")
	// }

	response.Success(ctx, gin.H{"category": viewCategory}, "")
}
