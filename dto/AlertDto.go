package dto

import (
	"encoding/json"
	"time"
)

type AlertPayload struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Status      string      `json:"status"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
	StartsAt    time.Time   `json:"startsAt"`
	EndsAt      time.Time   `json:"endsAt"`
	StartTime   string
	EndTime     string
	Count       int
	Fingerprint string `json:"fingerprint"`
}

type Labels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Serverity string `json:"serverity"` // 注意：原始 JSON 中为 "serverity"，可能是拼写错误（应为 "severity"）
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func (c *Alert) UnmarshalJSON(data []byte) error {
	var temp struct {
		Status      string      `json:"status"`
		Labels      Labels      `json:"labels"`
		Annotations Annotations `json:"annotations"`
		StartsAt    time.Time   `json:"startsAt"`
		EndsAt      time.Time   `json:"endsAt"`
		StartTime   string
		EndTime     string
		Count       int
		Fingerprint string `json:"fingerprint"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	c.Status = temp.Status
	c.Labels = temp.Labels
	c.Annotations = temp.Annotations
	c.StartsAt = temp.StartsAt
	c.EndsAt = temp.EndsAt
	c.StartTime = c.StartsAt.Format("2006-01-02 15:04:05")
	c.EndTime = c.EndsAt.Format("2006-01-02 15:04:05")
	c.Count = temp.Count
	c.Fingerprint = temp.Fingerprint
	return nil
}
