var SHARE_POPUP         = '#share-popup';
var SHARE_BUTTON        = '#share-button';
var IMPORT_POPUP        = '#import-popup';
var IMPORT_BUTTON       = '#import-button';

function initTopbar() {
    $(SHARE_BUTTON).on('click', showSharePopup);
    $(IMPORT_BUTTON).on('click', showImportPopup);
}

// Shows popup with share link
function showSharePopup() {
    $(SHARE_POPUP).foundation('reveal', 'open');
}

// Shows popup with import form
function showImportPopup() {
    $(IMPORT_POPUP).foundation('reveal', 'open');
}