package btclightclient

// 80 bytes of header
// TODO: Build function to ser and deser BTCHeader to readable format and
// query data from this header
type BTCHeader []byte

// 32 bytes hash of block
type BlockHash []byte

type BTCLightBlock struct {
	// height    uint64
	// header    BTCHeader
	// blockHash BlockHash
	Data string
}

type BTCLightClient struct {
	// TODO: manage BTCLightClient with cosmos storage
	latestBlock BTCLightBlock
}

func newBTCLightClient(lightBlock BTCLightBlock) BTCLightClient {
	return BTCLightClient{
		latestBlock: lightBlock,
	}
}

// TODO: rename this one for more sense in our context
func (btclc BTCLightClient) updateLightBlock(header BTCHeader) {
}
