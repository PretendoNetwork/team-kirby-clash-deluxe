package nex

import (
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/ticket-granting"
	"github.com/PretendoNetwork/team-kirby-clash-deluxe/globals"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticket_granting.NewCommonTicketGrantingProtocol(globals.AuthenticationServer)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.SetScheme("prudps")
	secureStationURL.SetAddress(os.Getenv("PN_TKCD_SECURE_SERVER_HOST"))
	secureStationURL.SetPort(os.Getenv("PN_TKCD_SECURE_SERVER_PORT"))
	secureStationURL.SetCID("1")
	secureStationURL.SetPID("2")
	secureStationURL.SetSID("1")
	secureStationURL.SetStream("10")
	secureStationURL.SetType("2")

	ticketGrantingProtocol.SetSecureStationURL(secureStationURL)
	ticketGrantingProtocol.SetBuildName(serverBuildString)

	globals.AuthenticationServer.SetPasswordFromPIDFunction(globals.PasswordFromPID)
}
