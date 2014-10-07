/* global document , $, U*/

function activeLink(elem) {
    'use strict';
    return {"" : "home",
            "newtask" : "newtask",
            "about" : "about"}[elem];
}

function setActive() {
 	'use strict';
    var loc = document.URL.split('/');
    U.navbarSetActive(activeLink(loc[loc.length-1]));
}
$(document).ready(function() {
    setActive();
    muteTask();
    $('.datepicker').datepicker({
        autoclose:true,
        format:"DD, M dd yyyy",
        todayHighlight:true});
});

function readDate(datestring) {

    return new Date(datestring);
}

function filterDone() {
    "use strict";
    var done = $('.statustext').filter(function() {
        return $(this).text() === "TODO";
    });
    return done;

    // return $(done).parent().parent();
}



// muteTask - Check if task.Status == "DONE".
// if it is set textcolor to a light grey to make it less visible.
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
