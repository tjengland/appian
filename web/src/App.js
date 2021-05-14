import React, { useEffect, useState } from "react";
import { EditableSection, WeatherCard } from "./components";
import "./App.css";
import { WeatherAPI } from "./api";

export default function App() {
  const [weather, setWeather] = useState([]);
  const [dateRange, setDateRange] = useState([]);
  const [town, setTown] = useState(null);

  useEffect(() => {
    async function fetchListings() {
      try {
        const w = await WeatherAPI.getWeather({
          from: getISOWithoutTZ(dateRange[0]),
          to: getISOWithoutTZ(dateRange[1]),
          town,
        });
        setWeather(w.data);
      } catch (err) {}
    }
    fetchListings();
  }, [dateRange, town, setWeather]);

  return (
    <div className="App">
      <EditableSection
        startDate={dateRange[0]}
        endDate={dateRange[1]}
        location={town}
        onStartDateChange={(date) => setDateRange([date, dateRange[1]])}
        onEndDateChange={(date) => setDateRange([dateRange[0], date])}
        onLocationChange={(e) => setTown(e.target.value || null)}
      />
      <div className="editable-section">
        {weather.map((item) => (
          <WeatherCard
            date={item.date}
            weather={item.weather}
            location={item.town}
          />
        ))}
      </div>
    </div>
  );
}

function getISOWithoutTZ(date) {
  return date
    ? new Date(date.getTime() - date.getTimezoneOffset() * 60000).toISOString()
    : null;
}
