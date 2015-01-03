var SIGN_IN_FORM = $('#sign-in-form');
var SIGN_IN_BUTTON = $('#sign-in-button');

$(document).ready(function() {
    $(SIGN_IN_BUTTON).on("click", function() {
        $(SIGN_IN_FORM).submit();
    });
});
