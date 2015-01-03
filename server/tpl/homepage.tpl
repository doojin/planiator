{{define "content"}}
    <div class="homepage-content row">
        <div class="small-16 columns">
            <div class="main">
                <div class="row">
                    <div class="small-16 medium-7 columns sign-in">
                        <form method="POST" action="/" id="sign-in-form">
                            <input type="hidden" name="action" value="sign up"/>
                            <div class="row">
                                <div class="small-16 columns">
                                    <label for="email-sign-in">Your E-mail Account:</label>
                                    <input type="text" name="email" id="email-sign-in"/>
                                </div>
                            </div>
                            <div class="row">
                                <div class="small-16 columns">
                                    <label for="password-sign-in">Your Password:</label>
                                    <input type="password" name="password" id="password-sign-in"/>
                                </div>
                            </div>
                            <div class="remember-me row">
                                <div class="small-16 columns">
                                    <input type="checkbox" id="remember-me-sign-in" id="remember-me-sign-in">
                                    <label for="remember-me-sign-in">Remember Me</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="green pn-button" id="sign-in-button">
                                    Sign In
                                </div>
                            </div>
                        </form>
                    </div>
                    <div class="small-16 medium-8 columns sign-up">
                        <form>
                            <div class="row">
                                <div class="small-16 columns">
                                    <label for="email-sign-up">Your E-mail:</label>
                                    <input type="text" name="email" id="email-sign-up"/>
                                </div>
                            </div>
                            <div class="row">
                                <div class="small-16 columns">
                                    <label for="password-sign-up">Pick Up Password:</label>
                                    <input type="password" name="password" id="password-sign-up"/>
                                </div>
                            </div>
                            <div class="row">
                                <div class="small-16 columns">
                                    <label for="password-repeat-sign-up">Your Password Again:</label>
                                    <input type="password" name="password-repeat" id="password-repeat-sign-up"/>
                                </div>
                            </div>
                            <div class="row">
                                <div class="pn-button">
                                    Sign Up
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
{{end}}
{{define "scripts"}}
    <script src="/assets/vendor/jquery/jquery.min.js"></script>
    <script src="/assets/js/homepage.js"></script>
{{end}}