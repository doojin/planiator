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
                <li class="has-dropdown">
                    <a href="#">Filter Calendars</a>
                    <ul class="dropdown">
                        <li><a href="#">All Calendars</a></li>
                        <li><a href="#">Tasks</a></li>
                        <li><a href="#">Mettings</a></li>
                        <li><a href="#">My Calendar</a></li>
                    </ul>
                </li>
                <li>
                    <a id="my-calendars-button">My Calendars</a>
                </li>
            </ul>
            <ul class="right">
                <li>
                    <a id="share-button">Share Calendar</a>
                </li>
                <li>
                    <a id="import-button">Import Tasks</a>
                </li>
                <li>
                    <a href="#">Export Tasks</a>
                </li>
                <li>
                    <a href="/logoff">Log Off</a>
                </li>
            </ul>
        </section>
    </nav>
{{end}}
