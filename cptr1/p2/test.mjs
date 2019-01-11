import * as assert from "assert";
import diff from "./main.mjs";

assert.equal(diff(0, 45), 45);
assert.equal(diff(0, 180), 180);
assert.equal(diff(0, 190), 170);
assert.equal(diff(120, 280), 160);