var cards_block = document.querySelector('.cards');
var cards = document.querySelectorAll('.cards > div');

function start() {
  var betamount = document.getElementById("BetAmount");
  var range = document.getElementById("Range");
  var multiply = document.getElementById("Multiply");
  var winchance = document.getElementById("WinChance");
  var profit = document.getElementById("Profit");

  var random;

  let data = {
    BetAmount: betamount.value,
    Range: range.value,
    Multiply: multiply.value,
    WinChance: winchance.value,
    Profit: profit.value,
    Number: null,
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

          betamount.value = result['BetAmount'];
          range.value = result['Range'];
          multiply.value = result['Multiply'];
          winchance.value = result['WinChance'];
          profit.value = result['Profit'];
          random = result['Number']

          
          var point = parseInt(random, 10)
          cards_block.style.left = -point * 100 + 'px';
          setTimeout(function() {

            cards[point].style.background = '#7B90F7';
            cards[point].style.color = 'white';
          }, 5000)

      });
  }).catch((error) => {
      console.log(error)
  });
}

