let ReactDOM = require('react-dom');
let React = require('react');

let fb = require('./script/fb');
let google = require('./script/google');
let RootComponent = require('./script/rootComponent.js');
let styles = require('./style/main.less');

ReactDOM.render(
	<RootComponent />,
	document.getElementById('main')
);