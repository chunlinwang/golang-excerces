const t1 = new Date();
let n = 0;
for (var i = 0; i <1000000000; ++i) {
  n += i;
}
const t2 = new Date();

console.log("t1 : ",t2.valueOf())
console.log("t1 : ",t1.valueOf())
console.log("time : ",t2.valueOf()-t1.valueOf())
