package gateways

type GatewayContainer struct {
	Ledger LedgerGateway
}

func InitGateway() GatewayContainer {
	return GatewayContainer{
		Ledger: NewLedgerGateway(),
	}
}
