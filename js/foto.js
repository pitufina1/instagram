$(document).ready(function() {
    console.log("holaaaa")

   /* Subir Foto();*/
    $("#inputFile").keyup(function(event) {
        if (event.keyCode === 13) {
            $("#btnSave").click();
        }
    });
    $("#btnSave").click(function() {
        var nombrefoto = $("#inputFile").val()
        
        console.log("oleee");

        var envio = {
            nombrefoto: nombrefoto
        };

        $.post({
            url:"/upload",
            method: "POST",
            data: JSON.stringify(envio),
            contentType: "application/json",
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#inputFile").val('')
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