package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DiceData struct {
	BetAmount string
	Range     string
	Multiply  string
	WinChance string
	Profit    string
	Number    string
}

type SendDiceData struct {
	BetAmount string
	Range     string
	Multiply  string
	WinChance string
	Profit    string
	Number    string
}

func GetDiceData(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var data DiceData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var send_data SendDiceData
	send_data.BetAmount = "200"
	send_data.Multiply = "2.5"
	send_data.Profit = "300"
	send_data.Range = "34"
	send_data.WinChance = "34"
	send_data.Number = "5"
	rw.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(send_data)
	return
}
