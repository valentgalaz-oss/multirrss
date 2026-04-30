import { InstagraNode } from "instagranode";
import json from "json";

const insta = new InstagraNode();
const profile = await insta.searchProfile("kyliejenner");

json.stringify(profile, null, 2);
