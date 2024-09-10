package managed

import (
	"context"

	"github.com/justefg/libocr/commontypes"
	"github.com/justefg/libocr/internal/loghelper"
	"github.com/justefg/libocr/offchainreporting2plus/types"
)

func loadConfigFromDatabase(ctx context.Context, database types.ConfigDatabase, logger loghelper.LoggerWithContext) *types.ContractConfig {
	cc, err := database.ReadConfig(ctx)
	if err != nil {
		logger.ErrorIfNotCanceled("loadConfigFromDatabase: Error during Database.ReadConfig", ctx, commontypes.LogFields{
			"error": err,
		})
		return nil
	}

	if cc == nil {
		logger.Info("loadConfigFromDatabase: Database.ReadConfig returned nil, no configuration to restore", nil)
		return nil
	}

	return cc
}
