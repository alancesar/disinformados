const rp = require('request-promise')

const urls = ['http://localhost:3000', 'http://localhost:3001'];
const requests = 100;

const ping = (url) => {
    console.log(`Fazendo requisição em ${url}`);
    return rp(url).then(response => response);
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