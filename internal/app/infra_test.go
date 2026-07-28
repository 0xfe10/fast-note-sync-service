package app

import (
	"path/filepath"
	"testing"

	"github.com/haierkeys/fast-note-sync-service/internal/config"
	"github.com/stretchr/testify/require"
)

func TestMachineUUIDPathUsesSQLiteDatabaseDirectory(t *testing.T) {
	cfg := &AppConfig{
		File: "config/config.yaml",
		Database: config.DatabaseConfig{
			Type: "sqlite",
			Path: "storage/database/db.sqlite3",
		},
	}

	require.Equal(t, filepath.Join("storage", "database", ".server_uuid"), machineUUIDPath(cfg))
}

func TestMachineUUIDPathFallsBackToConfigDirectory(t *testing.T) {
	for _, database := range []config.DatabaseConfig{
		{Type: "postgres"},
		{Type: "sqlite", Path: ":memory:"},
	} {
		cfg := &AppConfig{File: "config/config.yaml", Database: database}
		require.Equal(t, filepath.Join("config", ".server_uuid"), machineUUIDPath(cfg))
	}
}
