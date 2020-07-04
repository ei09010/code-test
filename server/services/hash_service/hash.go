package hash_service

// returns hash generated based on XOR with a pre-defined base value
func Generate(websiteUrl string, baseValue string) string {

	var xorbytes []byte

	xorbytes = make([]byte, len(websiteUrl))

	var i int
	for i = 0; i < len(websiteUrl); i++ {
		xorbytes[i] = websiteUrl[i] ^ baseValue[i]
	}

	return string(xorbytes)
}
