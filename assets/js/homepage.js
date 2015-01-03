var SIGN_IN_FORM = $('#sign-in-form');
var SIGN_IN_BUTTON = $('#sign-in-button');
var SIGN_UP_FORM = $('#sign-up-form');
var SIGN_UP_BUTTON = $('#sign-up-button');

$(document).ready(function() {
    $(SIGN_IN_BUTTON).on("click", function() {
        $(SIGN_IN_FORM).submit();
    });

    $(SIGN_UP_BUTTON).on("click", function() {
        $(SIGN_UP_FORM).submit();
    });
});
