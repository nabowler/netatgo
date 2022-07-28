package netatgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	GetStationsDataResponse struct {
		Body       Body    `json:"body"`
		Status     string  `json:"status"`
		TimeExec   float64 `json:"time_exec"`
		TimeServer int64   `json:"time_server"`
	}

	Body struct {
		Devices []Device `json:"devices"`
		User    User     `json:"user"`
	}

	Device struct {
		ID              string              `json:"_id"`
		StationName     string              `json:"station_name"`
		DateSetup       int64               `json:"date_setup"`
		LastSetup       int64               `json:"last_setup"`
		Type            string              `json:"type"`
		LastStatusStore int64               `json:"last_status_store"`
		ModuleName      string              `json:"module_name"`
		Firmware        int64               `json:"firmware"`
		LastUpgrade     int64               `json:"last_upgrade"`
		WifiStatus      int64               `json:"wifi_status"`
		Reachable       bool                `json:"reachable"`
		Co2Calibrating  bool                `json:"co2_calibrating"`
		DataType        []string            `json:"data_type"`
		Place           Place               `json:"place"`
		HomeID          string              `json:"home_id"`
		HomeName        string              `json:"home_name"`
		DashboardData   DeviceDashboardData `json:"dashboard_data"`
		Modules         []Module            `json:"modules"`
	}

	DeviceDashboardData struct {
		TimeUTC          int64   `json:"time_utc"`
		Temperature      float64 `json:"Temperature"`
		Co2              int64   `json:"CO2"`
		Humidity         int64   `json:"Humidity"`
		Noise            int64   `json:"Noise"`
		Pressure         float64 `json:"Pressure"`
		AbsolutePressure float64 `json:"AbsolutePressure"`
		MinTemp          float64 `json:"min_temp"`
		MaxTemp          int64   `json:"max_temp"`
		DateMaxTemp      int64   `json:"date_max_temp"`
		DateMinTemp      int64   `json:"date_min_temp"`
		TempTrend        string  `json:"temp_trend"`
		PressureTrend    string  `json:"pressure_trend"`
	}

	Module struct {
		ID             string              `json:"_id"`
		Type           string              `json:"type"`
		ModuleName     string              `json:"module_name"`
		LastSetup      int64               `json:"last_setup"`
		DataType       []string            `json:"data_type"`
		BatteryPercent int64               `json:"battery_percent"`
		Reachable      bool                `json:"reachable"`
		Firmware       int64               `json:"firmware"`
		LastMessage    int64               `json:"last_message"`
		LastSeen       int64               `json:"last_seen"`
		RFStatus       int64               `json:"rf_status"`
		BatteryVp      int64               `json:"battery_vp"`
		DashboardData  ModuleDashboardData `json:"dashboard_data"`
	}

	ModuleDashboardData struct {
		TimeUTC        int64    `json:"time_utc"`
		Temperature    *float64 `json:"Temperature,omitempty"`
		Humidity       *int64   `json:"Humidity,omitempty"`
		MinTemp        *float64 `json:"min_temp,omitempty"`
		MaxTemp        *float64 `json:"max_temp,omitempty"`
		DateMaxTemp    *int64   `json:"date_max_temp,omitempty"`
		DateMinTemp    *int64   `json:"date_min_temp,omitempty"`
		TempTrend      *string  `json:"temp_trend,omitempty"`
		Rain           *int64   `json:"Rain,omitempty"`
		SumRain1       *int64   `json:"sum_rain_1,omitempty"`
		SumRain24      *float64 `json:"sum_rain_24,omitempty"`
		Co2            *int64   `json:"CO2,omitempty"`
		WindStrength   *int64   `json:"WindStrength,omitempty"`
		WindAngle      *int64   `json:"WindAngle,omitempty"`
		GustStrength   *int64   `json:"GustStrength,omitempty"`
		GustAngle      *int64   `json:"GustAngle,omitempty"`
		MaxWindStr     *int64   `json:"max_wind_str,omitempty"`
		MaxWindAngle   *int64   `json:"max_wind_angle,omitempty"`
		DateMaxWindStr *int64   `json:"date_max_wind_str,omitempty"`
	}

	Place struct {
		Altitude int64     `json:"altitude"`
		City     string    `json:"city"`
		Country  string    `json:"country"`
		Timezone string    `json:"timezone"`
		Location []float64 `json:"location"`
	}

	User struct {
		Mail           string         `json:"mail"`
		Administrative Administrative `json:"administrative"`
	}

	Administrative struct {
		Lang         string `json:"lang"`
		RegLocale    string `json:"reg_locale"`
		Country      string `json:"country"`
		Unit         int64  `json:"unit"`
		Windunit     int64  `json:"windunit"`
		Pressureunit int64  `json:"pressureunit"`
		FeelLikeAlgo int64  `json:"feel_like_algo"`
	}

	ErrorResponse struct {
		Err APIError `json:"error"`
	}

	APIError struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	}
)

func (c Client) GetStationData(ctx context.Context, deviceID string, favorites bool) (GetStationsDataResponse, error) {
	u, _ := url.Parse("https://api.netatmo.com/api/getstationsdata?get_favorites=false")
	q := u.Query()
	if favorites {
		q.Set("get_favorites", "true")
	}

	if deviceID != "" {
		q.Set("device_id", deviceID)
	}

	u.RawQuery = q.Encode()

	var dataResp GetStationsDataResponse

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return dataResp, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return dataResp, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dataResp, err
	}

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return dataResp, err
		}
		return dataResp, errResp
	}

	return dataResp, json.Unmarshal(body, &dataResp)
}

func (err ErrorResponse) Error() string {
	return err.Err.Error()
}

func (err APIError) Error() string {
	return fmt.Sprintf("%d %s", err.Code, err.Message)
}
