let name = 'grady';
let age = 23;

let text = `name : ${name}, age : ${age}`;

console.log(text);

test = (strings, username, age) => {
  let string1 = strings[0];
  let string2 = strings[1];

  console.log(string1 + username + string2 + age);
}

let output = test`my name is ${name} and my age is ${age}`;