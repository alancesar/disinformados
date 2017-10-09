const http = require('http');

let a = 0, b = 0;

const response = (res, msg) => {
    res.writeHead(200);
    res.write(msg);
    res.end();
};

http.createServer((req, res) => {
    a++;
    console.log(`Request ${a} in localhost:3000`)
    response(res, `http://localhot:3000 -> ${a}`);
}).listen(3000);

http.createServer((req, res) => {
    b++;
    console.log(`Request ${b} in localhost:3001`)
    response(res, `http://localhot:3001 -> ${b}`);
}).listen(3001);