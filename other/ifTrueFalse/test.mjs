import {If, True, False} from "./main";
import {equal} from "assert";

equal(If(True)("fo")("bar"), "fo");
equal(If(False)("foo")("ba"), "ba");
equal(True()(), undefined);
equal(False()(), undefined);
