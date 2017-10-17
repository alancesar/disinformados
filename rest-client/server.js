const http = require('http');

http.createServer((req, res) => {
  const response = {
    cep: '13086-902',
    logradouro: 'Rua Doutor Ricardo Benetton Martins',
    complemento: 's/n',
    bairro: 'Polo II de Alta Tecnologia (Campinas)',
    localidade: 'Campinas',
    uf: 'SP',
    unidade: '',
    ibge: '3509502',
    gia: '244'
  };

  res.setHeader('Content-Type', 'application/json');
  res.write(JSON.stringify(response, null, 2));
  res.end();
}).listen(3000);
