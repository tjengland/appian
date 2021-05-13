import axios from "axios";

export class WeatherAPI {
  baseUrl = "http://appian.routecanary.com:8081";

  async getWeather(
    params = { town: null, from: null, to: null, limit: null, offset: null }
  ) {
    const res = await axios.get(`${this.baseUrl}/api/weather`, {
      params,
    });

    return res.data;
  }
}

export default new WeatherAPI();
