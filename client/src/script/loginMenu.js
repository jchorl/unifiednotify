let React = require('react');
let Auth = require('./auth');
let Google = require('./google.js');

class LoginMenu extends React.Component {
	loginWithFacebook() {
		FB.login(function(resp) {
			if (resp.status === 'connected') {
				Auth.auth('fb', FB.getAccessToken());
			}
		});
	}
	loginWithGoogle() {
		Google.getAuth2().grantOfflineAccess({'redirect_uri': 'postmessage'})
		.then(function(resp) {
			Auth.auth('google', resp.code);
		});
	}
	render() {
		return (
			<div id="login-menu">
				<button className="login-button fb" onClick={this.loginWithFacebook}>
					<span className="icon"><i className="fa fa-facebook"></i></span>
					<span className="text">Login with Facebook</span>
				</button>
				<button className="login-button google" onClick={this.loginWithGoogle}>
					<span className="icon"><i className="fa fa-google"></i></span>
					<span className="text">Login with Google</span>
				</button>
			</div>
		);
	}
}

module.exports = LoginMenu;
