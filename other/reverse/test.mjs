import reverse from "./main";
import {equal} from "assert";

equal(reverse("lol"), "lol");
equal(reverse("main"), "niam");
equal(reverse("qwe"), "ewq");
equal(reverse("qqqrrr"), "rrrqqq");
equal(reverse("rqrqr"), "rqrqr");
equal(reverse(""), "");
equal(reverse("qwertyuiop"), "poiuytrewq");
