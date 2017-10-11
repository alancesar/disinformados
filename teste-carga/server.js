const http = require('http');

let count = 1;

http.createServer((req, res) => {
    console.log(`Request ${count} em localhost:3000`);

    const value = Math.random() * 10000;
    const response = `http://localhot:3000 (${count++}) -> ${value}`;

    res.setHeader('Content-Type', 'application/json');
    res.write(response);
    res.end();
}).listen(3000);
