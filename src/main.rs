#![allow(non_snake_case)]
#[macro_use]
extern crate rocket;

use rocket::serde::json::Json;
use rocket::serde::Serialize;

#[derive(Serialize)]
struct SalesAnalyticsResponse {
    successful: bool,
    message: &'static str,
    analytics: Option<Vec<&'static str>>,
}

// #[get("/")]
// fn index() -> Result<HTML, String> {}

#[get("/")]
fn api() -> &'static str {
    "API Description"
}

#[get("/sales")]
fn sales() -> Result<Json<SalesAnalyticsResponse>, String> {
    let res = SalesAnalyticsResponse {
        successful: true,
        message: "Hola mundo!!",
        analytics: None,
    };
    let ventas = Json(res);
    //let ventasErr = String::from("No ha sido posible acceder a este contenido");

    return Ok(ventas);

    // let res = if ventas !== Null { Ok(ventas) } else { Err(ventasErr) };
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/api", routes![api, sales])
    //.mount("/", routes![index])
}
