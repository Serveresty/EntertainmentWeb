function deposit() {
    var money = document.getElementById("deposit").value;

    let data = {
        Money: money,
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
        });
    }).catch((error) => {
        console.log(error)
    });
}