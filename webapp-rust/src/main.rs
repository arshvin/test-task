use std::env;

pub mod handlers;

use seahorse::{App, Context, Flag, FlagType};
use warp::Filter;

fn main() {
    let args: Vec<String> = env::args().collect();

    let app = App::new(env!("CARGO_PKG_NAME"))
        .description(env!("CARGO_PKG_DESCRIPTION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .flag(
            Flag::new("address", FlagType::String)
                .description("Bind address to (default: 0.0.0.0)")
                .alias("a"),
        )
        .flag(
            Flag::new("port", FlagType::String)
                .description("Bind port to (required)")
                .alias("p"),
        )
        .action(default_action);

    app.run(args);
}

fn default_action(c: &Context) {
    let address = match c.string_flag("address") {
        Ok(addr) => addr,
        Err(_) => "0.0.0.0".to_string(),
    };

    let port = c.string_flag("port").expect("--port/-p flag is required");

    run_server(address, port);
}

#[tokio::main]
async fn run_server(addr: String, port: String) {
    let mut it = addr.split('.');
    let ip = if let [Some(octet0), Some(octet1), Some(octet2), Some(octet3), None] =
        [it.next(), it.next(), it.next(), it.next(), it.next()]
    {
        let mut arr: [u8; 4] = [0; 4];
        arr[0] = octet0.parse::<u8>().unwrap();
        arr[1] = octet1.parse::<u8>().unwrap();
        arr[2] = octet2.parse::<u8>().unwrap();
        arr[3] = octet3.parse::<u8>().unwrap();
        arr
    } else {
        panic!("Could not parse IP addr v4 argument")
    };

    let Ok(p) = port.parse::<u16>() else {
        panic!("Could not parse port argument")
    };

    let routes = warp::any().map(|| "Hello, World!");

    warp::serve(routes).run((ip, p)).await;
}
