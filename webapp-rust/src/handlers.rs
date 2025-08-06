use poem::{handler, web::Json};
use serde::Serialize;
use std::net::IpAddr::V4;
use sysinfo::{CpuRefreshKind, Networks, System};

#[derive(Serialize, Debug)]
enum EntityType {
    CPU_AMOUNT,
    IP_V4,
    TIME_NOW,
    DATE_NOW,
}

#[derive(Serialize, Debug)]
struct CpuAmount {
    #[serde(rename(serialize = "type"))]
    entity_type: EntityType,
    payload: usize,
}

#[handler]
pub fn cpu_info() -> Json<CpuAmount> {
    let mut sys = System::new();
    sys.refresh_cpu_list(CpuRefreshKind::everything());

    Json(CpuAmount {
        entity_type: EntityType::CPU_AMOUNT,
        payload: sys.cpus().len(),
    })
}

#[allow(dead_code)]
#[derive(Serialize, Debug)]
enum IpAddrParity {
    EVEN,
    ODD,
}
#[derive(Serialize, Debug)]
struct IfInfo {
    name: String,
    #[serde(rename(serialize = "ip"))]
    ip_addr: String,
}

#[derive(Serialize, Debug)]
struct NetInfo {
    #[serde(rename(serialize = "type"))]
    entity_type: EntityType,
    count: IpAddrParity,
    payload: Vec<IfInfo>,
}

#[handler]
pub fn net_info() -> Json<NetInfo> {
    let mut net = Networks::new_with_refreshed_list();

    let mut collector: Vec<(String, String)> = Vec::new();

    for (if_name, if_data) in &net {
        for net_data in if_data.ip_networks() {
            match net_data.addr {
                V4(addr) => {
                    collector.push((
                        if_name.clone(),
                        addr.octets().map(|o| o.to_string()).join("."),
                    ));
                }
                _ => {}
            }
        }
    }

    Json(NetInfo {
        entity_type: EntityType::IP_V4,
        count: match collector.len() % 2 {
            0 => IpAddrParity::EVEN,
            _ => IpAddrParity::ODD,
        },
        payload: collector
            .iter()
            .map(|t| IfInfo {
                name: t.0.clone(),
                ip_addr: t.1.clone(),
            })
            .collect(),
    })
}

#[handler]
pub fn get_time() -> &'static str {
    "Got TIME info"
}

#[handler]
pub fn get_date() -> &'static str {
    "Got DATE info"
}
