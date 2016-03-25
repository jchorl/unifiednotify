let React = require('react');

class LoginMenu extends React.Component {
	loginWithFacebook() {
		FB.login(function(resp) {
			if (resp.status === 'connected') {
				Auth.auth('fb', FB.getAccessToken());
			}
		});
	}
	render() {
		return <button onClick={this.loginWithFacebook}>Login with Facebook</button>
	}
}

module.exports = LoginMenu;
