module.exports = function(cb) {
	fetch('/notifications', {
		method: 'GET',
		credentials: 'same-origin',
	})
	.then(function(res) {
		return res.json();
	})
	.then(cb)
	.catch(function (error) {
		console.log(error);
	});
};
