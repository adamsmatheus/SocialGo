$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Passou aqui !")
}