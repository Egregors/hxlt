export const diff = (a, b) => {
    if (a == b) {
        return 0;
    }

    if (a != 0) {
        if (b < a) {
            let c; 
            c = a; 
            a = b; 
            b = c;
        }
        b = b - a;
        a = 0
    }

    if (b <= 180) {
        return b;
    }

    return 360 - b;
};


export default diff;