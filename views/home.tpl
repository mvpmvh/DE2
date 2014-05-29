<!DOCTYPE html>

<html>
  	<head>
    	<title>DE 2.0</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	
		<!-- Latest compiled and minified CSS -->
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">

		<!-- Optional theme -->
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
	</head>
  	
  	<body>
  		<div class="container">
			{{ if .Message }}
				<div class="alert alert-success">
					{{ .Message }}
				</div>
			{{ end }}
				
			{{ if .ErrorMessage }}
				<div class="alert alert-danger">
					{{ .ErrorMessage }}
				</div>
			{{ end }}

    	  <form action="/login" method="POST" class="form-signin" role="form">
    	    <h2 class="form-signin-heading">Please sign in</h2>
    	    <input id="email" name="email" type="email" class="form-control" placeholder="Email address" required autofocus>
    	    <input id="password" name="password" type="password" class="form-control" placeholder="Password" required>
    	    <label class="checkbox">
    	      <input type="checkbox" value="remember-me"> Remember me
    	    </label>
    	    <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
    	  </form>

    	</div> <!-- /container -->

		<!-- Latest compiled and minified JavaScript -->		
		<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
	</body>
</html>