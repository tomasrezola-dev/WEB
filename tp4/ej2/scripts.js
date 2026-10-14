console.log("Script cargado");
const title = document.getElementById("title");
title.textContent = "Ejercicio 2 - TP4";

const addButton = document.getElementById("add");
addButton.addEventListener("click", function () {
    const input = document.getElementById("text-item");
    const list = document.getElementById("list");
    const item = document.createElement("li");
    item.textContent = input.value;
    list.appendChild(item);
    input.value = "";  //cuando es input de texto, se trata como value el campo
});

const delButton = document.getElementById("delete");
delButton.addEventListener("click", function () {
    const list = document.getElementById("list");
    if (list.lastChild) {//devuelve el ultimo hiji del elemento, si no tiene hijos devuelve null
        list.removeChild(list.lastChild);
    }
});