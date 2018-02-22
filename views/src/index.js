import React from 'react';
import ReactDOM from 'react-dom';
import Login from './login.js';
import Signup from './signup.js';
import Dashbord from './Dashbord.js';

if (pageName=='Login'){

	ReactDOM.render(<Login />, document.getElementById('app'));
}else if(pageName=='Signup'){

	ReactDOM.render(<Signup />, document.getElementById('app'));
}else if(pageName=='Index'){
	ReactDOM.render(<Dashbord />, document.getElementById('app'));
}