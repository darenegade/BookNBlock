package main

import "../../message"

// main starts a whisper node that can be used for transmitting messages via rpc.
// see https://web3js.readthedocs.io/en/1.0/web3-shh.html
func main() {
	c := message.WhisperConfig{
		BootstrapNodes: []string{
			"enode://7d13360f5b1ddcf6947f244639113597a863abba0589d2fa5fffb2816ead0acea6211d5778a8be648e45e81ed881f4c1f5c9bbbf0e79065dfb54bcd97de3beab@127.0.0.1:30349",
		},
		HTTPPort: 8545,
	}
	w := message.StartNode(c)
	w.Node.Wait()
}
