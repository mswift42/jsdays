/*global document */
"use strict";
var U = {};
U.$ = function(id) {
    return document.getElementById(id);



};

U.$cllist = function(cls) {
    return document.getElementsByClassName(cls);
    
};

U.navbarSetActive = function(activeelement) {
    var elements = document.getElementById('navlinks').children;
    for (var i = 0;i<elements.length;i++) {
        if (activeelement === elements[i].id) {
        elements[i].className = "active";
        } else {
            elements[i].className = "";
        }
    }
};
