let name = 'Grady';
let age = 23;
getData = () => {
  return `my name ${name} and my age ${age}`;
}

let member = {
  name, age, getData
}

console.log(member);

let account = {
  a: 'alex',
  b: 20,
}

let { a, b : c } = account;

console.log(a);
console.log(c);
