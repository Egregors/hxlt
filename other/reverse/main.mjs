export const r = (s) => {
    return s.length <= 1 ? s : r(s.substr(1, s.length-1)) + s[0];
};

export default r;