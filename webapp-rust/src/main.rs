use std::env;

use seahorse::{App, Flag, FlagType, Context};
// use std::env;
fn main() {
    let args: Vec<String> = env::args().collect();

    let app = App::new(env!("CARGO_PKG_NAME"))
    .description(env!("CARGO_PKG_DESCRIPTION"))
    .author(env!("CARGO_PKG_AUTHORS"))
    .flag(
        Flag::new("address", FlagType::String)
        .description("Bind address to (default: 0.0.0.0)").alias("a")
    )
    .flag(
        Flag::new("port", FlagType::String)
        .description("Bind port to (required)").alias("p")
    ).action(default_action);

    app.run(args);
}

fn default_action(c: &Context){
    let address = match c.string_flag("address") {
        Ok(addr) => addr,
        Err(_) => "0.0.0.0".to_string(),
    };

    let port = c.string_flag("port").expect("'port' flag is required");
    
    println!("Gathered args: {}:{}",address,port);
}
