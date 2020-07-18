package prived.medved.controllers;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import prived.medved.response.*;
import prived.medved.response.enums.EntityType;
import prived.medved.response.enums.IPsCount;

import java.net.*;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.*;
import java.util.stream.Collectors;

@RestController
public class EntryPoint {

    @GetMapping("/date")
    public SimpleEntity getNowDate(){

        return new SimpleEntity(EntityType.DATE_NOW, LocalDate.now().format(DateTimeFormatter.ofPattern("yyyy MM dd")));
    }

    @GetMapping("/time")
    public SimpleEntity getNowTime(){

        return new SimpleEntity(EntityType.TIME_NOW, LocalDateTime.now().format(DateTimeFormatter.ofPattern("hh:mm:ss")));
    }

    @GetMapping("/ip")
    public IfaceEntity getIP() throws SocketException {

        Enumeration<NetworkInterface> ifaces = NetworkInterface.getNetworkInterfaces();

        List<IFaceIPPair> ipList = Collections.list(ifaces).stream().map(iface -> {

            String name = iface.getDisplayName();
            LinkedList<IFaceIPPair> addrList = new LinkedList<IFaceIPPair>();

            for (InetAddress ifAddr : Collections.list(iface.getInetAddresses())){

                if (ifAddr instanceof Inet4Address){
                    addrList.add(new IFaceIPPair(name, ifAddr.getHostAddress()));
                }
            }
            return addrList;

        }).flatMap(ifList -> ifList.stream()).collect(Collectors.toList());

        IPsCount count;

        if (ipList.size() % 2 == 0) {
            count = IPsCount.EVEN;
        } else {
            count = IPsCount.ODD;
        }
        return new IfaceEntity(EntityType.IP_V4, count, ipList);
    }

    @GetMapping("/cpu")
    public SimpleEntity getCPUs(){

        int cpuCount = Runtime.getRuntime().availableProcessors();

        return new SimpleEntity(EntityType.CPU_AMOUNT, String.valueOf(cpuCount));
    }
}
