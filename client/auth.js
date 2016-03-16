function authenticate(service, token) {
	$.post('/auth',
		   JSON.stringify({
			   service: service,
			   token: token
		   }),
		   function(jwt) {
			   $.ajaxSetup({
				   headers: {
					   'Authorization': 'Bearer ' + jwt
				   }
			   });
		   });
}
