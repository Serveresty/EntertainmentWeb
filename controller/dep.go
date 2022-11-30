package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type HowMuchMoney struct {
	Money   string
	Balance string
}

type Result struct {
	Money   string
	Balance string
}

func (s *DataBase) Deposit(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conditionsMap := map[string]any{}
	var data HowMuchMoney
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var send_data Result
	send_data.Money = data.Money
	user, status := s.GetUser(rw, r)
	_ = status
	conditionsMap["balance"] = user.balance
	uID := user.id
	perem := `
			INSERT INTO transactions (user_id, type, stat, summ) VALUES(?, ?, ?, ?)
		`

	//Добавляю в БД запись о регистрации, если нет ошибок
	insert, errdb := s.Data.Query(perem, uID, "Deposit", "Accept", send_data.Money)
	defer func() {
		if insert != nil {
		}
	}()
	if errdb != nil {
	}

	bal := conditionsMap["balance"].(string)
	n, _ := strconv.ParseFloat(bal, 64)
	dep, _ := strconv.ParseFloat(send_data.Money, 64)
	conditionsMap["balance"] = fmt.Sprint(n + dep)

	perem2 := `
			UPDATE users_account SET balance = ? WHERE id = ?;
		`

	insert2, errdb2 := s.Data.Query(perem2, conditionsMap["balance"], uID)
	defer func() {
		if insert2 != nil {
		}
	}()
	if errdb2 != nil {
	}

	send_data.Balance = conditionsMap["balance"].(string)
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(send_data)
	return
}
