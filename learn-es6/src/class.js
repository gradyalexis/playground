class Members {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  getData() {
    return `my name is ${this.name} and my age is ${this.age}`;
  }
}

let member = new Members('Grady', 23);
console.log(member);
