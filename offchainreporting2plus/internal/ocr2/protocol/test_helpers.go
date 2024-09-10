package protocol

import (
	"github.com/justefg/libocr/commontypes"
	"github.com/justefg/libocr/offchainreporting2plus/types"
)

// Used only for testing
type XXXUnknownMessageType struct{}

// Conform to protocol.Message interface
func (XXXUnknownMessageType) CheckSize(types.ReportingPluginLimits) bool { return true }
func (XXXUnknownMessageType) process(*oracleState, commontypes.OracleID) {}
