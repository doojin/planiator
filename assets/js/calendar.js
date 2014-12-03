var DAY_LI              = '.week li';
var NEW_RECORD_BUTTON   = '.add-new-record-button .pn-button';
var FROM_INPUT          = '#from';
var TO_INPUT            = '#to';
var TASK_CHECKBOX       = '.task input[type=checkbox]';
var DAY_POPUP           = '#day-popup';
var ADD_NEW_POPUP       = '#add-new-popup';
var HOVER_CLASS         = '.hover';

$(document).ready(function() {
    $(DAY_LI).on('mouseenter', showOverlay);
    $(DAY_LI).on('mouseleave', hideOverlay);
    $(DAY_LI).on('click', showDayPopup);
    $(NEW_RECORD_BUTTON).on('click', showAddNewPopup);

    // Timepicker widget
    var tpSettings = {
        'timeFormat': 'H:i',
        'step': 15
    };
    [FROM_INPUT, TO_INPUT].forEach(function(element) {
        $(element).timepicker(tpSettings);
    });

    $(TASK_CHECKBOX).on('change', doTask);
});

// Shows popup with information about daily tasks
function showDayPopup() {
    $(DAY_POPUP).foundation('reveal', 'open');
}

// Shows popup with form for adding new task
function showAddNewPopup() {
    $(ADD_NEW_POPUP).foundation('reveal', 'open');
}

// Strikes out done task
function doTask() {
    var parent = $(this).parent();
    if ($(this).is(':checked')) {
        $(parent).addClass('done');
    } else {
        $(parent).removeClass('done');
    }
}

// If overlay is already displayed over the li element
function isOverlayDisplayed(element) {
    var elements = $(element).children(HOVER_CLASS);
    return elements.length > 0;
}

// If li element has data-id property
function hasDataId(element) {
    return !!$(element).data('id');
}

// Displays overlay over the li element
function showOverlay() {
    if (hasDataId(this) && !isOverlayDisplayed(this)) {
        $('<div/>', {
            class: 'hover',
            text: 'Show'
        }).appendTo(this);
    }
}

// Hides overlay from the li element
function hideOverlay() {
    if (isOverlayDisplayed(this)) {
        $(this).children(HOVER_CLASS).remove();
    }
}
