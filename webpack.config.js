module.exports = {
	entry: "./client/src/index.js",
	output: {
		path: "./client/dest",
		filename: "bundle.js"
	},
	module: {
		loaders: [
			{
				test: /\.less/,
				loader: "style!css!less"
			},
			{
				test: /\.js?$/,
				loader: 'babel-loader',
				exclude: /node_modules/,
				query: {
					presets: ['es2015', 'react']
				}
			}
		]
	}
};