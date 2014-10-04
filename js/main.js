/* global document , $ */
'use strict';

function activeLink(elem) {
    return {"" : "home",
            "newtask" : "newtask",
            "about" : "about"}[elem];
}

function setActive() {
    var loc = document.URL.split('/');
    U.navbarSetActive(activeLink(loc[loc.length-1]));
}
$(document).ready(function() {
    $('.datepicker').datepicker({
        autoclose:true,
        format:"DD-M-yyyy",
        todayHighlight:true});
});


window.onload = setActive;
