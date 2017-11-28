import React from 'react';
import Nav from './Navigation.js';
import MainPanel from './MainPanel.js';

class Dashbord extends React.Component{

	constructor(props){
   		super()
   		console.log('constructor_Dashbord');
   		console.log(Login.Id);
   		this.state = {
   			UserFirstName : 'Loading',
   			UserLastName : 'Loading',
   			UserEmail: 'Loading',
   			login_status : 0
   		};
   	}

   	componentDidMount(){
   		console.log('componentDidMount_Dashbord');
   		var didMountscop= this;
   		$.ajax({
          type: "POST",
          url: Base.Url+"Dashbord",
            success: function(data){

				didMountscop.setState({
				        UserFirstName: data.fn,
				        UserLastName: data.ln,
				        UserEmail: data.email,
				        login_status: data.Login_status
			      });

            }
        });


   	}
	
	render() {
		console.log('render_Dashbord');
      return (
      	   	<div>
      	   		<Nav 
      	   			UserFirstName={this.state.UserFirstName}
      	   			UserLastName={this.state.UserLastName} />
              <MainPanel/>
      	   	</div>
   	  );
    }
}

export default Dashbord;