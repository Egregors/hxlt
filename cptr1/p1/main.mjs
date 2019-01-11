export const isPowerOfThree = num => {
  const base = 3;
  if (Math.pow(base, Math.round(Math.log(num) / Math.log(base))) === num) {
    return true;
  }

  return false;
};

export default isPowerOfThree;