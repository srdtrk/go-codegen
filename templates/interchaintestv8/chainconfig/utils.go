package chainconfig

import "fmt"

func getWasmdChains(chainNum uint8) string {
	var chains string
	for i := uint8(1); i <= chainNum; i++ {
		chains += "\n" + getNewWasmdConfig(fmt.Sprintf("%d", i))
	}
	return chains
}

func getNewWasmdConfig(id string) string {
	return fmt.Sprintf(`	// -- WASMD --
	{
		ChainConfig: ibc.ChainConfig{
			Type:    "cosmos",
			Name:    "wasmd-%s",
			ChainID: "wasmd-%s",
			Images: []ibc.DockerImage{
				{
					Repository: "cosmwasm/wasmd", // FOR LOCAL IMAGE USE: Docker Image Name
					Version:    "v0.50.0",        // FOR LOCAL IMAGE USE: Docker Image Tag
					UidGid:     "1025:1025",
				},
			},
			Bin:            "wasmd",
			Bech32Prefix:   "wasm",
			Denom:          "stake",
			GasPrices:      "0.00stake",
			GasAdjustment:  1.3,
			EncodingConfig: EncodingConfig(),
			TrustingPeriod: "508h",
			NoHostMount:    false,
		},
	},`, id, id)
}
