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

    <div id="my-calendars-popup" class="reveal-modal small" data-reveal>
        <dl class="accordion" data-accordion>
            <dd class="accordion-navigation">
                <a href="#c0" data-calendar="c0" style="background-color: #195C2B;">Family</a>
                <div id="c0" class="content">
                    <div class="row settings">
                        <div class="small-10 columns">
                            <div class="row collapse">
                                <div class="small-5 columns">
                                    <span class="prefix">Name:</span>
                                </div>
                                <div class="small-11 columns">
                                    <input type="text" data-calendar="c0" class="calendar-name" value="Family"/>
                                </div>
                            </div>
                        </div>
                        <div class="small-6 text-right columns">
                            <input type="text" data-calendar="c0" class="color-picker" value="#92cddc" />
                        </div>
                    </div>
                    <div class="row">
                        <div class="small-16 text-center">
                            <div class="red small pn-button">Delete Calendar</div>
                        </div>
                    </div>
                </div>

                <a href="#c1" data-calendar="c1" class="header" style="background-color: #5c2a4b;">University</a>
                <div id="c1" class="content">
                    <div class="row settings">
                        <div class="small-10 columns">
                            <div class="row collapse">
                                <div class="small-5 columns">
                                    <span class="prefix">Name:</span>
                                </div>
                                <div class="small-11 columns">
                                    <input type="text" data-calendar="c1" class="calendar-name" value="Family"/>
                                </div>
                            </div>
                        </div>
                        <div class="small-6 text-right columns">
                            <input type="text" data-calendar="c1" class="color-picker" value="#92cddc" />
                        </div>
                    </div>
                    <div class="row">
                        <div class="small-16 text-center">
                            <div class="red small pn-button">Delete Calendar</div>
                        </div>
                    </div>
                </div>
            </dd>
        </dl>

        <div class="row save-calendar-changes">
            <div class="small-16 columns text-center">
                <div class="green pn-button">Save Changes</div>
            </div>
        </div>

        <a class="close-reveal-modal">&#215;</a>
    </div>

    <div id="import-popup" class="reveal-modal small" data-reveal>
        <div class="row">
            <div class="small-16 columns">
                <input type="file" class="pn-button">
            </div>
        </div>
        <a class="close-reveal-modal">&#215;</a>
    </div>
{{end}}
