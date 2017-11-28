import React from 'react';
import ReactDOM from 'react-dom';

class Signup extends React.Component {
   
   constructor(){
   		super()

   		this.state = {
   			fn : '',
   			fnerr:'',
   			validfn:false,
   			ln   : '',
   			lnerr: '',
   			validln:false,
   			email : '',
   			emailerr: '',
   			validemail:false,
   			pwd : '',
   			pwderr:'',
   			validpwd:false
   		}
   }

   callhttp(){

		var data = {'fn':this.state.fn,'ln':this.state.ln,'email':this.state.email,'pwd':this.state.pwd}
        
        $.ajax({
          type: "POST",
          url: Base.Url+"Signup",
          data: data,
          cache: false,
            success: function(data){
              console.log(data);
              if(data.status=="S"){
                
              	this.setState({
              		emailerr: ''
              	})

              	window.location.href = Base.Url;

              }else if(data.status=="F"){
                
              	this.setState({
              		emailerr: 'err'
              	})
              }
            }.bind(this)
        });
   }

   signup(){

   		if(!this.state.validfn){
   			this.setState({
   				fnerr:'err'
   			})
   		}

   		if (!this.state.validln) {
   			this.setState({
   				lnerr:'err'
   			})
   		}

   		if (!this.state.validemail) {
   			this.setState({
   				emailerr:'err'
   			})
   		}

   		if (!this.state.validpwd) {
   			this.setState({
   				pwderr:'err'
   			})
   		}

   		if (this.state.validfn && this.state.validln && this.state.validemail && this.state.validpwd) {
   			this.callhttp()
   		}
   }

   handlechngfn(e){
		
		this.setState({
			fn : e.target.value.trim()
		})

		if(e.target.value.trim().length==0){
			this.setState({
				fnerr : 'err',
				validfn:false
			})
		}else if(e.target.value.trim().length>0){
			this.setState({
				fnerr : '',
				validfn:true
			})
		}
   }

   handlechngln(e){
		
		this.setState({
			ln : e.target.value.trim()
		})

		if(e.target.value.trim().length==0){
			this.setState({
				lnerr : 'err',
				validln:false
			})
		}else if(e.target.value.trim().length>0){
			this.setState({
				lnerr : '',
				validln:true
			})
		}
   }

   handlechngemail(e){
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

   handlechngpwd(e){
		
		this.setState({
			pwd : e.target.value
		})

		if(e.target.value.length==0){
			this.setState({
				pwderr : 'err',
				validpwd:false
			})
		}else if(e.target.value.length>0){
			this.setState({
				pwderr : '',
				validpwd:true
			})
		}
   }

   render() {
      return (

		<div className="container">
			<div className="form-signin">
				<h3 className="form-signin-heading text-center">Signup</h3>

				<div className="form-group">
				<label htmlFor="fn" className={this.state.fnerr} >First Name</label>
				<input 
						className="form-control" 
						name="fn" 
						id="fn" 
						value={this.state.fn}
						placeholder="First Name" 
						onChange={this.handlechngfn.bind(this)}
						autoComplete="off"></input>
				</div>

				<div className="form-group">
				<label htmlFor="ln" className={this.state.lnerr}>Last Name</label>
				<input 
						className="form-control" 
						name="ln" 
						id="ln" 
						value={this.state.ln}
						placeholder="Last Name" 
						onChange={this.handlechngln.bind(this)}
						autoComplete="off"></input>
				</div>

				<div className="form-group">
				<label htmlFor="email" className={this.state.emailerr} >Email</label>
				<input 
						className="form-control" 
						name="email" 
						id="email" 
						value={this.state.email}
						placeholder="Email" 
						onChange={this.handlechngemail.bind(this)}
						autoComplete="off"></input>
				</div>

				<div className="form-group">
				<label htmlFor="pwd" className={this.state.pwderr} >Password</label>
				<input 
						type="password" 
						className="form-control" 
						name="pwd" 
						id="pwd" 
						value={this.state.pwd}
						placeholder="Password" 
						onChange={this.handlechngpwd.bind(this)}
						autoComplete="off"></input>
				</div>

				<div className="form-group">
				<label htmlFor="reg"></label>
				<button className="btn btn-lg btn-primary btn-block" id="reg" type="submit" onClick={this.signup.bind(this)} >Register</button>
				</div>
			</div>
		</div>

      );
   }
}

export default Signup;