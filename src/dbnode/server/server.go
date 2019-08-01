	"strings"
	skipRaiseProcessLimitsEnvVar     = "SKIP_PROCESS_LIMITS_RAISE"
	skipRaiseProcessLimitsEnvVarTrue = "true"
	// By default attempt to raise process limits, which is a benign operation.
	skipRaiseLimits := strings.TrimSpace(os.Getenv(skipRaiseProcessLimitsEnvVar))
	if skipRaiseLimits != skipRaiseProcessLimitsEnvVarTrue {
		// Raise fd limits to nr_open system limit
		result, err := xos.RaiseProcessNoFileToNROpen()
		if err != nil {
			logger.Warn("unable to raise rlimit", zap.Error(err))
		} else {
			logger.Info("raised rlimit no file fds limit",
				zap.Bool("required", result.RaisePerformed),
				zap.Uint64("sysNROpenValue", result.NROpenValue),
				zap.Uint64("noFileMaxValue", result.NoFileMaxValue),
				zap.Uint64("noFileCurrValue", result.NoFileCurrValue))
		}
		SetTagDecoderPool(tagDecoderPool).
		SetMaxOutstandingWriteRequests(cfg.Limits.MaxOutstandingWriteRequests).
		SetMaxOutstandingReadRequests(cfg.Limits.MaxOutstandingReadRequests)
	bsGauge := instrument.NewStringListEmitter(scope, "bootstrappers")