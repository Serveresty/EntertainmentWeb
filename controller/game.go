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
		send_data.Multiply = data.Multiply
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
