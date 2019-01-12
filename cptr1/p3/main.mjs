const isPerfect = n => {
  if (n <= 0) {
    return false;
  }

  let sum = 0;
  for (let i = n - 1; i > 0; i--) {
    if (n % i === 0) {
      sum += i;
    }
  }
  return sum === n;
};

export default isPerfect;