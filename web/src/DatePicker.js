import React, { Component } from "react";
import DateTimePicker from "react-datetime-picker";
import "./styles/DatePicker.css";

class DatePicker extends Component {
  render() {
    return (
      <div className="inputWrapper">
        <p>{this.props.title}</p>
        <DateTimePicker
          onChange={this.props.onChange}
          value={this.props.value}
        />
      </div>
    );
  }
}

export default DatePicker;
