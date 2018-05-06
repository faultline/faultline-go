# faultline-go [![Build Status](https://travis-ci.org/faultline/faultline-go.svg?branch=master)](https://travis-ci.org/faultline/faultline-go)
![GitHub release](https://img.shields.io/github/release/faultline/faultline-go.svg)

> [faultline](https://github.com/faultline/faultline) exception and error notifier for Go.

## Requirement

- [faultline](https://github.com/faultline/faultline) >= 1.3.0

## Usage

``` go
package main

import (
	"errors"
	"github.com/faultline/faultline-go/faultline"
)

var notifications = []interface{}{
	faultline.Slack{
		Type:           "slack",
		Endpoint:       "https://hooks.slack.com/services/XXXXXXXXXX/BAC0D0N69/NacHbWgIfklAHH7XBEItGNcs",
		Channel:        "#random",
		Username:       "faultline-notify",
		NotifyInterval: 5,
		Threshold:      10,
		Timezone:       "Asia/Tokyo",
	},
}

var notifier = faultline.NewNotifier("faultline-go-project", "xxxxXXXXXxXxXXxxXXXXXXXxxxxXXXXXX", "https://xxxxxxxxx.execute-api.ap-northeast-1.amazonaws.com/v0", notifications)

func main() {
	defer notifier.Close()
	defer notifier.NotifyOnPanic()

	notifier.Notify(errors.New("operation failed"), nil)
}
```

## References

- faultline-go is based on [airbrake/gobrake](https://github.com/airbrake/gobrake)
    - Gobrake is licensed under [BSD 3-Clause](https://github.com/airbrake/gobrake/LICENSE).
