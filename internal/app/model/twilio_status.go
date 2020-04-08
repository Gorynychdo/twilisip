package model

type TwilioStatus struct {
    Status bool   `json:"VerificationStatus"`
    SID    string `json:"OutgoingCallerIdSid"`
}
