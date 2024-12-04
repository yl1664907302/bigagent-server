package user

import (
	"bigagent_server/db/mysqldb"
	"bigagent_server/model"
	responses "bigagent_server/web/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserApi struct{}

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
		responses.FailWithDetailed(c, "用户登入失败", map[string]any{
			"code": http.StatusInternalServerError,
		})

	} else {
		responses.LoginSuccessDetailed(c, "登入成功！", map[string]any{
			"username":    user.Username,
			"role":        user.Role,
			"roleId":      user.RoleId,
			"permissions": user.Permissions,
		})

	}

}
