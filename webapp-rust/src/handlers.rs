use poem::handler;
// use serde::Serialize;

// #[derive(Serialize, Debug)]
// struct CpuAmount {
//     r#type: String,
//     payload: u8,
// }

#[handler]
pub fn cpu_info() -> &'static str {
    "Got CPU info"
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
