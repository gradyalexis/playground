signIn = (username, password, age) => {
  console.log(username + ' has password ' + password + ' and age ' + age);
}

let data = ['grady', '123', 23];

signIn(...data);

member = (...members) => {
  console.log(members);
}

username = 'grady';
password = '312';
age = 23;

member(username, password, age);
