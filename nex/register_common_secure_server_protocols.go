package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
)

func registerCommonSecureServerProtocols() {
	_ = secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
}
