import React from 'react';
import Show_Users from './Show_Users.js';

class InfoPanel extends React.Component {
	constructor(props){
		super()
	}

   render() {
   		
   		let Info;
		Info = (
			<Show_Users/>
		)

      return (
		<div className="col-xs-12 col-sm-9 content">
            <div className="panel panel-default">
				<div className="panel-heading">
					<h3 className="panel-title">
						<a href="javascript:void(0);" className="toggle-sidebar">
							<span className="glyphicon glyphicon-transfer" data-toggle="offcanvas" title="Maximize Panel"></span>
						</a>
					</h3>
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