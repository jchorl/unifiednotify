function authenticate(service, token) {
	fetch('/auth', {
		method: 'POST',
		body: JSON.stringify({
			service: service,
			token: token
		})
	})
	.then(function(res) {
		return res.text();
	})
	.then(function(jwt) {
		console.log(jwt);
	})
	.catch(function (error) {
		console.log(error);
	});
}
