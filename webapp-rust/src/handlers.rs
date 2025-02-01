
use serde::Serialize;

#[derive(Serialize, Debug)]
struct cpu_amount{
    r#type: String,
    payload: u8
}
