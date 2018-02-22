import React from 'react';
import ReactDOM from 'react-dom';

class Login extends React.Component {
   
   constructor(){
   		
   		super()

   		this.state = {
   			email : '',
   			pwd   : '',
   			emailerr : '',
   			pwderr : '',
   			validemail: false,
   			validpwd: false
   		}
   }

   callhttp(){
        var data = {'email':this.state.email,'pwd':this.state.pwd}
        $.ajax({
          type: "POST",
          url: Base.Url+"Login",
          data: data,
          cache: false,
            success: function(data){
              if(data.status=="S"){
                
              	this.setState({
              		emailerr: '',
              		pwderr: ''
              	})

              	//window.location.href = window.location.href;
              	window.location.reload(true)

              }else if(data.status=="F"){
                
              	this.setState({
              		emailerr: 'err',
              		pwderr: 'err'
              	})
              }
            }.bind(this)
        });
   }

   submit(){

   		if(!this.state.validemail){
   			this.setState({
   				emailerr:'err'
   			})
   		}

   		if (!this.state.validpwd) {
   			this.setState({
   				pwderr:'err'
   			})
   		}

   		if (this.state.validemail && this.state.validpwd) 
   		{
   			this.callhttp()		
   		}
   }

	handlechngEmail(e){

		this.setState({
			email : e.target.value.trim()
		})

		var re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		var e_val =re.test(e.target.value.trim());

		if(e.target.value.trim().length==0 || !e_val){
   			this.setState({
   				emailerr:'err',
   				validemail:false
   			})
   		}

   		if (e.target.value.trim().length>0 && e_val) {
			this.setState({
				emailerr:'',
				validemail:true
			})
   		}

	}

	handlechngPwd(e){

		this.setState({
			pwd : e.target.value
		})

		if(e.target.value.length==0){
			this.setState({
				pwderr : 'err',
				validpwd: false
			})
		}

		if(e.target.value.length>0){
			this.setState({
				pwderr : '',
				validpwd:true
			})
		}

	}

   render() {
      return (
         <div className='container'>
	      <div className='form-signin'>
	        <h3 className='form-signin-heading text-center'>Login</h3>
	        <div className='form-group'>
	          <div className='input-group'>
	            <div className='input-group-addon'>
	              <i className={'glyphicon glyphicon-envelope '+this.state.emailerr}></i>
	            </div>
	            <input 
	            	type='text' 
	            	className='form-control' 
	            	name='email'
	            	id='email'
	            	value={this.state.email}
	            	onChange={this.handlechngEmail.bind(this)} 
	            	placeholder='Email' 
	            	autoComplete='off' />
	          </div>
	        </div>
	        <div className='form-group'>
	          <div className='input-group'>
	            <div className='input-group-addon'>
	              <i className={'glyphicon glyphicon-lock '+ this.state.pwderr}></i>
	            </div>
	            <input 
	            	type='password'
	            	className='form-control' 
	            	name='pwd'
	            	id='pwd'
	            	value={this.state.pwd}
	            	onChange={this.handlechngPwd.bind(this)}
	            	placeholder='Password'
	            	autoComplete='off' />
	          </div>
	        </div>
	        <button className='btn btn-lg btn-primary btn-block' id='signin' type='submit' onClick={this.submit.bind(this)} >Sign in</button>
	        <div className='etc-login-form'>
	          <p className='text-center'>Forgot Your Password? <a href={Base.Url} >Click Here</a></p>
	          <p className='text-center'>New User? <a href={Base.Url+'signup'} >Create New Account</a></p>
	        </div>
	      </div>
         </div>
      );
   }
}

export default Login;