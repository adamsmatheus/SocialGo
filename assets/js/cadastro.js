$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Passou aqui !")

    if($('#senha').val() != $('#confirmar-password').val()){
        alert("As senhas não são iguais!")
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    });
}