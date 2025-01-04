package kalshigo

import (
	"encoding/json"

	"github.com/arvindh-manian/kalshigo/structs"
)

const ANNOUNCEMENT_ENDPOINT = "/trade-api/v2/exchange/announcements"
const SCHEDULE_ENDPOINT = "/trade-api/v2/exchange/schedule"
const STATUS_ENDPOINT = "/trade-api/v2/exchange/status"

func (c *Client) GetExchangeAnnouncements() ([]structs.Announcement, error) {
	body, _, err := c.getRequest(ANNOUNCEMENT_ENDPOINT, nil)

	if err != nil {
		return []structs.Announcement{}, err
	}

	var announcements structs.GetExchangeAnnouncementsResponse

	err = json.Unmarshal(body, &announcements)

	if err != nil {
		return []structs.Announcement{}, err
	}

	return announcements.Announcements, nil
}

func (c *Client) GetExchangeSchedule() (*structs.GetExchangeScheduleResponse, error) {
	body, _, err := c.getRequest(SCHEDULE_ENDPOINT, nil)

	if err != nil {
		return &structs.GetExchangeScheduleResponse{}, err
	}

	var schedule structs.GetExchangeScheduleResponse

	err = json.Unmarshal(body, &schedule)

	if err != nil {
		return &structs.GetExchangeScheduleResponse{}, err
	}

	return &schedule, nil
}

func (c *Client) GetExchangeStatus() (*structs.ExchangeStatus, error) {
	body, _, err := c.getRequest(STATUS_ENDPOINT, nil)

	if err != nil {
		return &structs.ExchangeStatus{}, err
	}

	var status structs.ExchangeStatus

	err = json.Unmarshal(body, &status)

	if err != nil {
		return &structs.ExchangeStatus{}, err
	}

	return &status, nil
}
