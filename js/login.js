$(document).ready(function() {
    console.log("holaaaa");

    $("#btnLogin").click(function() {
         var correo = $("#txtEmail").val()
         var contrasena = $("#txtPassword").val()
         
         console.log(correo, contrasena);
 
         var envio = {
             correo: correo,
             contrasena: contrasena
         };
 
         $.post({
             url:"/loginusuario",
             data: JSON.stringify(envio),
             method: "POST",
             success: function(data, status, jqXHR) {
                 console.log(data);
                 
             },
             dataType: "json"
 
         }).done(function(data) {
             console.log("Petición realizada");
             if (data == true){
                 window.location.href="/main";
             }
             //ActualizarHistorial();
         
         }).fail(function(data) {
             console.log("Petición fallida");
         
         }).always(function(data){
             console.log("Petición completa");
         });
     });
 });