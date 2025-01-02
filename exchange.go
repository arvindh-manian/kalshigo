package kalshigo

import (
	"encoding/json"
)

const ANNOUNCEMENT_ENDPOINT = "/trade-api/v2/exchange/announcements"
const SCHEDULE_ENDPOINT = "/trade-api/v2/exchange/schedule"
const STATUS_ENDPOINT = "/trade-api/v2/exchange/status"

func (c *Client) GetExchangeAnnouncements() ([]Announcement, error) {
	body, _, err := c.getRequest(ANNOUNCEMENT_ENDPOINT, nil)

	if err != nil {
		return []Announcement{}, err
	}

	var announcements GetExchangeAnnouncementsResponse

	err = json.Unmarshal(body, &announcements)

	if err != nil {
		return []Announcement{}, err
	}

	return announcements.Announcements, nil
}

func (c *Client) GetExchangeSchedule() (*GetExchangeScheduleResponse, error) {
	body, _, err := c.getRequest(SCHEDULE_ENDPOINT, nil)

	if err != nil {
		return &GetExchangeScheduleResponse{}, err
	}

	var schedule GetExchangeScheduleResponse

	err = json.Unmarshal(body, &schedule)

	if err != nil {
		return &GetExchangeScheduleResponse{}, err
	}

	return &schedule, nil
}

func (c *Client) GetExchangeStatus() (*ExchangeStatus, error) {
	body, _, err := c.getRequest(STATUS_ENDPOINT, nil)

	if err != nil {
		return &ExchangeStatus{}, err
	}

	var status ExchangeStatus

	err = json.Unmarshal(body, &status)

	if err != nil {
		return &ExchangeStatus{}, err
	}

	return &status, nil
}
