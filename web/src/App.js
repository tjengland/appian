import React, { useEffect, useState } from "react";
import EditableSection from "./EditableSection";
import WeatherCard from "./WeatherCard";
import "./styles/App.css";
import { WeatherAPI } from "./api";

export default function App() {
  const [weather, setWeather] = useState([]);
  const [dateRange, setDateRange] = useState([]);
  const [town, setTown] = useState(null);

  useEffect(() => {
    async function fetchListings() {
      const from = dateRange[0]
        ? new Date(
            dateRange[0].getTime() - dateRange[0].getTimezoneOffset() * 60000
          ).toISOString()
        : null;
      const to = dateRange[1]
        ? new Date(
            dateRange[1].getTime() - dateRange[1].getTimezoneOffset() * 60000
          ).toISOString()
        : null;
      const w = await WeatherAPI.getWeather({
        from,
        to,
        town,
      });
      setWeather(w.data);
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
