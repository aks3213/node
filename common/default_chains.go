package common

func EthChain() Chain {
	return Chain{
		ChainName: ChainName_Eth,
		ChainId:   1,
	}
}

func GoerliChain() Chain {
	return Chain{
		ChainName: ChainName_Goerli,
		ChainId:   1337,
	}
}

func RopstenChain() Chain {
	return Chain{
		ChainName: ChainName_Ropsten,
		ChainId:   3,
	}
}

func BscMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_BscMainnet,
		ChainId:   56,
	}
}

func BscTestnetChain() Chain {
	return Chain{
		ChainName: ChainName_BscTestnet,
		ChainId:   97,
	}
}

func BaobabChain() Chain {
	return Chain{
		ChainName: ChainName_Baobab,
		ChainId:   1001,
	}
}
func ZetaChain() Chain {
	return Chain{
		ChainName: ChainName_ZetaChain,
		ChainId:   101,
	}
}
func BtcMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_Btc,
		ChainId:   55555,
	}
}

func BtcTestNetChain() Chain {
	return Chain{
		ChainName: ChainName_BtcTestNet,
		ChainId:   80001,
	}
}

func PolygonChain() Chain {
	return Chain{
		ChainName: ChainName_Polygon,
		ChainId:   137,
	}
}

func MumbaiChain() Chain {
	return Chain{
		ChainName: ChainName_Mumbai,
		ChainId:   80001,
	}
}

func DefaultChainsList() []*Chain {
	chains := []Chain{
		BtcTestNetChain(),
		BtcMainnetChain(),
		MumbaiChain(),
		PolygonChain(),
		BaobabChain(),
		BscTestnetChain(),
		BscMainnetChain(),
		GoerliChain(),
		RopstenChain(),
		EthChain(),
		ZetaChain(),
	}
	var c []*Chain
	for i := 0; i < len(chains); i++ {
		c = append(c, &chains[i])
	}
	return c
}