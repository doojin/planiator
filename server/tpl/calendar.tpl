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
    </div>
{{end}}

{{define "week"}}
<ul class="week small-block-grid-1 medium-block-grid-7">
    {{range .}} {{if .Id}} {{template "day" .}} {{else}} {{template "emptyday"}} {{end}} {{end}}
</ul>
{{end}}

{{define "day"}}
    <li data-id="{{.Id}}">
        <div class="day">{{.Day}}</div>
    </li>
{{end}}

{{define "emptyday"}}
    <li class="hide-for-small-only"></li>
{{end}}
