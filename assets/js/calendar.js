var DAY_LI              = '.week li';
var NEW_RECORD_BUTTON   = '.add-new-record-button .pn-button';
var FROM_INPUT          = '#from';
var TO_INPUT            = '#to';
var TASK_CHECKBOX       = '.task input[type=checkbox]';
var DAY_POPUP           = '#day-popup';
var ADD_NEW_POPUP       = '#add-new-popup';
var HOVER_CLASS         = '.hover';
var SAVE_TASK_BUTTON    = '#save-new-task-button';

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
    $(SAVE_TASK_BUTTON).on('click', saveTask);

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
    var dayId = $(this).data('id');
    if (dayId) {
        $(DAY_POPUP).foundation('reveal', 'open');
        $('#new-task-day-id').val(dayId);
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

// Saves task to database
function saveTask() {
    $.post('/save-task', getTaskData(), function(data) {
        var response = JSON.parse(data);
        if (response.success) {
            appendTask();
        } else {
            displayNewTaskErrors(respose);
        }
    });
}

// Creates object with task data
function getTaskData() {
    return {
        title: $('#task-title').val(),
        from: $('#from').val(),
        to: $('#to').val(),
        calendarId: $('#task-calendar').val(),
        dayId: $('#new-task-day-id').val(),
        desc: $('#task-description').val()
    };
}

// Creates new task in task list
function appendTask() {

}

// Adds errors to the task form
function displayNewTaskErrors(response) {

}

// Hides all errors from task form
function hideNewTaskErrors() {

}
