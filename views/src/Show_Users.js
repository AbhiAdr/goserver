import React from 'react';

class Show_Users extends React.Component{
	constructor(props){
		super(props)
 		this.state = {
   			data : {},
   			tr : ''
   		}
	}

	componentDidMount() {
        $.ajax({
          type: "POST",
          url: Base.Url+"Users",
          cache: true,
            success: function(data){
            	this.setState({
            		data : data
            	});
            }.bind(this)
        });
    }

	render() {
		
		console.log('render');
		let Info;
		if (this.state.data.length) {

			Info = $.map(this.state.data, function(value,key){
			    return (
			    	<tr key={value.id}>
                            <td>{value.id}</td>
                            <td>{value.fn}</td>
                            <td>{value.email}</td>
                            <td>{value.login_status}</td>
            </tr>
			    )
			});

		}else{
			// Info = (
				console.log('else');
			// )
		}
      return (
				<table className="table">
                        <thead>
                          <tr>
                            <th>id</th>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Login Status</th>
                          </tr>
                        </thead>
                        <tbody>
                        	{Info}
                        </tbody>
                </table>

      	)
    }
}

export default Show_Users