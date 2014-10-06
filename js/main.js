/* global document , $, U*/
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
    setActive();
    $('.datepicker').datepicker({
        autoclose:true,
        format:"DD, M yyyy",
        todayHighlight:true});
});

function readDate(datestring) {
    return new Date(datestring);
}

function findStatusDone(elem) {
    return elem.innerHTML.match('DONE');
}

function muteTask() {
    var elems = document.getElementsByClassName('singletask');
    for (var i = 0; i< elems.length; i++) {
        if (findStatusDone(elems[i].children[0])) {
            elems[i].className=('singletask-done');
        }
    }
}

window.onload = setActive;
