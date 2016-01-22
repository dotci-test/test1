const style = require('ansi-styles');

console.log(style.green.open + 'Hello world!' + style.green.close);
console.log("\033[31mRed\033[0m");