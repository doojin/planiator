var SHARE_POPUP             = '#share-popup';
var SHARE_BUTTON            = '#share-button';
var IMPORT_POPUP            = '#import-popup';
var IMPORT_BUTTON           = '#import-button';
var MY_CALENDARS_BUTTON     = '#my-calendars-button';
var MY_CALENDARS_POPUP      = '#my-calendars-popup';
var COLOR_PICKERS           = '.color-picker';
var CALENDAR_NAME_INPUT     = '.calendar-name';
var DEFAULT_CALENDAR_NAME   = 'Untitled Calendar';

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
    
    updateColorPickers();
    
    $(COLOR_PICKERS).on('change.color', function(event, color) {
        updateCalendarHeader(this, color);
    });
    
    $(CALENDAR_NAME_INPUT).on('input', function() {
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