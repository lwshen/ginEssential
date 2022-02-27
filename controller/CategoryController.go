package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"slw.app/ginessential/common"
	"slw.app/ginessential/model"
	"slw.app/ginessential/response"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{
		DB: db,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
		return
	}

	c.DB.Create(&requestCategory)

	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误,分类名称必填")
		return
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	// 更新分类
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "删除失败,请重试")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
