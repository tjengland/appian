import React, { Component } from "react";
import "./styles/TextField.css";

class TextField extends Component {
  render() {
    return (
      <div className="inputWrapper">
        <p>{this.props.title}</p>
        <input
          className="textField"
          type="text"
          onChange={this.props.onChange}
          value={this.props.value}
        />
      </div>
    );
  }
}

export default TextField;
