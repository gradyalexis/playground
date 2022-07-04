let s = {
  members: ['a', 'b', 'c'],
  getMembers() {
    this.members.map((name) => {
      return this;
    })
  },
}

console.log(s);