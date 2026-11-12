package androidcheck

import (
	maa "github.com/MaaXYZ/maa-framework-go/v4"
	"github.com/rs/zerolog/log"
)

var _ maa.CustomActionRunner = &SelfCheckAction{}

// SelfCheckAction 供 Android「Agent 自检」调用，验证 Agent 回调链路可用。
type SelfCheckAction struct{}

func (a *SelfCheckAction) Run(_ *maa.Context, _ *maa.CustomActionArg) bool {
	log.Info().Msg("Android agent self-check OK")
	return true
}

func Register() {
	if err := maa.AgentServerRegisterCustomAction("AndroidSelfCheck", &SelfCheckAction{}); err != nil {
		log.Error().Err(err).Msg("failed to register AndroidSelfCheck")
	}
}
