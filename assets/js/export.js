var INPUT_DATE_FROM = '#from';
var INPUT_DATE_TO = '#to';

var DATEPICKER_DATE_FORMAT = 'dd.mm.yy';
var DAY_MONDAY = 1;

var TASK_LI = '.task-list li';

var CALENDAR_LABELS = '.calendar-list .calendar';
var CALENDAR_CBXS = '.calendars input[type=checkbox]';

var datepickerConfig = {
    dateFormat: DATEPICKER_DATE_FORMAT,
    firstDay: DAY_MONDAY
};

$(document).ready(function() {
    $(INPUT_DATE_FROM).datepicker(datepickerConfig);
    $(INPUT_DATE_TO).datepicker(datepickerConfig);

    $(TASK_LI).on('click', function(e) {
        if (e.target == getFirstCheckbox(this)) { return; }
        toggleTaskCheckboxState(this);
    });

    $(CALENDAR_LABELS).on('click', function(e) {
        if (e.target == getFirstCheckbox(this)) { return; }
        toggleCalendarCheckboxState(this);
    });

    $(CALENDAR_CBXS).on('change', function() {
        toggleCalendarTaskStatuses(this);
    });
});

// Checks checkbox if it is unchecked, unchecks checkbox if it is checked
function switchCheckbox(checkbox) {
    $(checkbox).trigger('click');
}

// Returns task checkbox
function getFirstCheckbox(element) {
    return $(element).find('input[type="checkbox"]') && $(element).find('input[type="checkbox"]')[0];
}

// Changes state of task checkbox
function toggleTaskCheckboxState(li) {
    var checkbox = getFirstCheckbox(li);
    !$(li).hasClass('invs') && checkbox && switchCheckbox(checkbox);
}

// Changes state of calendar checkbox
function toggleCalendarCheckboxState(div) {
    var checkbox = getFirstCheckbox(div);
    checkbox && switchCheckbox(checkbox);
}

// Changes state of all tasks which belong to calendar
function toggleCalendarTaskStatuses(calendarCbx) {
    var id = getCalendarId(calendarCbx);
    var cbxs = getCbxsByCalendarId(id);
    $(cbxs).each(function(id) {
        if ($(this).is(':checked') != $(calendarCbx).is(':checked')) {
            switchCheckbox(cbxs[id]);
        }
    });
}

// Returns calendar id
function getCalendarId(calendarCbx) {
    return $(calendarCbx).attr('id');
}

// Returns list of checkboxes by calendar id
function getCbxsByCalendarId(id) {
    var tasks = $('.tasks');
    var checkboxes = $(tasks).find('input[type="checkbox"][data-calendar="' + id + '"]');
    return checkboxes;
}
