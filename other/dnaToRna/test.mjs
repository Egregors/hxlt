import dnaToRna from "./main";
import {equal} from "assert";

equal(dnaToRna("ACGTGGTCTTAA"), "UGCACCAGAAUU");
equal(dnaToRna("CCGTA"), "GGCAU");
equal(dnaToRna(""), "");
equal(dnaToRna("ACNTG"), null);
