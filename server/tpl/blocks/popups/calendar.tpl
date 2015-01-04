{{define "popups"}}
    <div id="share-popup" class="reveal-modal small" data-reveal>
        <div class="row">
            <div class="small-16 columns">
                Your calendar is shared now:
            </div>
            <div class="small-16 columns">
                <input type="text" value="http://planiator.tk/c/274927391203">
            </div>
        </div>
        <a class="close-reveal-modal">&#215;</a>
    </div>

    {{template "myCalendarsPopup" .Calendars}}

    <div id="import-popup" class="reveal-modal small" data-reveal>
        <div class="row">
            <div class="small-16 columns">
                <input type="file" class="pn-button">
            </div>
        </div>
        <a class="close-reveal-modal">&#215;</a>
    </div>
{{end}}
