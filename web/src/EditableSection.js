import React, { Component } from "react";
import "./styles/EditableSection.css";
import DatePicker from "./DatePicker.js";
import TextField from "./TextField.js";

class EditableSection extends Component {
  render() {
    return (
      <div className="editable-section">
        <DatePicker
          title="Start Date"
          onChange={this.props.onStartDateChange}
          value={this.props.startDate}
        />
        <DatePicker
          title="End Date"
          onChange={this.props.onEndDateChange}
          value={this.props.endDate}
        />
        <TextField
          title="Location"
          onChange={this.props.onLocationChange}
          value={this.props.location}
        />
      </div>
    );
  }
}

export default EditableSection;
