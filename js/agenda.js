
// indToMonth - return the month abbreviation
// for the month index of a date object.
function indToMonth(i) {
    'use strict';
    return ['Jan','Feb','Mar',
            'Apr','May', 'Jun',
            'Jul', 'Aug', 'Sep',
            'Oct', 'Nov', 'Dec'][i];
}
// indToWeekDay - return the Weekday
// for the getDay index method of a date object.
function indToWeekDay(i) {
    'use strict';
    return ['Sunday', 'Monday', 'Tuesday',
            'Wednesday', 'Thursday', 'Friday',
            'Saturday'][i];
}
function readDate(datestring) {
    return new Date(datestring);
}
// addDays - add n days to a given
// date object.
function addDays(day, n) {
    'use strict';
    var nday = new Date();
    return nday.setDate(day.getDate() + n);
}
// buildAgenda - return an array of 10 days
// from startday - 3 days to + 7 days.
function weekDates(startday) {
    'use strict';
    var dates = [];
    for (var i = -3 ; i<8 ; i++) {
        var day = new Date(addDays(startday,i));
        dates.push(day);
    }
    return dates;
}


// sameDay - compare two dates -
// if both are on the same day return
// true else false.
function sameDay(day1, day2) {
    'use strict';
    return day1.getDate() === day2.getDate();
}

function formatDate(day) {
    'use strict';
    return indToWeekDay(day.getDay()) + ", " + indToMonth(day.getMonth()) + " " + day.getDate() + " " + day.getFullYear();
}
