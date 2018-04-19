const fs = require("fs");
const solc = require('solc');

let source = fs.readFileSync('../../contracts/IRMAScheme.sol', 'utf8');

let compiledContract = solc.compile(source, 1);
// let abi = compiledContract.contracts['nameContract'].interface;
// let bytecode = compiledContract.contracts['nameContract'].bytecode;

console.log(compiledContract);
