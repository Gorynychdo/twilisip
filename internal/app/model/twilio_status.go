package model

import (
    "encoding/json"
)

type TwilioStatus struct {
    Status string `json:"status"`
    To     string `json:"to"`
    Date   string `json:"date"`
    SID    string `json:"sid"`
}

func (ts *TwilioStatus) Dump() string {
    b, _ := json.Marshal(ts)
    return string(b)
}
