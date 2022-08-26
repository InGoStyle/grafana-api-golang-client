package gapi

import (
	"bytes"
	"encoding/json"
)

type GettableGrafanaReceiver struct {
	DisableResolveMessage bool            `json:"disableResolveMessage,omitempty"`
	Name                  string          `json:"name,omitempty"`
	SecureFields          map[string]bool `json:"secureFields,omitempty"`
	Settings              *Json           `json:"settings,omitempty"`
	Type_                 string          `json:"type,omitempty"`
	Uid                   string          `json:"uid,omitempty"`
}

type Json struct {
}

type PostableGrafanaReceivers struct {
	GrafanaManagedReceiverConfigs []PostableGrafanaReceiver `json:"grafana_managed_receiver_configs,omitempty"`
}

type PostableGrafanaReceiver struct {
	DisableResolveMessage bool              `json:"disableResolveMessage,omitempty"`
	Name                  string            `json:"name,omitempty"`
	SecureSettings        map[string]string `json:"secureSettings,omitempty"`
	Settings              Settings          `json:"settings,omitempty"`
	Type_                 string            `json:"type,omitempty"`
	Uid                   string            `json:"uid,omitempty"`
}

type Settings struct {
	ChatId      string `json:"chatid,omitempty"`
	Addresses   string `json:"addresses,omitempty"`
	SingleEmail bool   `json:"singleEmail,omitempty"`
}

type PostableUserConfig struct {
	AlertmanagerConfig PostableApiAlertingConfig `json:"alertmanager_config,omitempty"`
	TemplateFiles      map[string]string         `json:"template_files"`
}

type PostableApiAlertingConfig struct {
	InhibitRules      []InhibitRule      `json:"inhibit_rules,omitempty"`
	MuteTimeIntervals []MuteTimeInterval `json:"mute_time_intervals,omitempty"`
	// Override with our superset receiver type
	Receivers []PostableApiReceiver `json:"receivers,omitempty"`
	Route     Route                 `json:"route,omitempty"`
	Templates []string              `json:"templates"`
}

type InhibitRule struct {
	Equal LabelNames `json:"equal,omitempty"`
	// SourceMatch defines a set of labels that have to equal the given value for source alerts. Deprecated. Remove before v1.0 release.
	SourceMatch    map[string]string `json:"source_match,omitempty"`
	SourceMatchRe  MatchRegexps      `json:"source_match_re,omitempty"`
	SourceMatchers Matchers          `json:"source_matchers,omitempty"`
	// TargetMatch defines a set of labels that have to equal the given value for target alerts. Deprecated. Remove before v1.0 release.
	TargetMatch    map[string]string `json:"target_match,omitempty"`
	TargetMatchRe  MatchRegexps      `json:"target_match_re,omitempty"`
	TargetMatchers Matchers          `json:"target_matchers,omitempty"`
}

type MuteTimeInterval struct {
	Name          string         `json:"name,omitempty"`
	TimeIntervals []TimeInterval `json:"time_intervals,omitempty"`
}

type PostableApiReceiver struct {
	EmailConfigs                  []interface{}             `json:"email_configs,omitempty"`
	GrafanaManagedReceiverConfigs []PostableGrafanaReceiver `json:"grafana_managed_receiver_configs,omitempty"`
	// A unique identifier for this receiver.
	Name string `json:"name,omitempty"`
}

type Route struct {
	Continue_     bool      `json:"continue,omitempty"`
	GroupBy       []string  `json:"group_by,omitempty"`
	GroupInterval *Duration `json:"group_interval,omitempty"`
	GroupWait     *Duration `json:"group_wait,omitempty"`
	// Deprecated. Remove before v1.0 release.
	Match             map[string]string `json:"match,omitempty"`
	MatchRe           *MatchRegexps     `json:"match_re,omitempty"`
	Matchers          *Matchers         `json:"matchers,omitempty"`
	MuteTimeIntervals []string          `json:"mute_time_intervals,omitempty"`
	ObjectMatchers    []interface{}     `json:"object_matchers,omitempty"`
	Provenance        *Provenance       `json:"provenance,omitempty"`
	Receiver          string            `json:"receiver,omitempty"`
	RepeatInterval    *Duration         `json:"repeat_interval,omitempty"`
	Routes            []Route           `json:"routes,omitempty"`
}

type LabelNames struct {
}

type MatchRegexps struct {
}

type Matchers struct {
}

type TimeInterval struct {
	DaysOfMonth []DayOfMonthRange `json:"days_of_month,omitempty"`
	Months      []MonthRange      `json:"months,omitempty"`
	Times       []TimeRange       `json:"times,omitempty"`
	Weekdays    []WeekdayRange    `json:"weekdays,omitempty"`
	Years       []YearRange       `json:"years,omitempty"`
}

type ObjectMatchers struct {
}

type Provenance struct {
}

type DayOfMonthRange struct {
	Begin int64 `json:"Begin,omitempty"`
	End   int64 `json:"End,omitempty"`
}

type MonthRange struct {
	Begin int64 `json:"Begin,omitempty"`
	End   int64 `json:"End,omitempty"`
}

type TimeRange struct {
	EndMinute   int64 `json:"EndMinute,omitempty"`
	StartMinute int64 `json:"StartMinute,omitempty"`
}

type WeekdayRange struct {
	Begin int64 `json:"Begin,omitempty"`
	End   int64 `json:"End,omitempty"`
}

type YearRange struct {
	Begin int64 `json:"Begin,omitempty"`
	End   int64 `json:"End,omitempty"`
}

func (c *Client) CreateContactPoint() ([]AlertNotification, error) {
	alertnotifications := make([]AlertNotification, 0)

	err := c.request("GET", "/api/alert-notifications/", nil, nil, &alertnotifications)
	if err != nil {
		return nil, err
	}

	return alertnotifications, err
}

func (c *Client) GetAlertConfig() (PostableUserConfig, error) {
	alertnotifications := PostableUserConfig{}

	err := c.request("GET", "api/alertmanager/grafana/config/api/v1/alerts", nil, nil, &alertnotifications)
	if err != nil {
		return PostableUserConfig{}, err
	}

	return alertnotifications, err
}

func (c *Client) AlertConfig(Config PostableUserConfig) (Ack, error) {

	data, err := json.Marshal(Config)
	if err != nil {
		return Ack{}, err
	}

	ack := Ack{}

	err = c.request("POST", "api/alertmanager/grafana/config/api/v1/alerts", nil, bytes.NewBuffer(data), &ack)

	if err != nil {
		return Ack{}, err
	}

	return ack, nil
}
