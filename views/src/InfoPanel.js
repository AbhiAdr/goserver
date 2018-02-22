import React from 'react';
import Show_Users from './Show_Users.js';
import Show_Testplan from './Show_Testplan.js';
import New_Testplan from './New_Testplan.js';

class InfoPanel extends React.Component {
	constructor(props){
		super(props)
		console.log(this.props.loadModule);
	}

   render() {

   		let Info;
   		if(this.props.loadModule=="c-admin"){
			Info = (
				<Show_Users/>
			)
   		}else if(this.props.loadModule=="users"){
   			console.log("users")
			Info = (
				<Show_Users/>
			)
   		}else if(this.props.loadModule=="testplan"){
   			console.log("users")
			Info = (
				<Show_Testplan/>
			)
   		}else if(this.props.loadModule=="newtestplan"){
   			console.log("New Testplan")
			Info = (
				<New_Testplan/>
			)
   		}

      return (
		<div className="col-xs-12 col-sm-9 content">
            <div className="panel panel-default">
				<div className="panel-heading">
					<div className="row">
						<div className="col-md-11">
							<h3 className="panel-title">
								<a href="javascript:void(0);" className="toggle-sidebar">
									<span className="glyphicon glyphicon-transfer " data-toggle="offcanvas" title="Maximize Panel"></span>
								</a>
							</h3>
						</div>
						<div className="col-md-1">
							<h3 className="panel-title">
								<a href="javascript:void(0);" className="toggle-sidebar">
									<span className="glyphicon glyphicon-file text-primary" data-toggle="offcanvas" title="Add New"></span>
								</a>
							</h3>
						</div>
					</div>
				</div>
              	<div className="panel-body">
              		{Info}
				</div>
			</div>
		</div>
   	  );
   }
}

export default InfoPanel;