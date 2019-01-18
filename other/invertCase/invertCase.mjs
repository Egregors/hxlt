export const invertCase = (s) => {
    let res = "";
    for (let i = 0; i < s.length; i++) {
        if (s[i] === s[i].toUpperCase()) {
            res += s[i].toLowerCase();
        } else {
            res += s[i].toUpperCase();
        }
    }
    return res;
};

export default invertCase;