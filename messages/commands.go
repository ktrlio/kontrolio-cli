package messages

import "github.com/kyokomi/emoji/v2"

const (
	SYNC_OFFLINE         = "You need to be connected to the internet in order to sync your offline and online data."
	SYNC_SERVICE_DOWN    = "Sorry. We can't sync your offline and online data right now because the service is unavailable."
	SYNC_CONFIG_MISSING  = "You need to have a configuration file set in order to sync your offline and online data."
	SYNC_API_KEY_MISSING = "You need to have an API key set on your configuration file in order to sync your offline and online data. Sign up @ https://kontrolio.com to get an API key."
	WORKDAY_STATUS       = "\nToday, you've worked: "
	PUNCH_SUCCESS        = " sucessfully."
)

var (
	YOURE_ONLINE  = emoji.Sprint("\n:earth_americas:You're online.\n")
	YOURE_OFFLINE = emoji.Sprint("\n:mobile_phone_off:You're offline.\n")
)
