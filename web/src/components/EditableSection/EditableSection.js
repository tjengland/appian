import React from "react";
import "./EditableSection.css";
import DatePicker from "../DatePicker/DatePicker.js";
import TextField from "../TextField/TextField.js";

function EditableSection(props) {
  return (
    <div className="editable-section">
      <DatePicker
        title="Start Date"
        onChange={props.onStartDateChange}
        value={props.startDate}
      />
      <DatePicker
        title="End Date"
        onChange={props.onEndDateChange}
        value={props.endDate}
      />
      <TextField
        title="Location"
        onChange={props.onLocationChange}
        value={props.location}
      />
    </div>
  );
}

export default EditableSection;
