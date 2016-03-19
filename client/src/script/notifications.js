module.exports = function() {
	fetch('/notifications', {
		method: 'GET',
		credentials: 'same-origin',
	})
	.then(function(res) {
		return res.json();
	})
	.then(function(parsed) {
		console.log(parsed);
	})
	.catch(function (error) {
		console.log(error);
	});
}
