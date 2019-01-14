import {equal} from "assert";
import {substr} from "./main";

equal(substr("abba", 0, 1), "a"); // => a
equal(substr("abda", 1, 2), "bd"); // => bd
equal(substr("abba", -10, 2), "ab"); // => ab
equal(substr("abba", -1, 100), "abba"); // => abba
equal(substr("abba", -1, -1), "a"); // => a
equal(substr("abda", 1, -10), "b"); // => b
equal(substr("abba", 1, 10), "bba"); // => bba
equal(substr("abba", 1, 0), ""); // => ''
equal(substr("abba", 100, 3), ""); // => ''
equal(substr(""), ""); // => ''
equal(substr("abba"), "abba"); // => abba
