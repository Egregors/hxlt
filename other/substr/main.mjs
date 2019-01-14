export const substr = (s, a = 0, l = s.length) => {
  if (l < 0) {
    l = 1;
  }

  if (a < 0) {
    a = 0;
  }

  if (a >= s.length) {
    return "";
  }

  if (l === 0) {
    return "";
  }

  let end_idx = l + a - 1;
  let res = "";

  if (end_idx > s.length - 1) {
    for (let i = a; i < s.length; i++) {
      res += s[i];
    }
    return res;
  }

  for (let i = a; i < a + l; i++) {
    res += s[i];
  }

  return res;
};
