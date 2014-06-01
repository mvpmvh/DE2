<!DOCTYPE html>
<html>
	<head>
		{{template "public/base/header.tpl" .}}
		{{template "head" .}}
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
                        <a href="{{urlfor `SessionController.Login`}}" class="btn btn-de"> <span class="glyphicon glyphicon-lock"></span>{{i18n .Lang "login"}}</a>
                    </div>
                {{end}}
            </div>
        </div>
	
		<noscript>{{i18n .Lang "enable_js_prompt"}}</noscript>
		
		<div class="outer-wrapper"> 
			{{template "body" .}}
		</div>

		{{template "public/base/footer.tpl" .}}
	</body>
</html>