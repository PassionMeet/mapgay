package umshandlers

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/db"
	"github.com/cmfunc/jipeng/pkg/password"
	"github.com/cmfunc/jipeng/pkg/usertoken"
	"github.com/gin-gonic/gin"
)

type LoginParam struct {
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
}

// ums login
func Login(c *gin.Context) {
	param := LoginParam{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("Login %s", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	row, err := db.GetUmsAccount(c, param.Account)
	if err != nil {
		log.Printf("Login param:%+v %s", param, err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	if !password.CheckPasswordHash(param.Password, row.Password) {
		log.Printf("Login auth failed param:%+v row.password:%s", param, row.Password)
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	// gen new account token
	token := usertoken.Gen()
	cache.SetUMSToken(c, param.Account, token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
