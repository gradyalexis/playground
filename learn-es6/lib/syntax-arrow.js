"use strict";

var members = ['a', 'b', 'c'];
members.forEach(function (member) {
  console.log(member);
});
members.forEach(function (member) {
  console.log(member);
});
members.forEach(function (member) {
  return console.log(member);
});