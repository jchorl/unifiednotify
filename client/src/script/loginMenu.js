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
		return (
			<div id="login-menu">
				<button className="login-button fb" onClick={this.loginWithFacebook}>
					<span className="icon"><i className="fa fa-facebook"></i></span>
					<span className="text">Login with Facebook</span>
				</button>
			</div>
		);
	}
}

module.exports = LoginMenu;
