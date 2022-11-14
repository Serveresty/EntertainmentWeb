
function RangeEdit(num) {

    var wrchance = document.getElementById("WrondRange")

    if (num < 0.01 || num == "") {
        wrchance.innerHTML = 'Min Range 0.01'; 
        return;
    }
    if (num > 0.01 && num != ""){
        wrchance.innerHTML = "";
        document.getElementById("WinChance").setAttribute('value', num/100);
    }

    if (num > 9400) {
        wrchance.innerHTML = 'Max Range 9400';
        return;
    }
    if (num < 9400) {
        wrchance.innerHTML = "";
        document.getElementById("WinChance").setAttribute('value', num/100);
    }
}
/*
function MultiplyEdit(num) {
    document.getElementById("Range").setAttribute('value', num);
    document.getElementById("WinChance").setAttribute('value', num);

    var wrchance = document.getElementById("WrondMultiply")

    if (num < 1.01 || num == "") {
        wrchance.innerHTML = 'Min Multiply 1.01'; 
    }else{
        wrchance.innerHTML = "";
    }

    if (num > 9500) {
        wrchance.innerHTML = 'Max Multiply 9500'; 
    }
}

function WinChanceEdit(num) {
    document.getElementById("Range").setAttribute('value', num * 100);
    document.getElementById("Multiply").setAttribute('value', num);

    var wrchance = document.getElementById("WrondChance")

    if (num < 0.01 || num == "") {
        wrchance.innerHTML = 'Min Chance 0.01%'; 
    }else{
        wrchance.innerHTML = "";
    }

    if (num > 94) {
        wrchance.innerHTML = 'Max Chance 94%'; 
    }
}

function BetEdit(num) {
    var wrchance = document.getElementById("WrondAmount")

    if (num == "") {
        wrchance.innerHTML = 'This field is required'; 
    }else{
        wrchance.innerHTML = "";
    }


}*/