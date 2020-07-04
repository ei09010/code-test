package hash_service

import "fmt"

// fletcher32 simplified implementation
func Generate(websiteUrl string) string {

	webSiteBytes := []byte(websiteUrl)

	sum1 := uint32(0)
	sum2 := uint32(0)

	for i := 0; i < len(webSiteBytes); i++ {
		sum1 = (sum1 + uint32(webSiteBytes[i])) % 0xffff
		sum2 = (sum2 + sum1) % 0xffff
	}

	finalSum := (sum2<<16 | sum1)
	// returns 32 bit hash in hexadecimal format
	return fmt.Sprintf("%x", finalSum)

}
