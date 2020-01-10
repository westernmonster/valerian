package model

const (
	//ConfigIng config ing.
	ConfigStateInProgress = int32(1)
	//ConfigEnd config end.
	ConfigStateEnd = int32(2)
)

const (
	DefaultEnv    = "dev"
	DefaultZone   = "hz001"
	DefaultRegion = "hz"

	//StatusShow status show
	AppStatusShow = int32(1)
	//StatusHidden status hidden
	AppStatusHidden = int32(2)
)

const (
	EnvUat  = "uat"
	EnvDev  = "dev"
	EnvProd = "prod"
)

const (
	PlatformAdmin     = int32(1)
	PlatformInfra     = int32(2)
	PlatformService   = int32(3)
	PlatformInterface = int32(4)
)
