package utils

import (
	"fmt"

	"github.com/create-go-app/fiber-go-template/pkg/const"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case _const.AdminRoleName:
		// Nothing to do, verified successfully.
	case _const.ModeratorRoleName:
		// Nothing to do, verified successfully.
	case _const.UserRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
