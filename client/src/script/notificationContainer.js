let React = require('react');
let Notification = require('./notification');
let GetNotifications = require('./notifications');

class NotificationContainer extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			notifications: [],
			loading: true
		};
	}

	componentDidMount() {
		GetNotifications(notifications => {
			this.setState({
				notifications: notifications,
				loading: false
			});
		});
	}

	render() {
		// first check if it is loading
		if (this.state.loading) {
			return <div className="loader">Loading...</div>;
		}

		let notifications = [];
		// if there are no notificaitons to display
		if (!this.state.notifications) {
			return <div id="main-content">No notifications to display</div>;
		}

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
