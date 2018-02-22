import React from 'react';

class Show_Testplan extends React.Component{
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
          url: Base.Url+"Testplans",
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
                            
                            <td>{value.name}</td>
                            <td>{value.lg_count}</td>
                            <td>{value.team_id}</td>
                            <td>{value.jprofile_id}</td>
                            <td>{value.modified_By}</td>
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
                            <th>Name</th>
                            <th>LG#</th>
                            <th>Team</th>
                            <th>Profile</th>
                            <th>Last Modified</th>
                          </tr>
                        </thead>
                        <tbody>
                        	{Info}
                        </tbody>
                </table>

      	)
    }
}

export default Show_Testplan