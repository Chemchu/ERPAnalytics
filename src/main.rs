#![allow(non_snake_case)]
#[macro_use]
extern crate rocket;
extern crate juniper;

use rocket::response::content;
use rocket::serde::json::from_str;
use rocket::serde::json::Json;
use rocket::serde::Deserialize;
use rocket::serde::Serialize;

#[derive(Deserialize)]
struct GraphQLReqBody {
    query: String,
    variables: Option<String>,
}

#[derive(Serialize)]
struct SalesAnalyticsResponse {
    successful: bool,
    message: &'static str,
    analytics: Option<Vec<&'static str>>,
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

#[get("/sales")]
fn sales() -> Result<Json<SalesAnalyticsResponse>, String> {
    let res: SalesAnalyticsResponse = SalesAnalyticsResponse {
        successful: true,
        message: "Hola mundo!!",
        analytics: None,
    };
    return Ok(Json(res));
}

// Se encarga de escuchar las peticiones en GraphQL para reenviar al microservicio indicado
#[post("/", format = "json", data = "<body>")]
fn graphql(body: String) -> Result<Json<String>, std::io::Error> {
    let res = String::from("GraphQL reforward");
    // let v = body.variables.clone().unwrap();

    // let json = println!(json);
    // let r: GraphQLReqBody = from_str(&body)?;

    // println!("{:?}", r);

    println!("{}", body);
    // println!("{}", v);

    return Ok(Json(res));
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![index])
        .mount("/api", routes![api, sales])
        .mount("/graphql", routes![graphql])
}
