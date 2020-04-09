package model

import (
    "encoding/json"
)

type TwilioCallerId struct {
    Status  int    `json:"status,omitempty"`
    Message string `json:"message,omitempty"`
    Code    string `json:"validation_code"`
    SID     string `json:"call_sid"`
}

func NewTwilioCallerId(b []byte) (*TwilioCallerId, error) {
    m := &TwilioCallerId{
        Status: 200,
    }

    if err := json.Unmarshal(b, m); err != nil {
        return nil, err
    }

    return m, nil
}

func (m *TwilioCallerId) Dump() string {
    m.Status = 0
    m.Message = ""

    b, _ := json.Marshal(m)
    return string(b)
}
