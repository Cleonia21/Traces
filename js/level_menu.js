

let focus_inpt = document.getElementById("input_1");


let inputs = document.getElementsByClassName('series_input');
for(let i = inputs.length; i--;) {
    inputs[i].addEventListener('focus', function () {
        focus_inpt = this;
    });
}

let buttons = document.getElementsByClassName("level_menu_button_area");
for(let i = buttons.length; i--;) {
    buttons[i].onmousedown = function () { return false; }
}

function change_focus(direction) {
    let new_inpt = document.getElementById("input_" + (parseInt(focus_inpt.id[6]) + direction).toString());
    if (new_inpt) {
        new_inpt.focus();
    }
    return true;
}
document.getElementById("arrow_up").onclick = function () {
    change_focus(-1);
}
document.getElementById("arrow_down").onclick = function () {
    change_focus(1);
}

document.getElementById("erase").onclick = function () {
    focus_inpt.value = focus_inpt.value.substring(0, focus_inpt.value.length - 1);
}

document.getElementById("clear").onclick = function () {
    focus_inpt.value = "";
}

document.getElementById("fill").onclick = function () {
    focus_inpt.value = "+";
    change_focus(1);
}

document.getElementById("tab").onclick = function () {
    focus_inpt.focus();
}

