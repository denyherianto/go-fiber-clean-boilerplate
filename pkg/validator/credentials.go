package utils

// "fmt"

// "github.com/denyherianto/go-fiber-clean-boilerplate/app/constants"

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role []string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// // Switch given role.
	// switch role {
	// case constants.AdminRoleName:
	// 	// Admin credentials (all access).
	// 	credentials = []string{
	// 		// constants.BookCreateCredential,
	// 		// constants.BookUpdateCredential,
	// 		// constants.BookDeleteCredential,
	// 	}
	// case constants.ModeratorRoleName:
	// 	// Moderator credentials (only book creation and update).
	// 	credentials = []string{
	// 		constants.BookCreateCredential,
	// 		constants.BookUpdateCredential,
	// 	}
	// case constants.UserRoleName:
	// 	// Simple user credentials (only book creation).
	// 	credentials = []string{
	// 		constants.BookCreateCredential,
	// 	}
	// default:
	// 	// Return error message.
	// 	return nil, fmt.Errorf("role '%v' does not exist", role)
	// }

	return credentials, nil
}
