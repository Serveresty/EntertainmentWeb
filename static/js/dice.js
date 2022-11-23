function start() {
  var betamount = document.getElementById("BetAmount");
  var range = document.getElementById("Range");
  var multiply = document.getElementById("Multiply");
  var winchance = document.getElementById("WinChance");
  var profit = document.getElementById("Profit");

  let data = {
    BetAmount: betamount.value,
    Range: range.value,
    Multiply: multiply.value,
    WinChance: winchance.value,
    Profit: profit.value,
    Number1: 0,
    Number2: 0,
    Number3: 0,
    Number4: 0,
    NotEnoughMoney: false,
    OutOfRange: false,
    OutOfMultiply: false,
    UnknownWinChance: false,
  };

  fetch("/dice", {
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
    },
      method: "POST",
      body: JSON.stringify(data)
  }).then((response) => {
      response.text().then(function (data) {
          let result = JSON.parse(data);
          console.log(result)

          if (result['NotEnoughMoney'] == true) {
            alert("Not Enough Money");
          }

          if (result['OutOfRange'] == true) {
            alert("Wrong Range");
          }

          if (result['OutOfMultiply'] == true) {
            alert("Wrong Multiply");
          }

          if (result['UnknownWinChance'] == true) {
            alert("Wrong Win Chance");
          }


          betamount.value = result['BetAmount'];
          range.value = result['Range'];
          multiply.value = result['Multiply'];
          winchance.value = result['WinChance'];
          profit.value = result['Profit'];
          num1 = result['Number1'];
          num2 = result['Number2'];
          num3 = result['Number3'];
          num4 = result['Number4'];

          roller = document.getElementById("resultDice");

          roller.textContent = num1.toString() + num2.toString() + num3.toString() + num4.toString();
      });
  }).catch((error) => {
      console.log(error)
  });
}

