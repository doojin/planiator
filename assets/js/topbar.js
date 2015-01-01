var SHARE_POPUP         = '#share-popup';
var SHARE_BUTTON        = '#share-button';
var IMPORT_POPUP        = '#import-popup';
var IMPORT_BUTTON       = '#import-button';
var MY_CALENDARS_BUTTON = '#my-calendars-button';
var MY_CALENDARS_POPUP  = '#my-calendars-popup';

function initTopbar() {
    $(SHARE_BUTTON).on('click', showSharePopup);
    $(IMPORT_BUTTON).on('click', showImportPopup);
    $(MY_CALENDARS_BUTTON).on('click', showMyCalendarsPopup);

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