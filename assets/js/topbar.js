var SHARE_POPUP                 = '#share-popup';
var SHARE_BUTTON                = '#share-button';
var IMPORT_POPUP                = '#import-popup';
var IMPORT_BUTTON               = '#import-button';
var MY_CALENDARS_BUTTON         = '#my-calendars-button';
var MY_CALENDARS_POPUP          = '#my-calendars-popup';
var COLOR_PICKERS               = '.color-picker';
var CALENDAR_NAME_INPUT         = '.calendar-name';
var DEFAULT_CALENDAR_NAME       = 'Untitled Calendar';
var NEW_CALENDAR_BUTTON         = '#new-calendar-button';
var CALENDAR_LIST               = '#calendar-list';
var SAVE_CALENDAR_LIST_BUTTON   = '#save-calendar-list-button';
var DELETE_CALENDAR_BUTTON      = '.delete-calendar-button';

function initTopbar() {
    $(SHARE_BUTTON).on('click', showSharePopup);
    $(IMPORT_BUTTON).on('click', showImportPopup);
    $(MY_CALENDARS_BUTTON).on('click', showMyCalendarsPopup);
    $(NEW_CALENDAR_BUTTON).on('click', addNewCalendar);
    $(SAVE_CALENDAR_LIST_BUTTON).on('click', saveCalendarList);
    $(DELETE_CALENDAR_BUTTON).on('click', removeCalendar);

    initColorPickers();
}

// Shows popup with share link
function showSharePopup() {
    $(SHARE_POPUP).foundation('reveal', 'open');
}

// Shows popup with import form
function showImportPopup() {
    $(IMPORT_POPUP).foundation('reveal', 'open');
}

// Show popup with calendar settings
function showMyCalendarsPopup() {
    $(MY_CALENDARS_POPUP).foundation('reveal', 'open');
}

function initColorPickers() {
    // Do you think this is a motherfucking game?
    $('.color-picker').colorpicker({
        showOn:'button',
        defaultPalette: 'web',
        history: false,
        displayIndicator: false
    });

    updateColorPickers();

    $(COLOR_PICKERS).on('change.color', function(event, color) {
        updateCalendarHeader(this, color);
    });

    $(CALENDAR_NAME_INPUT).on('input', function() {
        updateCalendarHeaderText(this);
    });
}

function initCalendarRow(colorPicker, calendarNameInput) {
    $(colorPicker).colorpicker({
        showOn:'button',
        defaultPalette: 'web',
        history: false,
        displayIndicator: false
    });

    var wrapper = $(colorPicker).parent()[0];
    $(wrapper).css('width', '100%');

    $(colorPicker).on('change.color', function(event, color) {
        updateCalendarHeader(this, color);
    });

    $(calendarNameInput).on('input', function() {
        updateCalendarHeaderText(this);
    });
}

// Returns the header of calendar
function getCalendarHeader(element) {
    var calendarId = $(element).data('calendar');
    return $('a[data-calendar=' + calendarId + ']');
}

// Changes background color of header element
function setCalendarHeaderColor(header, color) {
    $(header).css('background-color', color);
}

// Gets calendar header by element and changes it's color
function updateCalendarHeader(element, color) {
    var header = getCalendarHeader(element);
    setCalendarHeaderColor(header, color);
}

// Makes color picker not editable + applies some styles to it's wrapper
function updateColorPickers() {
    var colorPickers = $(COLOR_PICKERS);
    $(colorPickers).attr('disabled', true);
    colorPickers.each(function(i) {
       var colorPicker = colorPickers[i];
       var wrapper = $(colorPicker).parent()[0];
       $(wrapper).css('width', '100%');
    });
}

// Gets calendar's header and chanes it's text
function updateCalendarHeaderText(element) {
    var header = getCalendarHeader(element);
    var text = encode($(element).val());
    if (text === '') {
        text = DEFAULT_CALENDAR_NAME;
    }
    $(header).html(text);
}

// Html symbols escaping + trim
function encode(str) {
    var el = document.createElement("div");
    el.innerText = el.textContent = str;
    str = el.innerHTML;
    return $.trim(str);
}

// Adds new calendar
function addNewCalendar() {
    $.post('/new-calendar', function(data) {
        var newCalendar = JSON.parse(data);
        displayNewCalendar(newCalendar);
    });
}

function displayNewCalendar(calendar) {
    var container = CALENDAR_LIST;
    var calendarLink = $(
        '<a href="#c{calendarId}" data-calendar="c{calendarId}" style="background-color: {calendarBg};">{calendarName}</a>'
            .replace(/{calendarId}/g, calendar.ID)
            .replace('{calendarBg}', calendar.Color)
            .replace('{calendarName}', calendar.Name)
    );
    var content = $('<div id="c{calendarId}" class="content">'.replace('{calendarId}', calendar.ID));
    var settingsRow = $('<div class="row settings">');
    var small10Columns = $('<div class="small-10 columns">');
    var rowCollapse = $('<div class="row collapse">');
    var small5Columns = $('<div class="small-5 columns">');
    var nameSpan = $('<span class="prefix">Name:</span>');
    var small11Columns = $('<div class="small-11 columns">');
    var calendarNameInput = $(
        '<input type="text" data-calendar="c{calendarId}" class="calendar-name" value="{calendarName}"/>'
            .replace('{calendarId}', calendar.ID)
            .replace('{calendarName}', calendar.Name)
    );
    var small6Columns = $('<div class="small-6 text-right columns">');
    var colorPicker = $(
        '<input type="text" data-calendar="c{calendarId}" class="color-picker" value="{calendarColor}" />'
            .replace('{calendarId}', calendar.ID)
            .replace('{calendarColor}', calendar.Color)
    );
    var row = $('<div class="row">');
    var small16Columns = $('<div class="small-16 text-center">');
    var deleteButton = $('<div class="red small pn-button">Delete Calendar</div>');

    $(small5Columns).append(nameSpan);
    $(small11Columns).append(calendarNameInput);
    $(rowCollapse).append(small5Columns);
    $(rowCollapse).append(small11Columns);
    $(small10Columns).append(rowCollapse);
    $(settingsRow).append(small10Columns);
    $(small6Columns).append(colorPicker);
    $(settingsRow).append(small6Columns);

    $(small16Columns).append(deleteButton);
    $(row).append(small16Columns);

    $(content).append(settingsRow);
    $(content).append(row);

    $(container).append(calendarLink);
    $(container).append(content);
    initCalendarRow(colorPicker, calendarNameInput);
}

// Sends request for saving actual calendar list
function saveCalendarList() {
    var container = CALENDAR_LIST;
    var rows = $(container).find('.content');

    var calendars = getCalendarsFromRows(rows);

    $.post('/save-calendars', {
        calendars: JSON.stringify(calendars)
    }, function() {
        $(MY_CALENDARS_POPUP).foundation('reveal', 'close');
    });
}

// Parses row to get calendars data
function getCalendarsFromRows(rows) {
    var list = [];
    rows.each(function(i) {
        var calendar = {};
        var row = rows[i];
        calendar.name = $(row).find('.calendar-name').val();
        calendar.color = $(row).find('.color-picker').val();
        calendar.id = $(row).attr('id').slice(1);
        list.push(calendar);
    });
    return list;
}

// Removes calendar from calendar list
function removeCalendar() {
    var calendar = $(this).data('calendar');
    var link = $('a[data-calendar="' + calendar + '"]');
    var content = $('#' + calendar);
    $(link).remove();
    $(content).remove();
}
