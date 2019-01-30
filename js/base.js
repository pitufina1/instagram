$(document).ready(function() {
    
   /* ActualizarUsuario();*/
    $("#txtEmail").keyup(function(event) {
        if (event.keyCode === 13) {
            $("btnRegistrarse").click();
        }
    });
    $("#btnRegistrarse").click(function() {
        var nombre = $("#txtTexto").val()
        var correo = $("#txtEmail").val()
        var contrasena = $("#txtPassword").val()
        
        console.log(nombre, correo, contrasena);

        var envio = {
            nombre: nombre,
	        correo: correo,
	        contrasena: contrasena
        };

        $.post({
            url:"/insertarusuario",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#txtTexto").val('')
                $("#txtEmail").val('')
                $("#txtPassword").val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            //ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
});

/*function ActualizarUsuario() {
    var filtro = {
        correo: moment().format('YYYY-MM-DDT00:00:00Z')
    };
    $.ajax({
        url: "/listausuarios",
        method: "POST",
        data: JSON.stringify(filtro),
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            if(data != null)
                console.log(data.length + " objetos obtenidos");
            Historial_UI(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Usuario_UI(array) {
    var tbody = $("#usuario tbody");
    tbody.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            tbody.append(
                "<tr><td>" + array[x].ID + 
                "</td><td>" + array[x].Nombre + 
                "</td><td>" + array[x].Correo + 
                "</td><td>" + array[x].Contrasena + 
                "</td></tr>");
        }
    } else {
        tbody.append('<tr><td colspan="3">No hay registros de hoy</td></tr>');
        
    }
}*/