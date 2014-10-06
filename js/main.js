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
    muteTask();
    $('.datepicker').datepicker({
        autoclose:true,
        format:"DD, M yyyy",
        todayHighlight:true});
});

function readDate(datestring) {
    return new Date(datestring);
}


function muteTask() {
    var task = $('.taskstatus .statustext');
    $.each(task, function(index, value) {
        if (value.innerHTML == "DONE") {
            $(value).parent().parent().css('color', '#888888');
            $(value).parent().css('color','#777777');
            $(value).parent().parent().find('pre').css('color','#888888');

        }
    });
}

window.onload = setActive;
