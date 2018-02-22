import React from 'react';
import LeftPanel from './LeftPanel.js';
import InfoPanel from './InfoPanel.js';


class MainPanel extends React.Component {

	constructor(props){
		super(props)
   		this.state = {
   			moduleClicked: ''
   		}
		console.log('constructor_MainPanel');
	}

	toggModule(module){

		console.log(module);
		this.setState({
			moduleClicked: module
		})
	}

   render() {
	  console.log('render_MainPanel');
      return (
	    <div className="container-fluid">
			<div className="row row-offcanvas row-offcanvas-left">
				<LeftPanel toggModule={this.toggModule.bind(this)}/>
				<InfoPanel loadModule={this.state.moduleClicked}/>
			</div>
	    </div>
   	  );
   }
}

export default MainPanel;
