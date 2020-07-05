package hash_service

type Hash interface {
	Generate(websiteUrl string, baseValue string) interface{}
}
