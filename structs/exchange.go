package structs

import "time"

type GetExchangeAnnouncementsResponse struct {
	Announcements []Announcement `json:"announcements"`
}

type AnnouncementStatus string

const (
	AnnouncementStatusInfo    AnnouncementStatus = "info"    // the docs call this AnnouncementTypeInfo
	AnnouncementStatusWarning AnnouncementStatus = "warning" // the docs call this AnnouncementTypeWarning
	AnnouncementStatusError   AnnouncementStatus = "error"   // the docs call this AnnouncementTypeError
	AnnouncementStatusUnknown AnnouncementStatus = ""        // the docs call this AnnouncementTypeUnknown
)

type AnnouncementType string

const (
	AnnouncementTypeInfo    AnnouncementType = "info"
	AnnouncementTypeWarning AnnouncementType = "warning"
	AnnouncementTypeError   AnnouncementType = "error"
	AnnouncementTypeUnknown AnnouncementType = ""
)

type Announcement struct {
	DeliveryTime time.Time          `json:"delivery_time"`
	Message      string             `json:"message"`
	Status       AnnouncementStatus `json:"status"`
	Type         AnnouncementType   `json:"type"`
}

type GetExchangeScheduleResponse struct {
	Schedule ExchangeSchedule `json:"schedule"`
}

type ExchangeSchedule struct {
	MaintenanceWindows []MaintenanceWindow `json:"maintenance_windows"`
	StandardHours      []StandardHours     `json:"standard_hours"`
}

type MaintenanceWindow struct {
	EndDatetime   time.Time `json:"end_datetime"`
	StartDatetime time.Time `json:"start_datetime"`
}

type OneDayHours struct {
	CloseTime string `json:"close_time"`
	OpenTime  string `json:"open_time"`
}

type StandardHours struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	// Arrays are used because the time open is Union([midnight, 3 AM], [8 AM, midnight])
	Monday    []OneDayHours `json:"monday"`
	Tuesday   []OneDayHours `json:"tuesday"`
	Wednesday []OneDayHours `json:"wednesday"`
	Thursday  []OneDayHours `json:"thursday"`
	Friday    []OneDayHours `json:"friday"`
	Saturday  []OneDayHours `json:"saturday"`
	Sunday    []OneDayHours `json:"sunday"`
}

type ExchangeStatus struct {
	ExchangeActive              bool      `json:"exchange_active"`
	ExchangeEstimatedResumeTime time.Time `json:"exchange_estimated_resume_time"`
	TradingActive               bool      `json:"trading_active"`
}
