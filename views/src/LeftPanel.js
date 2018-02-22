import React from 'react';

class LeftPanel extends React.Component {
	
	constructor(props){
		super(props)
		console.log('constructor_LeftPanel');
	}

	handlechngModule(e){
		// console.log(e.target.className)
		this.props.toggModule(e.target.className)
	}

   render() {
	console.log('render_LeftPanel');

      return (
		<div className="col-xs-6 col-sm-3 sidebar-offcanvas" role="navigation">
			<ul className="list-group panel">
				<li className="list-group-item"><i className="glyphicon glyphicon-align-justify"></i> <b>SIDE PANEL</b></li>
				<li className="list-group-item"><input type="text" className="form-control search-query" placeholder="Search Something"></input></li>
				<li className="list-group-item">
					<a href="javascript:void(0)" 
					onClick={this.handlechngModule.bind(this)}
					className="users" >
					<i className="glyphicon glyphicon-th-list"></i>Users </a>
				</li>
			</ul>
		</div>
   	  );
   }
}

export default LeftPanel;