import React, { Component } from "react";
import EditableSection from "./EditableSection";
import WeatherCard from "./WeatherCard";
import "./styles/App.css";

class App extends Component {
  render() {
    return (
      <div className="App">
        <EditableSection
          startDate={null}
          endDate={null}
          location={null}
          onStartDateChange={() => {}}
          onEndDateChange={() => {}}
          onLocationChange={() => {}}
        />
        <div className="editable-section">
          {[
            {
              date: "2020-01-01 10:00AM",
              weather: "Cloudy",
              location: "Tyson's Corner",
            },
          ].map((item) => (
            <WeatherCard
              date={item.date}
              weather={item.weather}
              location={item.location}
            />
          ))}
        </div>
        <text>
          TODO: Implement the visual application using the provided components
          And by displaying data from test-data.csv or test-data.json
        </text>
      </div>
    );
  }
}

export default App;
