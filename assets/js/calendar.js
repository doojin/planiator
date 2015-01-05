var DAY_LI              = '.week li';
var NEW_RECORD_BUTTON   = '.add-new-record-button .pn-button';
var FROM_INPUT          = '#from';
var TO_INPUT            = '#to';
var TASK_CHECKBOX       = '.task input[type=checkbox]';
var DAY_POPUP           = '#day-popup';
var ADD_NEW_POPUP       = '#add-new-popup';
var HOVER_CLASS         = '.hover';

$(document).ready(function() {

    initTopbar();

    $(DAY_LI).on('mouseenter', function() {
        showOverlay(this);
    });

    $(DAY_LI).on('mouseleave', function() {
        hideOverlay(this);
    });

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

    $(TASK_CHECKBOX).on('change', function() {
        doTask(this);
    });
});

// Shows popup with information about daily tasks
function showDayPopup() {
    if ($(this).data('id')) {
        $(DAY_POPUP).foundation('reveal', 'open');
    }
}

// Shows popup with form for adding new task
function showAddNewPopup() {
    $(ADD_NEW_POPUP).foundation('reveal', 'open');
}

// Strikes out done task
function doTask(element) {
    var parent = $(element).parent();
    if ($(element).is(':checked')) {
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
function showOverlay(element) {
    if (hasDataId(element) && !isOverlayDisplayed(element)) {
        $('<div/>', {
            class: 'hover',
            text: 'Show'
        }).appendTo(element);
    }
}

// Hides overlay from the li element
function hideOverlay(element) {
    if (isOverlayDisplayed(element)) {
        $(element).children(HOVER_CLASS).remove();
    }
}
