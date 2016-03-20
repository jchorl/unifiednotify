var React = require('react');
var Notification = require('./notification');
var GetNotifications = require('./notifications');

class NotificationContainer extends React.Component {
	constructor(props) {
		super(props);
		this.state = {notifications: []};
	}

	componentDidMount() {
		GetNotifications(notifications => this.setState({
			notifications: notifications
		}));
	}

	render() {
		var notifications = [];
		this.state.notifications.map(n => notifications.push(<Notification key={n.id} data={n} />));
		return (
			<div id="main-content">
				<div id="notification-container">
					{notifications}
				</div>
			</div>
		);
	}
}

module.exports = NotificationContainer;
