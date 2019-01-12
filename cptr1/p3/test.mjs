import isPerfect from "./main.mjs";
import {equal} from "assert";

equal(isPerfect(0), false);
equal(isPerfect(-23), false);
equal(isPerfect(1), false);
equal(isPerfect(99), false);
equal(isPerfect(123), false);
equal(isPerfect(23), false);
equal(isPerfect(11111), false);
equal(isPerfect(6), true);
equal(isPerfect(28), true);
equal(isPerfect(8128), true);
