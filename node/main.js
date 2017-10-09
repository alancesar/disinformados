const http = require('http');
const urls = ['http://localhost:3000'];
const requests = 1000;

const ping = (url) => {
    return new Promise((resolve) => {
        console.log(`Fazendo requisição em ${url}`);
        
        let data = '';

        http.get(url, (response) => {
            response.on('data', chunk => data += chunk);
            response.on('end', () => resolve(data));
        });
    });
};

const main = () => {
    let init = new Date().getTime();
    let promises = [];

    for (let i = 0; i < requests; i++) {
        for (j in urls) {
            promises.push(ping(urls[j]));
        }
    }

    Promise.all(promises).then(responses => {
        responses.forEach(response => console.log(response));
        let end = new Date().getTime();
        console.log(`Tempo total -> ${end - init} milissegundos`)
    });
};

main();