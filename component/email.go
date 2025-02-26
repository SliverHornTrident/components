//go:build email

package component

import (
	"github.com/SliverHornTrident/components/config"
)

var (
	Email       *gomail.Dialer
	EmailConfig config.Email
)
