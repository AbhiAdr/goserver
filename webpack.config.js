var path = require('path');

module.exports = {
  entry: './views/src/index.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, './public/js')
  },
  module:{
  	rules:
  	[
	  	{
	 		test:/\.js$/,
	 		use:
	 		{
	 			loader:'babel-loader'
	 		}
	  	}
  	]
  }
};