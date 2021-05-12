package api

type Weather struct {
	Town    string `json:"town"`
	Weather string `json:"weather"`
	Date    string `json:"date"`
}

var WeatherData = []Weather{
	{
		Date:    "01/01/2020 12:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 12:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 12:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 12:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 01:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 01:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 01:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 01:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 02:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 02:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 02:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 02:00 AM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 03:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 03:00 AM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 03:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 03:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 AM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 AM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 05:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 05:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 05:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 05:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 06:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 06:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 06:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 06:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 07:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 07:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 07:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 07:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 08:00 AM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 AM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 09:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 09:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 09:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 09:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 10:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 10:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 10:00 AM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 10:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 11:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 11:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 11:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 11:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 12:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 12:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 12:00 PM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 12:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 01:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 01:00 PM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 01:00 PM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 01:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 02:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 02:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 02:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 02:00 PM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 03:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 03:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 03:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 03:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 04:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 04:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 04:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 05:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 05:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 05:00 PM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 05:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 06:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 06:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 06:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 06:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 07:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 07:00 PM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 07:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 07:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 08:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 08:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 09:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 09:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 09:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 09:00 PM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 10:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/01/2020 10:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 10:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 10:00 PM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 11:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/01/2020 11:00 PM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/01/2020 11:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/01/2020 11:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 12:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 12:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 12:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 12:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 01:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 AM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 02:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 02:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 02:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 02:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 03:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 03:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 03:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 03:00 AM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 04:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 04:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 04:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 04:00 AM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 05:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 06:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 AM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 AM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 07:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 07:00 AM",
		Town:    "Reston",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 07:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 07:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 08:00 AM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 08:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 AM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 AM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 09:00 AM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 10:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 10:00 AM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 10:00 AM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 10:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 11:00 AM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 11:00 AM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 11:00 AM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 11:00 AM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 12:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 12:00 PM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 12:00 PM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 12:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 01:00 PM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 01:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 02:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 02:00 PM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 02:00 PM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 02:00 PM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 03:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 03:00 PM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 03:00 PM",
		Town:    "Arlington",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 03:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 04:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 04:00 PM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 04:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 04:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 05:00 PM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 05:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 06:00 PM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 06:00 PM",
		Town:    "Fall's Church",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 07:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 07:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 07:00 PM",
		Town:    "Arlington",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 07:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 08:00 PM",
		Town:    "Fall's Church",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 PM",
		Town:    "Reston",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 09:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 09:00 PM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 10:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 10:00 PM",
		Town:    "Reston",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 10:00 PM",
		Town:    "Arlington",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 10:00 PM",
		Town:    "Fall's Church",
		Weather: "Rainy",
	},
	{
		Date:    "01/02/2020 11:00 PM",
		Town:    "Tyson's Corner",
		Weather: "Partly Cloudy",
	},
	{
		Date:    "01/02/2020 11:00 PM",
		Town:    "Reston",
		Weather: "Sunny",
	},
	{
		Date:    "01/02/2020 11:00 PM",
		Town:    "Arlington",
		Weather: "Cloudy",
	},
	{
		Date:    "01/02/2020 11:00 PM",
		Town:    "Fall's Church",
		Weather: "Partly Cloudy",
	},
}
