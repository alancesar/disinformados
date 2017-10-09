const http = require('http');

let a = 0;

http.createServer((req, res) => {
    console.log(`Request ${a} in localhost:3000`)
    
    a++;
    let value = Math.random() * (10 - 1) + 1;

    for (let i = 0; i < 100; i++) {
        value = value + Math.sqrt(Math.sqrt(value + (a * Math.pow(i, 2))));
    }

    res.writeHead(200);
    res.write(`http://localhot:3000 (${a}) -> ${value}`);
    res.end();
}).listen(3000);