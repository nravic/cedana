package types

// XXX: Config file should have a version field to manage future changes to schema
type (
	// Daemon config
	Config struct {
		Options    Options    `key:"options" json:"options" mapstructure:"options"`
		Connection Connection `key:"connection" json:"connection" mapstructure:"connection"`
		Storage    Storage    `key:"storage" json:"storage" mapstructure:"storage"`
		Profiling  Profiling  `key:"profiling" json:"profiling" mapstructure:"profiling"`
		CRIU       CRIU       `key:"criu" json:"criu" mapstructure:"criu"`
		CLI        CLI        `key:"cli" json:"cli" mapstructure:"cli"`
	}
	Options struct {
		Port     uint32 `key:"port" json:"port" mapstructure:"port"`
		Host     string `key:"host" json:"host" mapstructure:"host"`
		UseVSOCK bool   `key:"useVSOCK" json:"use_vsock" mapstructure:"use_vsock"`
	}
	Connection struct {
		// for cedana managed systems
		CedanaUrl       string `key:"cedanaUrl" json:"cedana_url" mapstructure:"cedana_url"`
		CedanaAuthToken string `key:"cedanaAuthToken" json:"cedana_auth_token" mapstructure:"cedana_auth_token"`
	}
	Storage struct {
		Remote      bool   `key:"remote" json:"remote" mapstructure:"remote"`
		DumpDir     string `key:"dumpDir" json:"dump_dir" mapstructure:"dump_dir"`
		Compression string `key:"compression" json:"compression" mapstructure:"compression"`
	}
	Profiling struct {
		Enabled bool `key:"enabled" json:"enabled" mapstructure:"enabled"`
		Otel    struct {
			Port int `key:"port" json:"port" mapstructure:"port"`
		} `key:"otel" json:"otel" mapstructure:"otel"`
	}
	CLI struct {
		WaitForReady bool `key:"waitForReady" json:"wait_for_ready" mapstructure:"wait_for_ready"`
	}
	CRIU struct {
		BinaryPath   string `key:"binaryPath" json:"binary_path" mapstructure:"binary_path"`
		LeaveRunning bool   `key:"leaveRunning" json:"leave_running" mapstructure:"leave_running"`
	}
)
