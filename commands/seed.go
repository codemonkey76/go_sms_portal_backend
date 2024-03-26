package commands

import "sms_portal/routes/api/users"

func SeedCommand(args []string) error {
	users.Seed()

	return nil
}
