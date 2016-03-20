var React = require('react');

class Notification extends React.Component {
	render() {
		let imgContainerStyle = {
			backgroundImage: 'url(' + this.props.data.iconUrl + ')'
		}
		return (
			<div className="notification">
				<div style={imgContainerStyle} className="notification-img-container"></div>
				<div className="notification-text-container">
					<div className="line1">{this.props.data.line1}</div>
					<div className="line2">{this.props.data.line2}</div>
					<div className="line3">{this.props.data.line3}</div>
				</div>
			</div>
		);
	}
}

module.exports = Notification;
