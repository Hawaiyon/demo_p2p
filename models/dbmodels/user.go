// User model in db

package dbmodels

import (
	"fmt"
	"demo_p2p_bak/models"
	"encoding/json"
	"log"
)

// User user
// swagger:model User



func AddUser(user* models.User) (err error){
	affected, err := engine.Insert(user)
	if affected < 1{
		data, _ := json.Marshal(user)
		return fmt.Errorf("fail to add user to db, user: %s, ", data)
	}
	return err
}

func GetUserInfo(userId int64) (oneUser* models.User, err error){
	oneUser = new(models.User)
	has, err := engine.Where("id=?", userId).Get(oneUser)
	if err != nil {
		err = fmt.Errorf("query db failed, id: %d", userId)
		return
	}
	if !has {
		err = fmt.Errorf("user not found, id: %d", userId)
		return
	}
	return oneUser, err

}