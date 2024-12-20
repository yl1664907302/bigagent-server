package user

import (
	"bigagent_server/db/mysqldb"
	"bigagent_server/model"
	responses "bigagent_server/web/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// Login godoc
// @Summary 用户登录接口
// @Description 处理用户登录请求
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param loginForm body model.User true "登录表单"
// @Router /v1/login [post]
func (*UserApi) Login(c *gin.Context) {
	// 获取请求参数
	var loginForm model.User
	// 使用ShouldBind方法将请求上下文（c）中的表单数据绑定到LoginForm实例
	if err := c.ShouldBind(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	user, err := mysqldb.LoginUser(loginForm.Username, loginForm.Password)
	if err != nil {
		log.Print(err)
		responses.ResponseApp.FailWithDetailed(c, "用户登入失败", map[string]any{
			"code": http.StatusInternalServerError,
		})

	} else {
		responses.ResponseApp.LoginSuccessDetailed(c, "登入成功！", map[string]any{
			"username":    user.Username,
			"role":        user.Role,
			"roleId":      user.RoleId,
			"permissions": user.Permissions,
		})

	}

}
