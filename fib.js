const rl = require('readline');

let interface = rl.createInterface({ input: process.stdin, output: process.stdout});

interface.question('What is the value of n?', (a)=>{
    let n = Number(a);
    if(isNaN(n))
        throw new Error('Please enter a number');
    const start = Date.now();
    let result = fib(n);
    const end = Date.now();
    console.log(`the value in the ${n}th position is ${result}, this calculation took ${end - start} milliseconds`)
});


function fib(n){
    if(n === 0) return 0;
    if(n < 2){
        return 1
    }
    return fib(n-1) + fib(n-2);
}