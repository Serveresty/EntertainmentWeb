package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type DiceData struct {
	BetAmount        string
	Range            string
	Multiply         string
	WinChance        string
	Profit           string
	Number1          int
	Number2          int
	Number3          int
	Number4          int
	NotEnoughMoney   bool
	OutOfRange       bool
	OutOfMultiply    bool
	UnknownWinChance bool
}

type SendDiceData struct {
	BetAmount        string
	Range            string
	Multiply         string
	WinChance        string
	Profit           string
	Number1          int
	Number2          int
	Number3          int
	Number4          int
	NotEnoughMoney   bool
	OutOfRange       bool
	OutOfMultiply    bool
	UnknownWinChance bool
}

func (s *DataBase) GetDiceData(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var data DiceData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var send_data SendDiceData

	bal := conditionsMap["balance"].(string)
	n, _ := strconv.ParseFloat(bal, 64)

	//Проверка наличия суммы ставки на балансе
	if k, _ := strconv.ParseFloat(data.BetAmount, 64); n >= k {
		send_data.BetAmount = data.BetAmount
	} else {
		send_data.NotEnoughMoney = true
	}

	//Проверка диапазона умножения
	if k, _ := strconv.ParseFloat(data.Multiply, 64); k <= 9500 && k >= 1.0106 {
		rng, _ := strconv.ParseFloat(data.Range, 64)
		wch, _ := strconv.ParseFloat(data.WinChance, 64)
		if k == 95/rng && k == 95/wch {
			send_data.Multiply = data.Multiply
		} else {
			send_data.OutOfMultiply = true
		}
	} else {
		send_data.OutOfMultiply = true
	}

	//Проверка диапазона ставки
	if k, _ := strconv.ParseFloat(data.Range, 64); k <= 94.0 && k >= 0.01 {
		send_data.Range = data.Range
	} else {
		send_data.OutOfRange = true
	}

	//Проверка шанса на победу
	if k, _ := strconv.ParseFloat(data.WinChance, 64); k <= 94.0 && k >= 0.01 {
		send_data.WinChance = data.WinChance
	} else {
		send_data.UnknownWinChance = true
	}

	//Проверка на наличие ошибок при ставке
	if send_data.NotEnoughMoney == false && send_data.OutOfRange == false && send_data.OutOfMultiply == false && send_data.UnknownWinChance == false {
		bet_am, _ := strconv.ParseFloat(send_data.BetAmount, 64)
		mult, _ := strconv.ParseFloat(data.Multiply, 64)

		send_data.Profit = fmt.Sprintf("%f", bet_am*mult-bet_am)
		send_data.Number1 = rand.Intn(9 + 1)
		send_data.Number2 = rand.Intn(9 + 1)
		send_data.Number3 = rand.Intn(9 + 1)
		send_data.Number4 = rand.Intn(9 + 1)
		num := fmt.Sprintf("%d%d.%d%d", send_data.Number1, send_data.Number2, send_data.Number3, send_data.Number4)
		s.transactionDice(num, send_data.Range, send_data.Profit, send_data.BetAmount, n)
	} else {
		send_data.Number1 = 0
		send_data.Number2 = 0
		send_data.Number3 = 0
		send_data.Number4 = 0
	}
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(send_data)
	return
}

func (s *DataBase) transactionDice(n string, rng string, pt string, bt string, bal float64) map[string]any {
	numb, _ := strconv.ParseFloat(n, 64)
	rang, _ := strconv.ParseFloat(rng, 64)
	prof, _ := strconv.ParseFloat(pt, 64)
	bet, _ := strconv.ParseFloat(bt, 64)
	var status string
	var summ string

	if rang > numb {
		bal += prof
		status = "win"
		summ = "+" + fmt.Sprint(prof)
	} else {
		bal -= bet
		status = "lost"
		summ = "-" + fmt.Sprint(bet)
	}

	perem := `
			INSERT INTO transactions (user_id, type, stat, summ) VALUES(?, ?, ?, ?)
		`

	//Добавляю в БД запись о регистрации, если нет ошибок
	insert, errdb := s.Data.Query(perem, conditionsMap["username"], "DiceBet", status, summ)
	defer func() {
		if insert != nil {
		}
	}()
	if errdb != nil {
	}

	conditionsMap["balance"] = fmt.Sprint(bal)

	perem2 := `
			UPDATE users_account SET balance = ? WHERE username = ?;
		`

	insert2, errdb2 := s.Data.Query(perem2, conditionsMap["balance"], conditionsMap["username"])
	defer func() {
		if insert2 != nil {
		}
	}()
	if errdb2 != nil {
	}

	fmt.Println(conditionsMap["balance"])
	return conditionsMap
}
