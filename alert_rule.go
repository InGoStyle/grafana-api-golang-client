package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type GettableRuleGroupConfig struct {
	Name     string                     `json:"name,omitempty"`
	Interval string                     `json:"interval,omitempty"`
	Rules    []GettableExtendedRuleNode `json:"rules,omitempty"`
}

type GettableExtendedRuleNode struct {
	Alert        string              `json:"alert,omitempty"`
	Annotations  map[string]string   `json:"annotations,omitempty"`
	Expr         string              `json:"expr,omitempty"`
	For_         string              `json:"for,omitempty"`
	GrafanaAlert GettableGrafanaRule `json:"grafana_alert,omitempty"`
	Labels       map[string]string   `json:"labels,omitempty"`
	Record       string              `json:"record,omitempty"`
}

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
	Updated         *time.Time   `json:"updated,omitempty"`
	Version         int64        `json:"version,omitempty"`
}

type PostableRuleGroupConfig struct {
	Interval string                     `json:"interval,omitempty"`
	Name     string                     `json:"name,omitempty"`
	Rules    []PostableExtendedRuleNode `json:"rules,omitempty"`
}

type PostableExtendedRuleNode struct {
	Alert        string              `json:"alert,omitempty"`
	Annotations  map[string]string   `json:"annotations,omitempty"`
	Expr         string              `json:"expr,omitempty"`
	For_         string              `json:"for,omitempty"`
	GrafanaAlert PostableGrafanaRule `json:"grafana_alert,omitempty"`
	Labels       map[string]string   `json:"labels,omitempty"`
	Record       string              `json:"record,omitempty"`
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
	Model Model `json:"model,omitempty"`
	// QueryType is an optional identifier for the type of query. It can be used to distinguish different types of queries.
	QueryType string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId             string             `json:"refId,omitempty"`
	RelativeTimeRange *RelativeTimeRange `json:"relativeTimeRange,omitempty"`
}

type Evaluator struct {
	Params []int64 `json:"params,omitempty"`
	Type   string  `json:"type,omitempty"`
}

type Operator struct {
	Type string `json:"type,omitempty"`
}

type Query struct {
	Params []string `json:"params,omitempty"`
}

type Reducer struct {
	Params []string `json:"params,omitempty"`
	Type   string   `json:"type,omitempty"`
}

type Conditions struct {
	Evaluator Evaluator `json:"evaluator,omitempty"`
	Operator  Operator  `json:"operator,omitempty"`
	Query     Query     `json:"query,omitempty"`
	Reducer   Reducer   `json:"reducer,omitempty"`
	Type      string    `json:"type,omitempty"`
}

type Datasource struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
}

type Model struct {
	Exemplar      bool         `json:"exemplar,omitempty"`
	Expr          string       `json:"expr"`
	Hide          bool         `json:"hide,omitempty"`
	Interval      string       `json:"interval,omitempty"`
	IntervalMs    int64        `json:"intervalMs,omitempty"`
	LegendFormat  string       `json:"legendFormat,omitempty"`
	MaxDataPoints int64        `json:"maxDataPoints,omitempty"`
	RefId         string       `json:"refId,omitempty"`
	Conditions    []Conditions `json:"conditions,omitempty"`
	Datasource    Datasource   `json:"datasource,omitempty"`
	Type          string       `json:"type,omitempty"`
}

type Duration struct {
}

type Ack struct {
	Message string `json:"message"`
}

type RelativeTimeRange struct {
	From int64 `json:"from,omitempty"`
	To   int64 `json:"to,omitempty"`
}

//1
func (c *Client) GetAlertRules(NameSpace string) ([]GettableRuleGroupConfig, error) {

	result := map[string][]GettableRuleGroupConfig{}
	err := c.request("GET", fmt.Sprintf("/api/ruler/grafana/api/v1/rules/%s", NameSpace), nil, nil, &result)

	if err != nil {
		return nil, err
	}

	return result[NameSpace], nil
}

// Create or update alert rule
func (c *Client) AlertRules(Rule PostableRuleGroupConfig, NameSpace string) (Ack, error) {

	data, err := json.Marshal(Rule)
	if err != nil {
		return Ack{}, err
	}

	ack := Ack{}

	err = c.request("POST", fmt.Sprintf("/api/ruler/grafana/api/v1/rules/%s", NameSpace), nil, bytes.NewBuffer(data), &ack)

	if err != nil {
		return Ack{}, err
	}

	return ack, nil
}

func (c *Client) DeleteAlertGroup(NameSpace, GroupName string) (Ack, error) {

	ack := Ack{}

	err := c.request("DELETE", fmt.Sprintf("/api/ruler/grafana/api/v1/rules/%s/%s", NameSpace, GroupName), nil, nil, &ack)

	if err != nil {
		return Ack{}, err
	}
	return ack, nil
}
