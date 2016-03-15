function authenticate(service, token) {
	$.post('/auth',
		   JSON.stringify({
			   service: service,
			   token: token
		   }),
		   function(jwt) {
			   console.log("got jwt: " + jwt);
		   });
}
