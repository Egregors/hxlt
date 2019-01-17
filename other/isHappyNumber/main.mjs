export const isHappyNumber = n => {
    n = String(n);
    for (let i = 0; i < 10; i++) {
        let sum = 0;
        for (let j = 0; j < n.length; j++) {
            sum += Number(n[j]) * Number(n[j]);
        }
        if (sum === 1) {
            return true;
        }
        n = String(sum);
    }
    return false;
};
