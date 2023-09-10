package errs

import (
	"fmt"
)

var FlagRequired = func(flag string) error {
	return fmt.Errorf("%s flag is required", flag)
}

var FlagRequiredForOtherFlag = func(flag, other string) error {
	return fmt.Errorf("%s flag required when using %s flag", flag, other)
}

var FlagRequiredForCommand = func(flag, command string) error {
	return fmt.Errorf("%s flag required when using %s command", flag, command)
}
