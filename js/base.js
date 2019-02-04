$(document).ready(function() {
    console.log("holaaaa")

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
