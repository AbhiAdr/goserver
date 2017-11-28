import React from 'react';
import LeftPanel from './LeftPanel.js';
import InfoPanel from './InfoPanel.js';


class MainPanel extends React.Component {

	constructor(props){
		super(props)
		console.log('constructor_MainPanel');
	}

   render() {
	  console.log('render_MainPanel');
      return (
	    <div className="container-fluid">
			<div className="row row-offcanvas row-offcanvas-left">
				<LeftPanel/>
				<InfoPanel/>
			</div>
	    </div>
   	  );
   }
}

export default MainPanel;
