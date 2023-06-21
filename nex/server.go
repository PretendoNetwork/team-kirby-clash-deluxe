package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/team-kirby-clash-deluxe-secure/globals"

	nex "github.com/PretendoNetwork/nex-go"
)

func StartNEXServer() {
	globals.NEXServer = nex.NewServer()
	globals.NEXServer.SetPRUDPVersion(1)
	globals.NEXServer.SetPRUDPProtocolMinorVersion(3)
	globals.NEXServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 10,
		Patch: 26,
	})
	globals.NEXServer.SetKerberosPassword(os.Getenv("KERBEROS_PASSWORD"))
	globals.NEXServer.SetAccessKey("e0c85605")

	globals.NEXServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("== TKCD - Secure ==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("======================")
	})

	// * Register the common handlers first so that they can be overridden if needed
	registerCommonProtocols()
	registerNEXProtocols()

	globals.NEXServer.Listen(":45161")
}
