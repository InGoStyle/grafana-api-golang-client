package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type GettableGrafanaRule struct {
	Condition       string       `json:"condition,omitempty"`
	Data            []AlertQuery `json:"data,omitempty"`
	ExecErrState    string       `json:"exec_err_state,omitempty"`
	Id              int64        `json:"id,omitempty"`
	IntervalSeconds int64        `json:"intervalSeconds,omitempty"`
	NamespaceId     int64        `json:"namespace_id,omitempty"`
	NamespaceUid    string       `json:"namespace_uid,omitempty"`
	NoDataState     string       `json:"no_data_state,omitempty"`
	OrgId           int64        `json:"orgId,omitempty"`
	RuleGroup       string       `json:"rule_group,omitempty"`
	Title           string       `json:"title,omitempty"`
	Uid             string       `json:"uid,omitempty"`
	Updated         time.Time    `json:"updated,omitempty"`
	Version         int64        `json:"version,omitempty"`
}

type PostableExtendedRuleNode struct {
	Alert        string               `json:"alert,omitempty"`
	Annotations  map[string]string    `json:"annotations,omitempty"`
	Expr         string               `json:"expr,omitempty"`
	For_         *Duration            `json:"for,omitempty"`
	GrafanaAlert *PostableGrafanaRule `json:"grafana_alert,omitempty"`
	Labels       map[string]string    `json:"labels,omitempty"`
	Record       string               `json:"record,omitempty"`
}

type PostableGrafanaRule struct {
	Condition    string       `json:"condition,omitempty"`
	Data         []AlertQuery `json:"data,omitempty"`
	ExecErrState string       `json:"exec_err_state,omitempty"`
	NoDataState  string       `json:"no_data_state,omitempty"`
	Title        string       `json:"title,omitempty"`
	Uid          string       `json:"uid,omitempty"`
}

type AlertQuery struct {
	// Grafana data source unique identifier; it should be '-100' for a Server Side Expression operation.
	DatasourceUid string `json:"datasourceUid,omitempty"`
	// JSON is the raw JSON query and includes the above properties as well as custom properties.
	Model interface{} `json:"model,omitempty"`
	// QueryType is an optional identifier for the type of query. It can be used to distinguish different types of queries.
	QueryType string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId             string             `json:"refId,omitempty"`
	RelativeTimeRange *RelativeTimeRange `json:"relativeTimeRange,omitempty"`
}

type Duration struct {
}

type RelativeTimeRange struct {
	From *Duration `json:"from,omitempty"`
	To   *Duration `json:"to,omitempty"`
}

//
func (c *Client) GetAlertRules() error {

	result := []GettableGrafanaRule{}

	data, err := json.Marshal(result)

	if err != nil {
		return err
	}

	return c.request("GET", "/api/ruler/grafana/api/v1/rules/", nil, bytes.NewBuffer(data), nil)

}

// Create or update alert rule
func (c *Client) AlertRules(Rule PostableExtendedRuleNode, NameSpace string) error {

	data, err := json.Marshal(Rule)
	if err != nil {
		return err
	}

	return c.request("POST", fmt.Sprintf("/api/ruler/grafana/api/v1/rules/%s", NameSpace), nil, bytes.NewBuffer(data), nil)

}
