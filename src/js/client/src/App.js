import React, {Component} from 'react';
import { StreamingAPIClient } from "./proto/service_grpc_web_pb";
import { DisplayRequest } from "./proto/service_pb";

class App extends Component  {
  state = {msg: ""}

  componentDidMount(){
    const client = new StreamingAPIClient("http://localhost:8080", null, null);
    const req = new DisplayRequest();
    req.setProcessid = "0000"

    let stream = client.displayStream(req, null);
    stream.on('data', (resp) => {
      console.log(resp.getInfo())
      this.setState({msg:resp.getInfo()})
    })
  }
  
  render(){
    return (
      <div className="App">
        {this.state.msg}
      </div>
    );
  }
  
}

export default App;
