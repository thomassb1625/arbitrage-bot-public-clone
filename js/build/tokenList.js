"use strict";
// To parse this data:
//
//   import { Convert, TokenList } from "./file";
//
//   const tokenList = Convert.toTokenList(json);
//
// These functions will throw an error if the JSON doesn't
// match the expected interface, even if the JSON is valid.
Object.defineProperty(exports, "__esModule", { value: true });
exports.Convert = exports.Tag = void 0;
var Tag;
(function (Tag) {
    Tag["Erc20"] = "erc20";
    Tag["MetaTx"] = "metaTx";
    Tag["Plasma"] = "plasma";
    Tag["Pos"] = "pos";
    Tag["Stablecoin"] = "stablecoin";
    Tag["Swapable"] = "swapable";
})(Tag = exports.Tag || (exports.Tag = {}));
// Converts JSON strings to/from your types
// and asserts the results of JSON.parse at runtime
class Convert {
    static toTokenList(json) {
        return cast(JSON.parse(json), r("TokenList"));
    }
    static tokenListToJson(value) {
        return JSON.stringify(uncast(value, r("TokenList")), null, 2);
    }
}
exports.Convert = Convert;
function invalidValue(typ, val, key = '') {
    if (key) {
        throw Error(`Invalid value for key "${key}". Expected type ${JSON.stringify(typ)} but got ${JSON.stringify(val)}`);
    }
    throw Error(`Invalid value ${JSON.stringify(val)} for type ${JSON.stringify(typ)}`);
}
function jsonToJSProps(typ) {
    if (typ.jsonToJS === undefined) {
        const map = {};
        typ.props.forEach((p) => map[p.json] = { key: p.js, typ: p.typ });
        typ.jsonToJS = map;
    }
    return typ.jsonToJS;
}
function jsToJSONProps(typ) {
    if (typ.jsToJSON === undefined) {
        const map = {};
        typ.props.forEach((p) => map[p.js] = { key: p.json, typ: p.typ });
        typ.jsToJSON = map;
    }
    return typ.jsToJSON;
}
function transform(val, typ, getProps, key = '') {
    function transformPrimitive(typ, val) {
        if (typeof typ === typeof val)
            return val;
        return invalidValue(typ, val, key);
    }
    function transformUnion(typs, val) {
        // val must validate against one typ in typs
        const l = typs.length;
        for (let i = 0; i < l; i++) {
            const typ = typs[i];
            try {
                return transform(val, typ, getProps);
            }
            catch (_) { }
        }
        return invalidValue(typs, val);
    }
    function transformEnum(cases, val) {
        if (cases.indexOf(val) !== -1)
            return val;
        return invalidValue(cases, val);
    }
    function transformArray(typ, val) {
        // val must be an array with no invalid elements
        if (!Array.isArray(val))
            return invalidValue("array", val);
        return val.map(el => transform(el, typ, getProps));
    }
    function transformDate(val) {
        if (val === null) {
            return null;
        }
        const d = new Date(val);
        if (isNaN(d.valueOf())) {
            return invalidValue("Date", val);
        }
        return d;
    }
    function transformObject(props, additional, val) {
        if (val === null || typeof val !== "object" || Array.isArray(val)) {
            return invalidValue("object", val);
        }
        const result = {};
        Object.getOwnPropertyNames(props).forEach(key => {
            const prop = props[key];
            const v = Object.prototype.hasOwnProperty.call(val, key) ? val[key] : undefined;
            result[prop.key] = transform(v, prop.typ, getProps, prop.key);
        });
        Object.getOwnPropertyNames(val).forEach(key => {
            if (!Object.prototype.hasOwnProperty.call(props, key)) {
                result[key] = transform(val[key], additional, getProps, key);
            }
        });
        return result;
    }
    if (typ === "any")
        return val;
    if (typ === null) {
        if (val === null)
            return val;
        return invalidValue(typ, val);
    }
    if (typ === false)
        return invalidValue(typ, val);
    while (typeof typ === "object" && typ.ref !== undefined) {
        typ = typeMap[typ.ref];
    }
    if (Array.isArray(typ))
        return transformEnum(typ, val);
    if (typeof typ === "object") {
        return typ.hasOwnProperty("unionMembers") ? transformUnion(typ.unionMembers, val)
            : typ.hasOwnProperty("arrayItems") ? transformArray(typ.arrayItems, val)
                : typ.hasOwnProperty("props") ? transformObject(getProps(typ), typ.additional, val)
                    : invalidValue(typ, val);
    }
    // Numbers can be parsed by Date but shouldn't be.
    if (typ === Date && typeof val !== "number")
        return transformDate(val);
    return transformPrimitive(typ, val);
}
function cast(val, typ) {
    return transform(val, typ, jsonToJSProps);
}
function uncast(val, typ) {
    return transform(val, typ, jsToJSONProps);
}
function a(typ) {
    return { arrayItems: typ };
}
function u(...typs) {
    return { unionMembers: typs };
}
function o(props, additional) {
    return { props, additional };
}
function m(additional) {
    return { props: [], additional };
}
function r(name) {
    return { ref: name };
}
const typeMap = {
    "TokenList": o([
        { json: "name", js: "name", typ: "" },
        { json: "version", js: "version", typ: r("Version") },
        { json: "logoURI", js: "logoURI", typ: "" },
        { json: "keywords", js: "keywords", typ: a("") },
        { json: "tags", js: "tags", typ: r("Tags") },
        { json: "timestamp", js: "timestamp", typ: Date },
        { json: "tokens", js: "tokens", typ: a(r("Token")) },
    ], false),
    "Tags": o([
        { json: "stablecoin", js: "stablecoin", typ: r("CustomWithdrawEventSig") },
        { json: "swapable", js: "swapable", typ: r("CustomWithdrawEventSig") },
        { json: "erc20", js: "erc20", typ: r("CustomWithdrawEventSig") },
        { json: "pos", js: "pos", typ: r("CustomWithdrawEventSig") },
        { json: "plasma", js: "plasma", typ: r("CustomWithdrawEventSig") },
        { json: "metaTx", js: "metaTx", typ: r("CustomWithdrawEventSig") },
        { json: "customWithdrawEventSig", js: "customWithdrawEventSig", typ: r("CustomWithdrawEventSig") },
        { json: "noDeposit", js: "noDeposit", typ: r("CustomWithdrawEventSig") },
        { json: "noWithdraw", js: "noWithdraw", typ: r("CustomWithdrawEventSig") },
    ], false),
    "CustomWithdrawEventSig": o([
        { json: "name", js: "name", typ: "" },
        { json: "description", js: "description", typ: "" },
    ], false),
    "Token": o([
        { json: "chainId", js: "chainId", typ: 0 },
        { json: "name", js: "name", typ: "" },
        { json: "symbol", js: "symbol", typ: "" },
        { json: "decimals", js: "decimals", typ: 0 },
        { json: "address", js: "address", typ: "" },
        { json: "logoURI", js: "logoURI", typ: "" },
        { json: "tags", js: "tags", typ: a(r("Tag")) },
        { json: "extensions", js: "extensions", typ: r("Extensions") },
    ], false),
    "Extensions": o([
        { json: "rootAddress", js: "rootAddress", typ: "" },
    ], false),
    "Version": o([
        { json: "major", js: "major", typ: 0 },
        { json: "minor", js: "minor", typ: 0 },
        { json: "patch", js: "patch", typ: 0 },
    ], false),
    "Tag": [
        "erc20",
        "metaTx",
        "plasma",
        "pos",
        "stablecoin",
        "swapable",
    ],
};
