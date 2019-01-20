export default m => {
    return f(Math.floor(m / 60)) + ":" + f(m % 60);    
};

const f = (n) => {
    return String(n).length === 2 ? String(n) : String("0" + n);
};