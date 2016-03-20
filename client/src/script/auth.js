module.exports = function(service, token) {
	fetch('/auth', {
		method: 'POST',
		credentials: 'same-origin',
		body: JSON.stringify({
			service: service,
			token: token,
		})
	})
	.then(function(res) {
		if (res.ok) {
			console.log('authd')
		}
	})
	.catch(function (error) {
		console.log(error);
	});
};
