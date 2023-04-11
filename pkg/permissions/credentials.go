package permissions

import "fmt"

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role string) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case AdminRole:
		// Admin credentials (all access).
		credentials = []string{
			BookCreateCredential,
			BookUpdateCredential,
			BookDeleteCredential,
		}
	case ModeratorRole:
		// Moderator credentials (only book creation and update).
		credentials = []string{
			BookCreateCredential,
			BookUpdateCredential,
		}
	case UserRole:
		// Simple user credentials (only book creation).
		credentials = []string{
			BookCreateCredential,
		}
	case ReadOnlyRole:
		// Simple read only credentials (only read books).
		credentials = []string{
			BookReadCredential,
		}
	default:
		// Return error message.
		return []string{}, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
