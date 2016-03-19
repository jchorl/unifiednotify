function authenticate(service, token) {
	fetch('/auth', {
		method: 'POST',
		credentials: 'same-origin',
		body: JSON.stringify({
			service: service,
			token: token,
		})
	})
	.catch(function (error) {
		console.log(error);
	});
}
