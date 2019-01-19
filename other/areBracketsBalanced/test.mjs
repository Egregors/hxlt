import roundBracketsValidator from "./main.mjs";
import * as assert from "assert";

assert.equal(roundBracketsValidator(""), true); 
assert.equal(roundBracketsValidator("()()"), true);
assert.equal(roundBracketsValidator("((()))"), true);
assert.equal(roundBracketsValidator("((())"), false);
assert.equal(roundBracketsValidator(")("), false);
assert.equal(roundBracketsValidator("(()()))("), false);
