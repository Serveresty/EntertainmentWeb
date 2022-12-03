function deposit() {
    var money = document.getElementById("deposit").value;
    var bal =  document.getElementById("balance");

    let data = {
        Money: money,
        Balance: bal.textContent,
    };
    fetch("/profile", {
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
            document.getElementById("resDep").textContent = "Deposit success, +" + result['Money'] + "$";
            bal.textContent = result['Balance'] + " $";
        });
    }).catch((error) => {
        console.log(error)
    });
}

function eddt() {
    document.getElementById("resDep").textContent = "";
}