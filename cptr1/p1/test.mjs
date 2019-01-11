import isPowerOfThree from "./main.mjs";
import * as assert from "assert";

assert.equal(isPowerOfThree(1), true);
assert.equal(isPowerOfThree(2), false);
assert.equal(isPowerOfThree(9), true);
