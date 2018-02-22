import React from 'react';
import ReactDOM from 'react-dom';
import Login from './login.js';
import Dbconfig from './Dbconfig.js';
import Signup from './signup.js';
import Dashbord from './Dashbord.js';

if (pageName=='Login'){

	ReactDOM.render(<Login />, document.getElementById('app'));
}else if(pageName=='Dbconfig'){

	ReactDOM.render(<Dbconfig />, document.getElementById('app'));
}else if(pageName=='Signup'){

	ReactDOM.render(<Signup />, document.getElementById('app'));
}else if(pageName=='Index'){
	console.log('here');
	ReactDOM.render(<Dashbord />, document.getElementById('app'));
}