import React from "react";
import DateTimePicker from "react-datetime-picker";
import "./DatePicker.css";

function DatePicker(props) {
  return (
    <div className="inputWrapper">
      <p>{props.title}</p>
      <DateTimePicker onChange={props.onChange} value={props.value} />
    </div>
  );
}

export default DatePicker;
