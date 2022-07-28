package netatgo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c Client) GetStationData(ctx context.Context, deviceID string, favorites bool) (*http.Response, []byte, error) {
	u, _ := url.Parse("https://api.netatmo.com/api/getstationsdata?get_favorites=false")
	q := u.Query()
	if favorites {
		q.Set("get_favorites", "true")
	}

	if deviceID != "" {
		q.Set("device_id", deviceID)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return resp, nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, body, err
	}

	return resp, body, nil
}
