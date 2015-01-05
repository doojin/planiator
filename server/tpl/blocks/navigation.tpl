{{define "navigation"}}
    <nav class="top-bar" role="navigation" data-topbar>
        <ul class="title-area">
            <li class="name"></li>
            <!-- Menu icon -->
            <li class="toggle-topbar menu-icon">
                <a href="#">
                    <span></span>
                </a>
            </li>
        </ul>
        <section class="top-bar-section">
            <ul class="left">
                <li>
                    <a id="my-calendars-button">My Calendars</a>
                </li>
            </ul>
            <ul class="right">
                <li>
                    <a href="/logoff">Log Off</a>
                </li>
            </ul>
        </section>
    </nav>
{{end}}
