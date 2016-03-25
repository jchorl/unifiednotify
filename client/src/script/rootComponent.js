let React = require('react');

let Auth = require('./auth');
let LoginMenu = require('./loginMenu.js');
let NotificationContainer = require('./notificationContainer.js');

class RootComponent extends React.Component {
	render() {
		if (Auth.isAuthd()) {
			return <NotificationContainer />;
		}
		return <LoginMenu />;
	}
}

module.exports = RootComponent;
