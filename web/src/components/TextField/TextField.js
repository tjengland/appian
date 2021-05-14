import React from "react";
import "./TextField.css";

function TextField(props) {
  return (
    <div className="inputWrapper">
      <p>{props.title}</p>
      <input
        className="textField"
        type="text"
        onChange={props.onChange}
        value={props.value}
      />
    </div>
  );
}

export default TextField;
