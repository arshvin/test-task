use poem::{handler, web::Json};
use serde::Serialize;
use sysinfo::{CpuRefreshKind, System};

#[derive(Serialize, Debug)]
enum EntityType {
    CPU_AMOUNT,
    IP_V4,
    TIME_NOW,
    DATE_NOW,
}

#[allow(dead_code)]
#[derive(Serialize, Debug)]
enum IpAddrParity {
    EVEN,
    ODD,
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

#[handler]
pub fn net_info() -> &'static str {
    "Got NET info"
}

#[handler]
pub fn get_time() -> &'static str {
    "Got TIME info"
}

#[handler]
pub fn get_date() -> &'static str {
    "Got DATE info"
}
