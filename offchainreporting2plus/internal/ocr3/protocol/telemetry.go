package protocol

import (
	"github.com/justefg/libocr/commontypes"
	"github.com/justefg/libocr/offchainreporting2plus/types"
)

type TelemetrySender interface {
	RoundStarted(
		configDigest types.ConfigDigest,
		epoch uint64,
		seqNr uint64,
		round uint64,
		leader commontypes.OracleID,
	)
}
