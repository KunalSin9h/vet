package analyzer

import (
	"fmt"

	specmodels "github.com/safedep/vet/gen/models"
	"github.com/safedep/vet/pkg/common/logger"
	"github.com/safedep/vet/pkg/models"
)

const lfpAnalyzerName = "LockfilePoisoningAnalyzer"

type LockfilePoisoningAnalyzerConfig struct {
	FailFast            bool
	TrustedRegistryUrls []string
}

type lockfilePoisoningAnalyzer struct {
	config     LockfilePoisoningAnalyzerConfig
	detections []*AnalyzerEvent
}

type lockfilePoisoningAnalyzerPlugin interface {
	Analyze(manifest *models.PackageManifest, handler AnalyzerEventHandler) error
}

type lockfileAnalyzerPluginBuilder = func(config *LockfilePoisoningAnalyzerConfig) lockfilePoisoningAnalyzerPlugin

var lockfilePoisoningAnalyzers = map[specmodels.Ecosystem]lockfileAnalyzerPluginBuilder{
	specmodels.Ecosystem_Npm: func(config *LockfilePoisoningAnalyzerConfig) lockfilePoisoningAnalyzerPlugin {
		return &npmLockfilePoisoningAnalyzer{
			config: *config,
		}
	},
}

func NewLockfilePoisoningAnalyzer(config LockfilePoisoningAnalyzerConfig) (Analyzer, error) {
	return &lockfilePoisoningAnalyzer{
		config:     config,
		detections: make([]*AnalyzerEvent, 0),
	}, nil
}

func (lfp *lockfilePoisoningAnalyzer) Name() string {
	return lfpAnalyzerName
}

func (lfp *lockfilePoisoningAnalyzer) Analyze(manifest *models.PackageManifest,
	handler AnalyzerEventHandler) error {
	logger.Debugf("LockfilePoisoningAnalyzer: Analyzing [%s] %s",
		manifest.GetSpecEcosystem(), manifest.GetDisplayPath())

	pluginBuilder, ok := lockfilePoisoningAnalyzers[manifest.GetSpecEcosystem()]
	if !ok {
		logger.Warnf("No lockfile poisoning analyzer for ecosystem %s", manifest.Ecosystem)
		return nil
	}

	plugin := pluginBuilder(&lfp.config)
	return plugin.Analyze(manifest, func(event *AnalyzerEvent) error {
		lfp.detections = append(lfp.detections, event)
		if lfp.config.FailFast {
			_ = handler(&AnalyzerEvent{
				Source:  lfpAnalyzerName,
				Type:    ET_AnalyzerFailOnError,
				Message: "Identified lockfile poisoning attempt in " + manifest.GetDisplayPath(),
				Err:     fmt.Errorf("fail-fast on lockfile poisoning at %s", manifest.GetDisplayPath()),
			})
		}

		return handler(event)
	})
}

func (lfp *lockfilePoisoningAnalyzer) Finish() error {
	if len(lfp.detections) > 0 {
		logger.Infof("LockfilePoisoningAnalyzer: Found %d instances", len(lfp.detections))
	}

	return nil
}
