package twiliserver

import (
    "encoding/json"
)

func mustJson(v interface{}) string {
    b, err := json.Marshal(v)
    if err != nil {
        return ""
    }

    return string(b)
}
