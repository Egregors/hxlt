// cherch nums -> nums
const print = f => f(x => x + 1)(0);

// nums
const Zero = f => n => n;
const One = f => n => f(n);
const Two = f => n => f(f(n));

console.log(print(Zero));
console.log(print(One));
console.log(print(Two));

// Succ
const Succ = n => (s => z => s (n (s) (z)));

console.log(print(Succ(Succ(Zero))));
