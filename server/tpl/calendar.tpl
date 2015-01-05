{{define "content"}}
    <div class="calendar-content">
        <div class="calendar-navigation text-center">
            <a href="/calendar/m{{PrevOffset .Offset}}"><i class="fi-arrow-left"></i></a>
            <div class="month-year">{{.MonthName}} {{.YearName}}</div>
            <a href="/calendar/m{{NextOffset .Offset}}"><i class="fi-arrow-right large"></i></a>
        </div>
        <div class="calendar">
            <ul class="day-names hide-for-small-only small-block-grid-1 medium-block-grid-7">
                <li>Monday</li>
                <li>Tuesday</li>
                <li>Wednesday</li>
                <li>Thursday</li>
                <li>Friday</li>
                <li>Saturday</li>
                <li>Sunday</li>
            </ul>
            {{range .Weeks}}
                {{template "week" .}}
            {{end}}
        </div>

        <div id="day-popup" class="reveal-modal small" data-reveal>
            <div class="row">
                <div class="day-of-week small-8 columns">Monday</div>
                <div class="date small-8 columns text-right">21.03.1991</div>
            </div>

            <dl class="accordion" data-accordion>
                <dd class="accordion-navigation">

                    <a href="#task1">Task1</a>
                    <div id="task1" class="content">
                        <div class="row">
                            <div class="time small-16 medium-9 columns">
                                From <span><i class="fi-clock"></i> 12:30</span> to <span><i class="fi-clock"></i> 14:30</span>
                            </div>
                            <div class="calendar small-16 medium-7 columns">
                                <i class="fi-star"></i> Family
                            </div>
                        </div>
                        <div class="row">
                            <div class="description small-16 columns">
                                Need to visit my grandma
                            </div>
                        </div>
                        <div class="row">
                            <div class="controls small-16 columns text-center">
                                <div class="green pn-button">Mark as Done</div>
                                <div class="pn-button">Edit Task</div>
                                <div class="red pn-button">Delete Task</div>
                            </div>
                        </div>
                    </div>

                    <a href="#task2" class="done">Clean my room</a>
                    <div id="task2" class="content">
                        <div class="row">
                            <div class="time small-16 medium-9 columns">
                                <span><i class="fi-clock"></i> All the day</span>
                            </div>
                            <div class="calendar small-16 medium-7 columns">
                                <i class="fi-star"></i> House
                            </div>
                        </div>
                        <div class="row">
                            <div class="controls small-16 columns text-center">
                                <div class="green pn-button">Mark as Undone</div>
                                <div class="pn-button">Edit Task</div>
                                <div class="red pn-button">Delete Task</div>
                            </div>
                        </div>
                    </div>

                </dd>
            </dl>

            <div class="add-new-record-button row">
                <div class="small-16 text-center columns">
                    <div class="green pn-button">Create New Task</div>
                </div>
            </div>
            <a class="close-reveal-modal">&#215;</a>
        </div>

        <div id="add-new-popup" class="reveal-modal small" data-reveal>
            <form>
                <div class="row collapse">
                    <div class="small-5 columns">
                        <span class="prefix">Task Title:</span>
                    </div>
                    <div class="small-11 columns">
                        <input type="text">
                    </div>
                </div>
                <div class="row collapse">
                    <div class="small-7 columns">
                        <div class="row collapse">
                            <div class="small-6 columns">
                                <span class="prefix">From:</span>
                            </div>
                            <div class="small-10 columns">
                                <input type="text" id="from">
                            </div>
                        </div>
                    </div>
                    <div class="small-2 columns"></div>
                    <div class="small-7 columns">
                        <div class="row collapse">
                            <div class="small-6 columns">
                                <span class="prefix">To:</span>
                            </div>
                            <div class="small-10 columns">
                                <input type="text" id="to">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="row collapse">
                    <div class="small-5 columns">
                        <span class="prefix">Task Type:</span>
                    </div>
                    <div class="small-11 columns">
                        <select>
                            <option value="1">Family</option>
                            <option value="2">Work</option>
                            <option value="3">University</option>
                            <option value="4">Meeting</option>
                        </select>
                    </div>
                </div>
                <div class="row">
                    <div class="small-16 text-center columns">
                        <div class="blue pn-button">Save</div>
                    </div>
                </div>
            </form>
            <a class="close-reveal-modal">&#215;</a>
        </div>
    </div>
{{end}}

{{define "week"}}
<ul class="week small-block-grid-1 medium-block-grid-7">
    {{range .}} {{if .Id}} {{template "day" .}} {{else}} {{template "emptyday"}} {{end}} {{end}}
</ul>
{{end}}

{{define "day"}}
    <li data-id=".Id">
        <div class="day">{{.Day}}</div>
    </li>
{{end}}

{{define "emptyday"}}
    <li class="hide-for-small-only"></li>
{{end}}
