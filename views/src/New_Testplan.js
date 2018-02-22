import React from 'react';

class New_Testplan extends React.Component{


	render() {
  		return (
         <div className="panel-body">
              <div className="form-group">
                <label className="col-md-2 control-label">Name</label>
                <div className="col-md-10">
                  <input required="" placeholder="Title" id="title" className="form-control" name="title" type="text"></input>
                </div>
              </div>
              <div className="form-group">
                <label className="col-md-2 control-label">Team</label>
                <div className="col-md-10">
                  
                  <div className="selecter  open" tabIndex="0">
                  <select name="selecter_basic" className="selecter_2 selecter-element" data-selecter-options="" tabIndex="-1">
                    <option value="#">Google Search</option>
                    <option value="#">BoingBoing</option>
                    <option value="#">CNN News</option>
                  </select>
                  <span className="selecter-selected placeholder">Google Search</span>
                  <div className="selecter-options" styleName={{display : "block"}} >
                    <a className="selecter-item selected" href="#">Google Search</a>
                    <a className="selecter-item" href="#">BoingBoing</a>
                    <a className="selecter-item" href="#">CNN News</a>
                  </div>
                </div>


                </div>
              </div>
              <div className="form-group">
                <label className="col-md-2 control-label" >Description</label>
                <div className="col-md-10">
                  <textarea required="" className="form-control" placeholder="Description" rows="10" cols="30" id="description" name="description"></textarea>
                </div>
              </div>
              <div className="form-group">
                <div className="col-md-offset-2 col-md-10">
                  <button className="btn btn-info" type="submit">Submit</button>
                </div>
              </div>
          </div>
      );
  }
}

export default New_Testplan