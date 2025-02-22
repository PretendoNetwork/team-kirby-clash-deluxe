package nex

import (
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	common_secure.NewCommonProtocol(secureProtocol)
}
