package handlers

import (
	"net/http"
	"net/netip"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	gopsutilCpu "github.com/shirou/gopsutil/v4/cpu"
	gopsutilNet "github.com/shirou/gopsutil/v4/net"
)

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

func CurrentDateTime(c echo.Context, dateOnly bool) error {
	response := &DateTimeNow{}
	now := time.Now()

	switch dateOnly {
	case true:
		response.EntityType = entityTypeOutput(dateNow)
		response.Payload = now.Format(time.DateOnly)
	case false:
		response.EntityType = entityTypeOutput(timeNow)
		response.Payload = now.Format(time.TimeOnly)
	}

	return c.JSON(http.StatusOK, response)
}
