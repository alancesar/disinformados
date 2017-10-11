fs = require('fs');

// (x² * √x) / (√x² + 1)
const calc = (number) => {
  return new Promise((resolve) => {
    resolve((Math.pow(number, 2) * Math.sqrt(number)) / (Math.sqrt(Math.pow(number, 2) + 1)));
  });
};

let init = new Date().getTime();

fs.readFile('numbers.txt', 'utf8', (error, data) => {
  let numbers = data.split('\n');
  let promises = numbers.map(calc);

  Promise.all(promises).then(responses => {
    responses.forEach(response => {
      console.log(response)
    });

    let end = new Date().getTime();
    console.log(`Tempo total: ${end - init} milissegundos`);
  });
});
