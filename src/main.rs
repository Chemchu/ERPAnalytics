#![allow(non_snake_case)]
#[macro_use]
extern crate rocket;
extern crate juniper;

pub mod salesAnalyse;
pub mod structs;

use rocket::response::content;
use rocket::serde::json::{Json, Value};
use rocket::serde::{Deserialize, Serialize};
use salesAnalyse::{GetVentasPorHora, VentaPorHora};
use structs::Venta;

// #[path = "./sales/salesAnalytics.rs"]
// use SalesAnalytics::Venta;

#[derive(Serialize, Deserialize, Debug)]
struct GraphQLReqBody {
    query: String,
    variables: Option<Value>,
}

#[derive(Serialize, Deserialize, Debug)]
struct GraphQLVariables {
    find: Option<String>,
    variables: Option<Value>,
}

// #[derive(Serialize, Deserialize, Debug)]
// struct SalesAnalyticsResponse {
//     successful: bool,
//     message: &'static str,
//     analytics: Option<Vec<&'static str>>,
// }

#[derive(Serialize, Deserialize, Debug)]
struct SalesAnalyticsResponse {
    successful: bool,
    message: &'static str,
    analytics: Vec<VentaPorHora>,
}

#[get("/")]
fn index() -> Result<content::RawHtml<String>, String> {
    let html = include_str!("./html/index.html");
    return Ok(content::RawHtml(String::from(html)));
}

#[get("/")]
fn api() -> Result<content::RawHtml<String>, String> {
    let html = include_str!("./html/apiDescription.html");
    return Ok(content::RawHtml(String::from(html)));
}

#[post("/sales", format = "application/json", data = "<sales>")]
fn sales(sales: Json<Vec<Venta>>) -> Result<Json<SalesAnalyticsResponse>, String> {
    let ventasPorHora = GetVentasPorHora(sales.to_vec());

    let res: SalesAnalyticsResponse = SalesAnalyticsResponse {
        successful: true,
        message: "Hola mundo!!",
        analytics: ventasPorHora.unwrap(),
    };
    return Ok(Json(res));
}

// Se encarga de escuchar las peticiones en GraphQL para reenviar al microservicio indicado
#[post("/", format = "json", data = "<body>")]
fn graphql(body: Json<GraphQLReqBody>) -> Result<Json<String>, std::io::Error> {
    let res = String::from("GraphQL reforward");
    println!("{:#?}", body.variables);

    // Modificar campo en JSON
    // if let Some(name) = v.get_mut("query") {
    //     *name = "new name".into();
    // }

    return Ok(Json(res));
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index])
        .mount("/api", routes![api, sales])
        .mount("/graphql", routes![graphql])
}
