package utils

import (
	"os"
	"path/filepath"

	pathutils "github.com/xozrc/cloud/pkg"
)

var (
	baseDir = os.Getenv("MACHINE_STORAGE_PATH")

	defaultDeployPath   = ".xocloud"
	defaultMachinesPath = "machines"
	defaultCertPath     = "certs"
)

func GetBaseDir() string {
	if baseDir == "" {
		baseDir = filepath.Join(pathutils.GetHomeDir(), defaultDeployPath)
	}
	return baseDir
}

func GetMachineDir() string {
	return filepath.Join(GetBaseDir(), defaultMachinesPath)
}

func GetMachineCertDir() string {
	return filepath.Join(GetBaseDir(), defaultCertPath)
}
