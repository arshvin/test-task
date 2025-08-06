package handlers

const (
	dateNow int = iota * 2
	timeNow
	cpuAmount
	addrIpV4Amount

	even int = iota << 2
	odd
)

var (
	entityTypePretty map[int]string
	parityPretty     map[int]string
)

func init() {
	entityTypePretty = make(map[int]string)
	entityTypePretty[dateNow] = "DATE_NOW"
	entityTypePretty[timeNow] = "TIME_NOW"
	entityTypePretty[cpuAmount] = "CPU_AMOUNT"
	entityTypePretty[addrIpV4Amount] = "IP_V4"

	parityPretty = make(map[int]string)
	parityPretty[even] = "EVEN"
	parityPretty[odd] = "ODD"
}

func entityTypeOutput(n int) string {
	return entityTypePretty[n]
}

func parityOutput(n int) string {
	return parityPretty[n]
}
