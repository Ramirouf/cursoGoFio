async function getData() {
    const response = await fetch('http://127.0.0.1:5500/JsonMRR/datos.json');
    const json = await response.json();

    console.log(json);
    console.log(json.nombre);
    console.log(json.experiencia);
    json.experiencia.forEach(element => {
        console.log(element.empresa);
    });
}

getData();