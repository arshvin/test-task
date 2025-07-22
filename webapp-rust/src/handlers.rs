use serde::Serialize;

#[derive(Serialize, Debug)]
struct CpuAmount {
    r#type: String,
    payload: u8,
}
