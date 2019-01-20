import formattedTime from "./main";
import {equal} from "assert";

equal(formattedTime(5), "00:05");
equal(formattedTime(15), "00:15");
equal(formattedTime(60), "01:00");
equal(formattedTime(67), "01:07");
equal(formattedTime(130), "02:10");
equal(formattedTime(175), "02:55");
equal(formattedTime(549), "09:09");
equal(formattedTime(600), "10:00");
equal(formattedTime(754), "12:34");
equal(formattedTime(1293), "21:33");
