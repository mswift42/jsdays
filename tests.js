QUnit.test( "hello test", function( assert ) {
  assert.ok( 1 == "1", "Passed!" );
});

QUnit.test("test indToMonth", function( assert ) {
    assert.equal(indToMonth(2) , "Mar");
    assert.equal(indToMonth(0) , "Jan");
    assert.equal(indToMonth(11) , "Dec");
    assert.notEqual(indToMonth(4), "Apr");
});

QUnit.test("test indToWeekDay", function(assert) {
    assert.equal(indToWeekDay(0), "Sunday");
    assert.equal(indToWeekDay(1), "Monday");
    assert.equal(indToWeekDay(6), "Saturday");
    assert.equal(indToWeekDay(3), "Wednesday");
    assert.notEqual(indToWeekDay(3), "Tuesday");
});

QUnit.test("test addDays", function(assert) {
    var day1 = new Date(2014,10,1);
    var day2 = new Date(addDays(day1,1));
    var day3 = new Date(addDays(day1,2));
    var day4 = new Date(addDays(day2,2));
    var day5 = new Date(addDays(day3,-2)); // check if subtraction of days works
    var day6 = new Date(addDays(day1,-1));
    var day7 = new Date(addDays(day1,0)); // check if no additions returns same day.
    assert.equal(day2.getMonth(), 9);
    assert.equal(day2.getDate(), 2);
    assert.equal(day3.getDate(), 3);
    assert.equal(day4.getDate(), 4);
    assert.equal(indToWeekDay(day4.getDate()), "Thursday");
    assert.equal(day5.getDate(), 1);
    assert.equal(day6.getDate(),30);
    assert.equal(day6.getMonth(),8);
    assert.equal(day7.getDate(),1);
});

QUnit.test("test buildAgenda", function(assert) {
    var day = new Date(2014,10,1);
    var ag = buildAgenda(day);
    assert.equal(ag[0].getDate(),28);
    assert.equal(ag[0].getMonth(), 8);
    assert.equal(ag[3].getDate(),1);
    assert.equal(ag[10].getDate(),8);
    assert.equal(ag.length,11);
});
