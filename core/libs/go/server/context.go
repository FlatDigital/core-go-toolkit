package server

// Environment is a string that contains the current environment
// in which the application should boot.
type Environment string

const (
	// EnvProduction is the environment used in production environments
	EnvProduction Environment = "production"

	// EnvStaging is the environment used in staging environments
	EnvStaging Environment = "staging"

	// EnvDevelop is the environment used in development or staging environments
	EnvDevelop Environment = "develop"

	// EnvTest is the environment used in testing environment
	EnvTest Environment = "test"
)

// Role is a string that contains the role in which the application
// should bootstrap.
type Role string

const (

	// RoleRead role will bootstrap the server in read mode.
	// This mode does not provide searching capabilities, but instead provides
	// reading assets by primary id.
	RoleRead Role = "read"

	// RoleWrite role will bootstrap the server in write mode.
	// This mode enables the endpoints needed for writing things to a backing store.
	RoleWrite Role = "write"

	// RoleWorker role will bootstrap the server in worker mode.
	// This mode should be used only by endpoints that receive data from BigQ and do a specific task
	RoleWorker Role = "worker"

	// RoleMiddleEnd role will bootstrap the server in middleend mode.
	// This mode should be used by middle end apps. It allow read & write traffic, it jumps over mlauth validation
	RoleMiddleEnd Role = "middleend"
)

// ApplicationContext contains the necessary information for bootstraping the server
type ApplicationContext struct {
	Environment Environment `json:"environment"`
	Role        Role        `json:"role"`
	Tag         string      `json:"tag"`
}
