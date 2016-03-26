let Auth = require('./auth')
let auth2 = null;

window.googleCallback = function() {
	gapi.load('auth2', function(){
		// Retrieve the singleton for the GoogleAuth library and set up the client.
		auth2 = gapi.auth2.init({
			client_id: '991584590724-ait5qge3pq4g37re0sbse69ine6lnc6a.apps.googleusercontent.com',
			cookiepolicy: 'single_host_origin',
			scope: 'profile'
		});
	});
};

module.exports = {
	getAuth2: function() {
		return auth2;
	}
};
