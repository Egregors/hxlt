export const True = a => b => a;
export const False = a => b => b;
export const If = f => a => b => f(a)(b);
