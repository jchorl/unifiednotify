var React = require('react');
var ReactDOM = require('react-dom');

var auth = require('./script/auth')
var fb = require('./script/fb')

class HelloMessage extends React.Component {
	render() {
		return <h1>Hello, Josh!</h1>
	}
}

ReactDOM.render(
	<HelloMessage />,
	document.getElementById('main')
);