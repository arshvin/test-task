use std::env;

pub mod handlers;

use poem::{Route, Server, get, listener::TcpListener};
use seahorse::{App, Context, Flag, FlagType};

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
            Flag::new("port", FlagType::Uint)
                .description("Bind port to (required)")
                .alias("p"),
        )
        .action(process_args);

    app.run(args);
}

fn process_args(c: &Context) {
    let address = match c.string_flag("address") {
        Ok(addr) => addr,
        Err(_) => "0.0.0.0".to_string(),
    };

    let port = c.uint_flag("port").expect("--port/-p flag is required");

    let _ = ignition(format!("{}:{}", address, port));
}

#[tokio::main]
async fn ignition(addr: String) -> Result<(), std::io::Error> {
    let routes = Route::new()
        .at("/cpu", get(handlers::cpu_info))
        .at("/ip", get(handlers::net_info))
        .at("/time", get(handlers::get_time))
        .at("/date", get(handlers::get_date));

    Server::new(TcpListener::bind(addr)).run(routes).await
}
