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
    assert.equal(day2.getMonth(), 9);
    assert.equal(day2.getDate(), 2);
    assert.equal(day3.getDate(), 3);
    assert.equal(day4.getDate(), 4);
    assert.equal(indToWeekDay(day4.getDate()), "Thursday");
});
