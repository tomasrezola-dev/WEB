const title = document.getElementById("title");
title.textContent = "Ejercicio 1 - TP4";

//el querySelector solamente funciona para cosas que tienen class
const text = document.querySelector(".text");
text.style.border = "1px solid black";

const miBoton = document.getElementById("changeText");

miBoton.addEventListener("click", function () {
    alert("¡Botón presionado!");
    text.textContent = "Hola mundo";
    text.style.display = "block";

    const content = document.getElementById("dynamicContent");
    content.innerHTML =
    "<h2>Contenido dinamico</h2>" +
    "<ul><li>Elemento 1</li><li>Elemento 2</li></ul>"
});