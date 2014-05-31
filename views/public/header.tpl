<!DOCTYPE html>
{{/* [if lt IE 7]><html class="no-js lt-ie9 lt-ie8 lt-ie7 bootstrapped"><![endif] */}}
{{/* [if IE 7]><html class="no-js lt-ie9 lt-ie8 bootstrapped"><![endif] */}}
{{/* [if IE 8]><html class="no-js lt-ie9 bootstrapped"><![endif] */}}
{{/* [if gt IE 8]><!--><html class="no-js bootstrapped"><!--<![endif] */}}
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <title>
            <cfoutput>#rc.pageTitle#</cfoutput>
        </title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">

        <script>window.scrollTo(0,1);</script>

		{{if .IncludeLayoutCSS}}
            <link rel="stylesheet" type="text/css" href="/static/css/bootstrap/bootstrap.min.css">
            <link rel="stylesheet" type="text/css" href="/static/css/public/main/main.css">
        {{end}}
        
        <script type="text/javascript" src="/static/js/modernizr/modernizr-2.5.3.min.js"></script>

        {{/* HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries */}}
        {{/* if lt IE 9 */}}
          <script src="/static/js/html5shiv/html5shiv.js"></script>
          <script src="/static/js/respond/respond.min.js"></script>
        {{/* endif */}}
    </head>
	
	<body>
		{{/*<cfif local.isExternalUser>
            <cfoutput>
                <div id="return-container">
                    <a href="#local.returnUrl#" id="return-link">Return to #local.partnerName#</a>
                </div>
            </cfoutput>
        </cfif>
		*/}}
		
        <div id="shell-outer">
            <div id="shell" class="container clearfix">
                <div id="de-logo" class="pull-left">
                    <a href="{{config `string` `wwwurl` `google.com`}}" title="Discovery Education"></a>
                </div>

                {{if not .User}}
                    <div class="pull-right" id="public-login-btn">
                        <a href="{{urlfor `SessionController.Login`}}" class="btn btn-de"> <span class="glyphicon glyphicon-lock"></span> LOGIN</a>
                    </div>
                {{end}}

            </div>
        </div>