{{define "content"}}
    <div class="calendar-content">
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
            <ul class="week small-block-grid-1 medium-block-grid-7">
                <li class="hide-for-small-only"></li>
                <li class="hide-for-small-only"></li>
                <li class="hide-for-small-only"></li>
                <li class="hide-for-small-only"></li>
                <li data-id="01012014">
                    <div class="day">1</div>
                </li>
                <li data-id="02012014">
                    <div class="day">2</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="03012014">
                    <div class="day">3</div>
                </li>
            </ul>
            <ul class="week small-block-grid-1 medium-block-grid-7">
                <li data-id="04012014">
                    <div class="day">4</div>
                </li>
                <li data-id="05012014">
                    <div class="day">5</div>
                    <div class="task">Task1</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="06012014">
                    <div class="day">6</div>
                </li>
                <li data-id="07012014">
                    <div class="day">7</div>
                </li>
                <li data-id="08012014">
                    <div class="day">8</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="09012014">
                    <div class="day">9</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="10012014">
                    <div class="day">10</div>
                </li>
            </ul>
            <ul class="week small-block-grid-1 medium-block-grid-7">
                <li data-id="11012014">
                    <div class="day">11</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="12012014">
                    <div class="day">12</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="13012014">
                    <div class="day">13</div>
                </li>
                <li data-id="14012014">
                    <div class="day">14</div>
                    <div class="task">Task1</div>
                    <div class="task">Task1</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="15012014">
                    <div class="day">15</div>
                </li>
                <li data-id="16012014">
                    <div class="day">16</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="17012014">
                    <div class="day">17</div>
                </li>
            </ul>
            <ul class="week small-block-grid-1 medium-block-grid-7">
                <li data-id="18012014">
                    <div class="day">18</div>
                    <div class="task">Task1</div>
                    <div class="task">Task1</div>
                    <div class="task">Task1</div>
                    <div class="more">...</div>
                </li>
                <li data-id="19012014">
                    <div class="day">19</div>
                </li>
                <li data-id="20012014">
                    <div class="day">20</div>
                </li>
                <li data-id="21012014">
                    <div class="day">21</div>
                </li>
                <li data-id="22012014">
                    <div class="day">22</div>
                    <div class="task">Task1 Task1 Task1 Task1 Task1 Task1 Task1</div>
                </li>
                <li data-id="23012014">
                    <div class="day">23</div>
                </li>
                <li data-id="24012014">
                    <div class="day">24</div>
                    <div class="task">Task1</div>
                </li>
            </ul>
            <ul class="week small-block-grid-1 medium-block-grid-7">
                <li data-id="25012014">
                    <div class="day">25</div>
                </li>
                <li data-id="26012014">
                    <div class="day">26</div>
                    <div class="done task">Task1</div>
                </li>
                <li data-id="27012014">
                    <div class="day">27</div>

                </li>
                <li data-id="28012014">
                    <div class="day">28</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="29012014">
                    <div class="day">29</div>
                    <div class="task">Task1</div>
                </li>
                <li data-id="30012014">
                    <div class="day">30</div>
                </li>
                <li class="hide-for-small-only"></li>
            </ul>
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
