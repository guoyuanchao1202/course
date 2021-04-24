package handle

import (
	"context"
	"course/dal"
	"course/request"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func DoLogin(ctx context.Context, loginReq *request.LoginReq) (*dal.User, error) {
	user, err := dal.QueryUserByUserName(loginReq.UserName)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("query user [ ", loginReq.UserName, " ] failed: ", err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		log.Println("no user [ ", loginReq.UserName, " ]")
		return nil, fmt.Errorf("no user [%v], please retry !", loginReq.UserName)
	}
	if user.PassWord == loginReq.PassWord {
		log.Println(loginReq.UserName, " login success")
		return user, nil
	}
	log.Println(loginReq.UserName, " login failed")
	return nil, fmt.Errorf("login failed: passWd is incorrect !")
}
