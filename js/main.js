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
    agendaHtml(new Date());
    $('.datepicker').datepicker({
        autoclose:true,
        format:"DD, M dd yyyy",
        todayHighlight:true});
});

// return jquery object of .singletask divs
// with a taskstatus of "TODO".
function filterTodo() {
    "use strict";
    return $('.singletask').filter(function() {
        return $(this).find('.statustext').text() === "TODO";
    });
}

function readScheduled(task) {
    'use strict';
    return new Date($(task).find('.scheduleddate').text());
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

function Agenda(day) {
    'use strict';
    this.day = day;
    this.tasks = [];
}

function buildAgenda(startday) {
    'use strict';
    var wd = weekDates(startday);
    var res = [];
    var tasks = filterTodo();
    for (var i = 0; i<wd.length; i++) {
        var ag = new Agenda(wd[i]);
        for (var j = 0; j<tasks.length; j++) {
            if (sameDay(wd[i], readScheduled(tasks[j]))) {
                ag.tasks.push(tasks[j]);
            }
        }
        res.push(ag);
    }
    return res;
}

function agendaHtml(startday) {
    'use strict';
    var ag = buildAgenda(startday);
    var target = $('.agenda');
    var ht = "";
    for (var i = 0; i < ag.length; i++) {

        ht += "<div class='singleag'>" + "<p>" +
            formatDate(ag[i].day) + "</p>" + "</div>";
        for (var j = 0; j < ag[i].tasks.length;j++) {
            ht += "<div class='singleagtask'>" +
                "<a href=#" + $(ag[i].tasks[j]).find('#hiddenid').val() + ">" +
                $(ag[i].tasks[j]).find('h4').text() + "</a></div>";
        }

    }
    $(target).append(ht);
}

function muteNonTask() {
    'use strict';
    var ag = $('.singleag');
    for (var i = 0; i < ag.length; i++) {
        if ($(ag[i]).next().find('a').length===0) {
            $(ag[i]).css('color',"#a2a2a2");
        }
    }
}


        



window.onload = setActive;
