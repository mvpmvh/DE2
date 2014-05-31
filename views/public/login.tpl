{{template "public/header.tpl" params "IncludeLayoutCSS" true "IncludeLayoutJS" true "IncludeHeader" true}}
	<div class="outer-wrapper"> 
		{{if .ErrorMessages}}
			<div class="container errors">
    		    <div class="alert alert-danger alert-dismissable">
    		        <button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>
					{{range $error := .ErrorMessages}}
						{{$error}}
					{{end}}
				</div>
			</div>
		{{end}}
		
		<div class="de-public-container container">
    		<div id="loginDiv" class="row">
    		    <div id="userNameLogin" class="col-md-6">
    	        <h2>{{.LoginTemplate.Title}}</h2>

            	<div id="loginText">
            	    {{if .LoginTemplate.Body}}
            	        {{.LoginTemplate.Body}}
                	{{else}}
                	    <p>Use your existing Discovery Education username and password.</p>
                	{{end}}
            	</div>

            	<form action="{{.LoginTemplate.LoginAction}}" method="post" class="col-md-8">
                	{{range $key, $value := .HiddenInputs}}
                    	<input type="hidden" name="$key" value="$value" />
                	{{end}}

                	<div class="form-group">
                	    <label for="username">Username</label>
                	    <input class="form-control" autocapitalize="off" autocomplete="on" autocorrect="off" id="username" name="username" title="Username" type="text" value="{{.Username}}" autofocus="autofocus" />
                	</div>

                	<div class="form-group">
                	    <label for="password">Password</label>
                	    <input class="form-control" autocapitalize="off" autocomplete="on" autocorrect="off" id="password" name="password" title="Password" type="password" />
                	</div>

                	<div id="loginButtonContainer" class="clearfix">
                	    <ul class="form-links pull-left">
                	        {{if .LoginTemplate.ShowForgot}}
                            	<li>
                            	    <a href="{{urlfor `MainController.Forgot`}}">Forgot username or password?</a>
                            	</li>
							{{end}}
                        	{{if .LoginTemplate.ShowSignUp}}
	                            <li>
	                                <a href="{{urlfor `MainController.Signup` `newUser` `1`}}">Passcode/Create New User</a>
	                            </li>
							{{end}}
                    	</ul>
                    	<button id="loginButton" name="loginButton" type="submit" class="btn btn-de pull-right">&nbsp;&nbsp;&nbsp;LOGIN&nbsp;&nbsp;&nbsp;</button>
                	</div>
            	</form>
        	</div>

        	{{if .LoginTemplate.ShowSignUp}}
            	<div id="sign-up" class="col-md-6">
            	    <h2>Register</h2>
            	    <p>
            	        If you don't have a registered account with Discovery Education, create an account now. It's free and easy, and allows you to:
            	    </p>

            	    <ul id="freeAccess">
            	        <li>Access exclusive classroom resources to engage students</li>
            	        <li>Explore on-demand professional development resources</li>
            	        <li>Stay up-to-date with the latest news from Discovery Education<br />and its partners</li>
            	    </ul>

            	    <div id="signUp" class="col-md-6">
            	        <div class="clearfix">
            	            <a href="{{.SignUpURL}}" class="btn btn-info btn-de">Create an account</a>
            	        </div>
            	    </div>
            	</div>
        	{{end}}
    	</div>
	</div>
{{template "public/footer.tpl"}}