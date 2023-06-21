package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"
)

func registerCommonProtocols() {
	_ = secureconnection.NewCommonSecureConnectionProtocol(globals.NEXServer)
}
