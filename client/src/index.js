var ReactDOM = require('react-dom');
var React = require('react');

var auth = require('./script/auth');
var fb = require('./script/fb');

var NotificationContainer = require('./script/notificationContainer.js');
var styles = require('./style/main.less');

ReactDOM.render(
	<NotificationContainer />,
	document.getElementById('main')
);