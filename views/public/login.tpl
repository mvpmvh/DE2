{{template "public/base/base.tpl" append (params "IncludeLayoutCSS" true "IncludeLayoutJS" true "IncludeHeader" true) . }}
{{define "head"}}{{end}}
{{define "body"}}
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
                	    <p>{{i18n .Lang "login.description"}}</p>
                	{{end}}
            	</div>

            	<form action="{{.LoginTemplate.LoginAction}}" method="post" class="col-md-8">
                	{{range $key, $value := .HiddenInputs}}
                    	<input type="hidden" name="$key" value="$value" />
                	{{end}}

                	<div class="form-group">
                	    <label for="username">{{i18n .Lang "username"}}</label>
                	    <input class="form-control" autocapitalize="off" autocomplete="on" autocorrect="off" id="username" name="username" title="{{i18n .Lang "login description"}}" type="text" value="{{.Username}}" autofocus="autofocus" />
                	</div>

                	<div class="form-group">
                	    <label for="password">{{i18n .Lang "password"}}</label>
                	    <input class="form-control" autocapitalize="off" autocomplete="on" autocorrect="off" id="password" name="password" title="{{i18n .Lang "password"}}" type="password" />
                	</div>

                	<div id="loginButtonContainer" class="clearfix">
                	    <ul class="form-links pull-left">
                	        {{if .LoginTemplate.ShowForgot}}
                            	<li>
                            	    <a href="{{urlfor `MainController.Forgot`}}">{{i18n .Lang "forgot username or password"}}</a>
                            	</li>
							{{end}}
                        	{{if .LoginTemplate.ShowSignUp}}
	                            <li>
	                                <a href="{{urlfor `MainController.Signup` `newUser` `1`}}">Passcode/Create New User</a>
	                            </li>
							{{end}}
                    	</ul>
                    	<button id="loginButton" name="loginButton" type="submit" class="btn btn-de pull-right">&nbsp;&nbsp;&nbsp;{{i18n .Lang "login"}}&nbsp;&nbsp;&nbsp;</button>
                	</div>
            	</form>
        	</div>

        	{{if .LoginTemplate.ShowSignUp}}
            	<div id="sign-up" class="col-md-6">
            	    <h2>{{i18n .Lang "Register"}}</h2>
            	    <p>
            	        {{i18n .Lang "Register.why"}}
            	    </p>

            	    <ul id="freeAccess">
            	        <li>{{i18n .Lang "Register.access"}}</li>
            	        <li>{{i18n .Lang "Register.explore"}}</li>
            	        <li>{{i18n .Lang "Register.stay"}}</li>
            	    </ul>

            	    <div id="signUp" class="col-md-6">
            	        <div class="clearfix">
            	            <a href="{{.SignUpURL}}" class="btn btn-info btn-de">{{i18n .Lang "create account"}}</a>
            	        </div>
            	    </div>
            	</div>
        	{{end}}
    	</div>
{{end}}