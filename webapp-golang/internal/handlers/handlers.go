package handlers

import (
	"fmt"
	"net/http"
	"net/netip"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	gopsutilCpu "github.com/shirou/gopsutil/v4/cpu"
	gopsutilNet "github.com/shirou/gopsutil/v4/net"
)

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

func CpuAmount(c echo.Context) error {
	nproc, err := gopsutilCpu.Counts(true)
	if err != nil {
		log.Errorf("Could not get CPU amount: %s", err)
		return err
	}

	return c.JSON(http.StatusOK, &CpuResponse{
		EntityType: entityTypeOutput(cpuAmount),
		Payload:    nproc,
	})
}

func AddrIpV4Amount(c echo.Context) error {
	interfaces, err := gopsutilNet.Interfaces()
	if err != nil {
		log.Errorf("Could not get Net Interfaces info: %s", err)
	}

	response := &AddrIpV4Response{
		EntityType: entityTypeOutput(addrIpV4Amount),
	}
	iFsInfo := make([]*IfInfo, 0, 5)

	for _, ifConf := range interfaces {
		for _, ifIpAddr := range ifConf.Addrs {

			ipAddr := strings.Split(ifIpAddr.Addr, "/")[0] // Because gopsutil returns ip addr with mask in one string which will not be parsed
			ipParsed, err := netip.ParseAddr(ipAddr)
			if err != nil {
				log.Errorf("Could not parse IP address: %s: %s", ipAddr, err)
				return err
			}

			if ipParsed.Is4() {
				fmt.Fprintf(os.Stdout, "Adding of ifConf.Name:%s ifConf.Addrs:%s\n", ifConf.Name, ifConf.Addrs)

				iFsInfo = append(iFsInfo, &IfInfo{
					Name:      ifConf.Name,
					IpAddress: ipParsed.String(),
				})
			}
		}
	}

	response.Payload = iFsInfo

	if len(iFsInfo)%2 != 0 {
		response.Count = parityOutput(odd)
	} else {
		response.Count = parityOutput(even)
	}

	return c.JSON(http.StatusOK, response)
}
