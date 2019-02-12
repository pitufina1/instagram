$(document).ready(function() {
    console.log("holaaaa")

    ActualizarFotos();
    
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

function ActualizarFotos() {
    $.ajax({
        url: "/lista",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            if(data != null)
                console.log(data.length + " objetos obtenidos");
            Historial_Fotos(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Historial_Fotos(array) {
    var div = $("#fotos");
    div.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            div.append(
                "<div>"
                    +"<img src='/files/"+array[x].NombreFoto+"' width='250px' height='150px'>"+
                "</div>");
        }
    } else {
        div.append('<div colspan="3">No hay registros de hoy</div>');
        
    }
}