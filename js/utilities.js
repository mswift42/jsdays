"use strict";
var U = {};
U.$ = function(elem) {
    var target = elem.charAt(0);
    if (target === '#') {
        return document.getElementById(elem.substring(1));
    }
    return document.getElementsByClassName(elem.substring(1));

};
