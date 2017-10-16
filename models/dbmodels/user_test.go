package dbmodels

import (
	"testing"
	"time"
	"fmt"
	"math/rand"
	"demo_p2p_bak/models"
)

func TestAddUser(t *testing.T) {
	rand.Seed(int64(time.Now().UnixNano()))
	var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	name := make([]byte, 5)
	for i := range name {
		name[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	for i:=0; i<3; i++{
		var balance int64 = 200
		s := string(name)
		name := fmt.Sprintf("%s_%d", s, i)
		u := models.User{Username: &name , Balance: &balance}
		AddUser(&u)
		if u.Id == 0{
			t.Errorf("no id after insert: %s", name)
		}
		dbUser, _ := GetUserInfo(u.Id)
		if *(dbUser.Username) != name{
			t.Errorf("name not matching to db: %s vs %s", name, dbUser.Username)
		}

	}

}