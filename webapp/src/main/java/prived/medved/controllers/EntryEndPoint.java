package prived.medved.controllers;

import org.springframework.stereotype.Component;
import prived.medved.response.*;
import prived.medved.response.enums.EntityType;
import prived.medved.response.enums.IPsCount;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import java.net.*;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.*;
import java.util.stream.Collectors;

@Component
@Path("/stat/")
@Produces("application/json")
public class EntryEndPoint {

    @GET
    @Path("date")
    public SimpleEntity getNowDate(){

        return new SimpleEntity(EntityType.DATE_NOW, LocalDate.now().format(DateTimeFormatter.ofPattern("yyyy MM dd")));
    }

    @GET
    @Path("time")
    public SimpleEntity getNowTime(){

        return new SimpleEntity(EntityType.TIME_NOW, LocalDateTime.now().format(DateTimeFormatter.ofPattern("hh:mm:ss")));
    }

    @GET
    @Path("ip")
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

    @GET
    @Path("cpu")
    public SimpleEntity getCPUs(){

        int cpuCount = Runtime.getRuntime().availableProcessors();

        return new SimpleEntity(EntityType.CPU_AMOUNT, String.valueOf(cpuCount));
    }
}
