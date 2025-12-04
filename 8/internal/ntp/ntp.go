package ntp

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

const ntpServer = "0.beevik-ntp.pool.ntp.org"

func GetCurrentTime() (time.Time, error) {
	currentTime, err := ntp.Time(ntpServer)
	if err != nil {
		return time.Time{}, fmt.Errorf("can't get time via NTP: %w", err)
	}

	return currentTime, nil
}
