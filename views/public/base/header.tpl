<!DOCTYPE html>
{{/* [if lt IE 7]><html class="no-js lt-ie9 lt-ie8 lt-ie7 bootstrapped"><![endif] */}}
{{/* [if IE 7]><html class="no-js lt-ie9 lt-ie8 bootstrapped"><![endif] */}}
{{/* [if IE 8]><html class="no-js lt-ie9 bootstrapped"><![endif] */}}
{{/* [if gt IE 8]><!--><html class="no-js bootstrapped"><!--<![endif] */}}
    
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