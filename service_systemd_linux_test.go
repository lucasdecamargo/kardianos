package kardianos

import (
	"bytes"
	"testing"
)

func TestSystemdCustomConfig(t *testing.T) {
	options := make(KeyValue)
	options["Group"] = "glue"
	options["LogOutput"] = true
	options["LogDirectory"] = "/var/log"
	options["PIDFile"] = "/var/run/glue.pid"
	options["Restart"] = "on-success"
	options["RestartSec"] = 2
	// Restart for exit code 2 from Panic, do not restart on Fatal (exit code 1)
	options["SuccessExitStatus"] = "0 2 SIGKILL"
	options["LimitNOFILE"] = -1

	cfg := &Config{
		Name:             "myapp",
		DisplayName:      "My Custom Systemd App",
		Description:      "A custom systemd application",
		WorkingDirectory: "/var/lib/myapp",
		Arguments:        []string{"run"},
		Option:           options,

		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target",
		},
	}

	s := &systemd{
		Config: cfg,
	}

	var to = &struct {
		*Config
		Path                 string
		Group                string
		HasOutputFileSupport bool
		ReloadSignal         string
		PIDFile              string
		LimitNOFILE          int
		Restart              string
		RestartSec           int
		SuccessExitStatus    string
		LogOutput            bool
		LogDirectory         string
	}{
		s.Config,
		"/usr/bin/myapp",
		s.Option.string(optionGroup, ""),
		s.hasOutputFileSupport(),
		s.Option.string(optionReloadSignal, ""),
		s.Option.string(optionPIDFile, ""),
		s.Option.int(optionLimitNOFILE, optionLimitNOFILEDefault),
		s.Option.string(optionRestart, "always"),
		s.Option.int(optionRestartSec, 120),
		s.Option.string(optionSuccessExitStatus, ""),
		s.Option.bool(optionLogOutput, optionLogOutputDefault),
		s.Option.string(optionLogDirectory, defaultLogDirectory),
	}

	buffer := new(bytes.Buffer)
	if err := s.template().Execute(buffer, to); err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	content := buffer.String()
	t.Log(content)
}
