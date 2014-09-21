"use strict";
var U = {};
U.$ = function(id) {
    return document.getElementById(id);

};

U.$cllist = function(cls) {
    return document.getElementsByClassName(cls);
};
