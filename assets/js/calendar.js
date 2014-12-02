$(document).ready(function() {
    $('.week li').on('click', showDayPopup);

    $('.add-new-record-button .pn-button').on('click', showAddNewPopup);

    // Timepicker widget
    var tpSettings = {
        'timeFormat': 'H:i',
        'step': 15
    };
    ['#from', '#to'].forEach(function(element) {
        $(element).timepicker(tpSettings);
    });

    $('.task input[type=checkbox]').on('change', doTask);
});

function showDayPopup() {
    $('#day-popup').foundation('reveal', 'open');
}

function showAddNewPopup() {
    $('#add-new-popup').foundation('reveal', 'open');
}

function revealAddNew() {
    $('#add-new').show();
}

function doTask() {
    var parent = $(this).parent();
    if ($(this).is(':checked')) {
        $(parent).addClass('done');
    } else {
        $(parent).removeClass('done');
    }
}
