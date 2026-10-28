var global = 'usuarios';

function cargarUsuario(entity, body) {

    const fila = document.createElement("tr");

    const celdaUsuario = document.createElement("td");
    celdaUsuario.textContent = entity.nombre;

    const celdaMail = document.createElement("td");
    celdaMail.textContent = entity.mail;

    const deleteButton = document.createElement("td");
    deleteButton.innerHTML = `<button class="delete-usuario" data-id="${entity.id_usuario}">Eliminar</button>`;

    fila.appendChild(celdaUsuario);
    fila.appendChild(celdaMail);
    fila.appendChild(deleteButton)

    body.appendChild(fila);
}

function cargarMensaje(entity, body) {

    const fila = document.createElement("tr");

    const celdaUsuario = document.createElement("td");
    celdaUsuario.textContent = entity.id_usuario.Int32;

    const celdaRespMensaje = document.createElement("td");
    celdaRespMensaje.textContent = entity.id_respuesta_a_mensaje.Int32

    const celdaTexto = document.createElement("td");
    celdaTexto.textContent = entity.texto;

    const celdaMeGusta = document.createElement("td");
    celdaMeGusta.textContent = entity.me_gusta.Int32;

    const deleteButton = document.createElement("td");
    deleteButton.innerHTML = `<button class="delete-usuario" data-id="${entity.id_mensaje}">Eliminar</button>`;

    fila.appendChild(celdaUsuario);
    fila.appendChild(celdaRespMensaje);
    fila.appendChild(celdaTexto);
    fila.appendChild(celdaMeGusta);
    fila.appendChild(deleteButton)

    body.appendChild(fila);
}

const contenedor = document.getElementById("data-container");
async function cargarDatos(entity) {
    try {
        const respuesta = await fetch(
            `http://192.168.0.160:8080/api/${entity}`, {method: "GET"}
        );

        if (!respuesta.ok) {
            throw new Error(`Error HTTP: ${respuesta.status}`);
        }

        const data = await respuesta.json();

        const cuerpo = document.getElementById("usuarios-body");
        cuerpo.replaceChildren();
        if (data != null) {
            if (entity == 'usuarios') {
                data.forEach(usuario => {cargarUsuario(usuario, cuerpo);});
            }
            else {
                data.forEach(mensaje => {cargarMensaje(mensaje, cuerpo);});
            }
        }

    } catch (error) {
        console.error("Error al consultar la API:", error);
        alert('Ocurrio un error al cargar las entidades');
    }
}
cargarDatos('usuarios');


const formulario = document.getElementById("form");
formulario.addEventListener("submit", async(event) => {
    event.preventDefault();

    const var_1 = document.getElementById(`${global}_1`);
    const var_2 = document.getElementById(`${global}_2`);
    const var_3 = document.getElementById(`${global}_3`);
    if (global == 'usuarios') {
        param = JSON.stringify({
            nombre: var_1.value,
            mail: var_2.value,
            contrasenia: var_3.value
        })
    }
    else {
        id = Number(var_1.value);
        if (var_2.value != '') {
            id_msnj = Number(var_2.value)
            param = JSON.stringify({
                id_usuario: id,
                id_respuesta_a_mensaje: id_msnj,
                texto: var_3.value
            })
        }
        else {
            param = JSON.stringify({
                id_usuario: id,
                texto: var_3.value
            })
        }
    }

    //insertar
    try {
        const respuesta = await fetch(`http://192.168.0.160:8080/api/${global}`, {
            method: "POST",
            headers: {"Content-Type": "application/json"},
            body: param
        });
        var_1.value = "";
        var_2.value = "";
        var_3.value = "";
        user = await respuesta.json();
        if (!respuesta.ok) {
            throw new Error(`Error HTTP: ${respuesta.status}`);
        }
        
    }
    catch (error) {
        console.error("Error al consultar la API:", error);
        alert('Ocurrio un error al crear el usuario');
        return
    }

    //anadir dinamicamente a la tabla
    const cuerpo = document.getElementById("usuarios-body");
    if (global == 'mensajes') {
        cargarMensaje(user, cuerpo);
    }
    else {
        cargarUsuario(user, cuerpo);
    }
})

const cuerpo = document.getElementById("usuarios-body");
cuerpo.addEventListener("click", async(event) => {
    const boton = event.target.closest(".delete-usuario");
    const idUsuario = boton.dataset.id;
    try {
        const respuesta = await fetch(
            `http://192.168.0.160:8080/api/${global}/${idUsuario}`, {method: "DELETE"}
        );

        if (!respuesta.ok) {
            throw new Error(`Error HTTP: ${respuesta.status}`);
        }
    }
    catch {
        console.error("Error al consultar la API:", error);
    }

    boton.closest("tr").remove();
})

const toggle = document.getElementById("switch-label");
toggle.addEventListener("change", async(event) => {

    const encabezado = document.getElementById("encabezado");
    const formulario = document.getElementById("form");
    const tabla = document.getElementById("table");
    tabla.classList.remove("tabla-usuarios", "tabla-mensajes");
    formulario.classList.remove("form-usuarios", "form-mensajes");
    if (event.target.checked) {
        global = 'mensajes'
        tabla.classList.add('tabla-mensajes');
        formulario.classList.add('tabla-mensajes');

        encabezado.innerHTML = `
            <tr>
                <th>User</th>
                <th>Mnsj resp.</th>
                <th>text</th>
                <th>Likes</th>
                <th>Accion</th>
            </tr>
        `;
        
        formulario.innerHTML = `

            <label>id_usuario:</label><br>
            <input type="text" id="mensajes_1"><br><br>
            <label>id_mensaje_a_responder:</label><br>
            <input type="text" id="mensajes_2"><br><br>
            <label>Texto:</label><br>
            <input type="text" id="mensajes_3"><br><br>
            <button type="submit">Enviar</button>
        `;
        
        cargarDatos('mensajes');

    }
    else {
        global = 'usuarios'
        tabla.classList.add('tabla-usuarios');
        formulario.classList.add('tabla-usuarios');

        encabezado.innerHTML = `
            <tr>
                <th>User</th>
                <th>Mail</th>
                <th>Accion</th>
            </tr>
        `;
        
        formulario.innerHTML = `

            <label>Nombre:</label><br>
            <input type="text" id="usuarios_1"><br><br>
            <label>Mail:</label><br>
            <input type="text" id="usuarios_2"><br><br>
            <label>Password:</label><br>
            <input type="password" id="usuarios_3"><br><br>
            <button type="submit">Enviar</button>
        `;

        cargarDatos('usuarios');
    }
})