package dbmodels

import (
	"testing"
	"demo_p2p_bak/models"
	"math/rand"
	"time"
	"fmt"
	"log"
)

var user_ids = make([]int64, 0)
var a, b int64
func initFunc() {
	rand.Seed(int64(time.Now().UnixNano()))
	var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	name := make([]byte, 5)
	for i := range name {
		name[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	for i:=0; i<2; i++{
		var balance int64 = 200
		s := string(name)
		name := fmt.Sprintf("%s_%d", s, i)
		u := models.User{Username: &name , Balance: &balance}
		AddUser(&u)
		user_ids = append(user_ids, u.Id)
	}
	a, b = user_ids[0], user_ids[1]
	log.Print("v", a, b)

}

func TestTransaction(t *testing.T) {
	initFunc()
	log.Print(a, b)
	AddTransaction(a, b, 100, "loan")
	checkBalance(100, 300, t, a, b )
	checkDebt(100, 0, t,a,b)

	AddTransaction(b, a, 50, "load")
	checkBalance(150, 250, t,a,b)
	checkDebt(100, 50, t,a,b)

	AddTransaction(b, a, 50, "repay")
	checkBalance(200, 200, t,a,b)
	checkDebt(50, 50, t,a,b)


}

func checkBalance(a_num int64, b_num int64, t *testing.T, a int64, b int64){
	userA, _ :=  GetUserInfo(a)
	userB, _:= GetUserInfo(b)
	if *(userA.Balance) != a_num || *(userB.Balance) != b_num{
		t.Errorf("expect %d %d, got %d %d", a_num, b_num, userA.Balance, userB.Balance)
	}
}

func checkDebt(exLend int64, exBorrow int64, t *testing.T, a int64, b int64){

	if lend, borrow, _ := GetUserDebt(a, b); lend!=exLend || borrow!=exBorrow{
		t.Errorf("expect debt %d %d, got %d %d", exLend, exBorrow, lend, borrow)
	}
}