// User model in db

package dbmodels

import (
	"demo_p2p_bak/models"
	"fmt"
	"time"
	"github.com/go-xorm/xorm"
	"log"
)

// User user
// swagger:model User

type Transaction struct {
	// id
	ID int64 `json:"id,omitempty"`

	// from_user_id
	FromUserId int64 `json:"from_user_id"`

	// to_user_id
	ToUserId int64 `json:"to_user_id"`

	// amount
	Amount int64 `json:"amount"`

	// created_date
	CreatedDate time.Time `json:"created_date"`

	// status
	Status string `json:"status"`

	IsRepay bool
}

func getTransaction(transId int64){

}

func GetUserDebt(lenderId int64, borrowerId int64) (lend int64, borrow int64,  err error){
	lender, _ := GetUserInfo(lenderId)
	borrower, _ := GetUserInfo(borrowerId)
	if lender ==nil || borrower==nil{
		err = fmt.Errorf("user not found, %d %d", lenderId, borrowerId)
		return
	}
	lend = getDebt(engine, lenderId, borrowerId)
	borrow = getDebt(engine, borrowerId, lenderId)
	return
}


func AddTransaction(fromUserId int64, toUserId int64, amount int64, trans_type string) (trans Transaction, err error) {

	fromUser, _ := GetUserInfo(fromUserId)
	toUser, _ := GetUserInfo(toUserId)
	if  fromUser == nil {
		return trans, fmt.Errorf("user not found, id: %d", fromUserId)
	}
	if toUser == nil{
		return trans, fmt.Errorf("user not found, id: %d", toUserId)
	}
	// 先判余额是否大于转出的金额，后续事务中还会判断
	if *(fromUser.Balance) < amount {
		return trans, fmt.Errorf("user %d has balance of %d, less than %d", fromUserId,
			*(fromUser.Balance), amount)
	}

	// 开始事务
	session := engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return trans, fmt.Errorf("error to begin transaction %s", err.Error())
	}
	tmpUser := new(models.User)

	// 保证加锁顺序，转入的余额增加
	if toUserId < fromUserId {
		affected, _ := session.Where("id=?", toUserId).Incr("balance", amount).Update(tmpUser)

		if affected < 1 {
			session.Rollback()
			return trans, fmt.Errorf("update user balance failed: %d", fromUserId)
		}
	}
	if trans_type == "repay"{
		debt := getDebtForUpdate(session, toUserId, fromUserId)
		if amount > debt{
			session.Rollback()
			return trans, fmt.Errorf("%d owes %d to %d, less than provided repay number %d ", toUserId,
				debt, fromUserId, amount)
		}
	}

	// 转出的人余额减少
	affected, err := session.Where("id=?", fromUserId).And("balance >= ?", amount).Decr("balance", amount).Update(tmpUser)

	if affected < 1 {
		session.Rollback()
		return trans, fmt.Errorf("user %d has balance of %d, less than %d", fromUserId,
			fromUser.Balance, amount)
	}

	// 保证加锁顺序，转入的余额增加
	if toUserId > fromUserId {
		affected, _ := session.Where("id=?", toUserId).Incr("balance", amount).Update(tmpUser)
		if affected < 1 {
			session.Rollback()
			return trans, fmt.Errorf("update user balance failed: %d", fromUserId)
		}
	}

	// 记录交易信息
	isRepay := false
	if trans_type == "repay"{
		isRepay = true
	}
	trans = Transaction{FromUserId: fromUserId, ToUserId: toUserId, Amount: amount,
		Status: "success", CreatedDate: time.Now(), IsRepay: isRepay}
	_, err = session.Insert(&trans)
	if err != nil {
		log.Print(err.Error())
		session.Rollback()
		return trans, fmt.Errorf("add transaction failed: %d -> %d", fromUserId, toUserId)
	}
	session.Commit()

	return trans, err

}

func getDebt(session *xorm.Engine, lenderId int64, borrowerId int64) int64{
	trans := new(Transaction)
	total_lend, _ := session.Where("from_user_id = ?", lenderId).And("to_user_id=?",  borrowerId).
		And("is_repay=?", 0).SumInt(trans, "amount")
	total_repay, _ := session.Where("from_user_id = ?", borrowerId).And("to_user_id=?",  lenderId).
		And("is_repay=?", 1).SumInt(trans, "amount")

	return total_lend - total_repay
}

func getDebtForUpdate(session *xorm.Session, lenderId int64, borrowerId int64) int64{
	trans := new(Transaction)
	total_lend, _ := session.Where("from_user_id = ?", lenderId).And("to_user_id=?",  borrowerId).
		And("is_repay=?", 0).ForUpdate().SumInt(trans, "amount")
	total_repay, _ := session.Where("from_user_id = ?", borrowerId).And("to_user_id=?",  lenderId).
		And("is_repay=?", 1).ForUpdate().SumInt(trans, "amount")

	return total_lend - total_repay
}
