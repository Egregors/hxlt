export const roundBracketsValidator = (s) => {
    if (s.length === 0) { return true; }
    if (s.length % 2 !== 0) { return false; }

    let l = 0, r = 0;

    for (let i = 0; i < s.length; i ++) {
        if (s[i] === "(") { l++; } else { r++; }
        if (r > l) { return false; }
    }

    if (r !== l) { return false; }
    return true;
};

export default roundBracketsValidator;
