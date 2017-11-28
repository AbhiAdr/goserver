import React from 'react';


class Nav extends React.Component {

  constructor(props){
   	super(props)
  }

  handlSignout(){
        $.ajax({
          type: "POST",
          url: Base.Url+"Logout",
          cache: false,
            success: function(data){
              console.log(data);
              if(data.status=="S"){
              	console.log("success")
              	window.location.href = Base.Url;
              }else if(data.status=="F"){
                console.log("Fail")
              }
            }.bind(this)
        });
  }

   render() {
      return (
		    <nav role="navigation" className="navbar navbar-custom">
		        <div className="container-fluid">

		          <div className="navbar-header">
		            <button data-target="#bs-content-row-navbar-collapse-5" data-toggle="collapse" className="navbar-toggle" type="button">
		              <span className="sr-only">Toggle navigation</span>
		              <span className="icon-bar"></span>
		              <span className="icon-bar"></span>
		              <span className="icon-bar"></span>
		            </button>
		            <a href="#" className="navbar-brand">Go-Web</a>
		          </div>

		          <div id="bs-content-row-navbar-collapse-5" className="collapse navbar-collapse">
		            <ul className="nav navbar-nav navbar-right">		              
		              <li className="dropdown">
		                <a data-toggle="dropdown" className="dropdown-toggle" href="#"> {this.props.UserFirstName} {this.props.UserLastName} <b className="caret"></b></a>
		                <ul role="menu" className="dropdown-menu">
		                  <li className="dropdown-header">Setting</li>
		                  <li><a href="#">Action</a></li>
		                  <li className="divider"></li>
		                  <li className="active"><a href="#">Separated link</a></li>
		                  <li className="divider"></li>
		                  <li className="disabled"><a href="#" onClick={this.handlSignout.bind(this)} >Signout</a></li>
		                </ul>
		              </li>
		            </ul>

		          </div>
		        </div>
		      </nav>
   	  );
   }
}

export default Nav;
