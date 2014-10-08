
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
    var nday = new Date();
    return nday.setDate(day.getDate() + n);
}
