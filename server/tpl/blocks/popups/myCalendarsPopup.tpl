{{define "myCalendarsPopup"}}
    <div id="my-calendars-popup" class="reveal-modal small" data-reveal>
        <div class="row">
            <div class="small-16 text-center">
                <div class="green pn-button" id="new-calendar-button">Add New</div>
            </div>
        </div>
        <dl class="accordion" data-accordion>
            <dd class="accordion-navigation" id="calendar-list">

                {{range .}}
                    {{template "calendarRow" .}}
                {{end}}

            </dd>
        </dl>

        <div class="row save-calendar-changes">
            <div class="small-16 columns text-center">
                <div class="green pn-button" id="save-calendar-list-button">Save Changes</div>
            </div>
        </div>

        <a class="close-reveal-modal">&#215;</a>
    </div>
{{end}}

{{define "calendarRow"}}
    <a href="#c{{.ID}}" data-calendar="c{{.ID}}" style="background-color: {{.Color}};">{{.Name}}</a>
    <div id="c{{.ID}}" class="content">
        <div class="row settings">
            <div class="small-10 columns">
                <div class="row collapse">
                    <div class="small-5 columns">
                        <span class="prefix">Name:</span>
                    </div>
                    <div class="small-11 columns">
                        <input type="text" data-calendar="c{{.ID}}" class="calendar-name" value="{{.Name}}" autocomplete="off"/>
                    </div>
                </div>
            </div>
            <div class="small-6 text-right columns">
                <input type="text" data-calendar="c{{.ID}}" class="color-picker" value="{{.Color}}" autocomplete="off"/>
            </div>
        </div>
        <div class="row">
            <div class="small-16 text-center">
                <div class="red small pn-button delete-calendar-button" data-calendar="c{{.ID}}">Delete Calendar</div>
            </div>
        </div>
    </div>
{{end}}
