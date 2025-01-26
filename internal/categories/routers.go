package categories

import (
	"context"
	"fmt"
	"github.com/benedictngth/NUSphereBackend/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Categories(router *gin.RouterGroup, categoriesService CategoriesService) {
	router.POST("/create", CreateChildCategoryHandler(categoriesService))
	router.GET("", GetCategoriesHandler(categoriesService))
	// // router.PUT("/edit/:id", EditCategoryHandler(categoriesService))
	// router.DELETE("/delete/:id", DeleteCategoryHandler(categoriesService))
}

func CreateChildCategoryHandler(categoriesService CategoriesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req NewCategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.INVALID_INPUT})
			return
		}
		err := categoriesService.CreateChildCategory(context.Background(), req.Name, req.Description, req.ParentID, req.CreatedBy)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": CREATE_CATEGORY_FAILED})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": CREATE_CATEGORY_SUCCESS})
	}
}

func GetCategoriesHandler(categoriesService CategoriesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := categoriesService.GetCategories(context.Background())
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": GET_CATEGORIES_FAILED})
			return
		}
		fmt.Println(categories)
		c.JSON(http.StatusOK, categories)
	}
}

// func GetParentCategoriesHandler(categoriesService CategoriesService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		categories, err := categoriesService.GetParentCategories(context.Background())
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get categories"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, categories)
// 	}
// }

//implement if necessary

// func DeleteCategoryHandler(categoriesService CategoriesService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		publicID := c.Param("id")
// 		err := categoriesService.DeleteCategoryByPublicID(context.Background(), publicID)
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "category deletion failed"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
// 	}
// }

//implement if necessary
// func EditCategoryHandler(categoriesService CategoriesService) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		publicID := c.Param("id")
// 		var req EditCategoryRequest
// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
// 			return
// 		}
// 		err := categoriesService.EditCategory(context.Background(), publicID, req.Name, req.Description)
// 		if err != nil {
// 			c.Error(err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "category edit failed"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "category edited"})
// 	}
// }
