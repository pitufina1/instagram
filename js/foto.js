/*$(document).ready(function() {

    console.log("hola holita")
    ActualizarFotos();
    $("#inputFile").keyup(function(event) {
        if (event.keyCode === 13) {
            $("#btnSave").click();
        }
    });
  
    $("#btnSave").click(function() {
        var texto = $("#inputFile").val();
        console.log(texto);
        
        var foto = {
            nombrefoto: nombrefoto
        };

        $.post({
            url:"/main",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#inputFile").val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            ActualizarFoto();
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
});*/



/*function ActualizarFotos() {
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
    var div = $("#fotos");//aki
    div.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            div.append(
                "div"
                    +"img src='/files/"+array[x].NombreFoto+"' width='250px' height='150px'>"+
                "</div>");
        }
    } else {
        div.append('<div colspan="3">No hay registros de hoy</div>');
        
    }
}*/