use chrono::{DateTime, NaiveDateTime, Timelike, Utc};
use rocket::serde::{Deserialize, Serialize};
use std::collections::hash_map::Entry::{Occupied, Vacant};
use std::{
    collections::HashMap,
    time::{Duration, UNIX_EPOCH},
};

use crate::structs::Venta;

#[derive(Serialize, Deserialize, Debug)]
pub struct VentaPorHora {
    pub _id: String,
    pub numVentas: u32,
    pub valor: f32,
    pub hora: u8,
}

fn RequestSales(ventas: Vec<Venta>) -> Result<(), Box<dyn std::error::Error>> {
    // let mut res = reqwest::blocking::get("http://localhost:8080/graphql")?;
    // let mut body = String::new();
    // res.read_to_string(&mut body)?;

    // println!("Status: {}", res.status());
    // println!("Headers:\n{:#?}", res.headers());
    // println!("Body:\n{}", body);

    Ok(())
}

// fn Media() -> Result<f64, String> {}

// fn NumVentas() -> Result<f64, String> {}

fn EpochToDateTime(fecha: String) -> DateTime<Utc> {
    let fecha: u64 = fecha.parse().unwrap();
    let d = UNIX_EPOCH + Duration::from_millis(fecha);
    // Create DateTime from SystemTime
    let datetime = DateTime::<Utc>::from(d);
    return datetime;
    // // Formats the combined date and time with the specified format string.
    // let timestamp_str = datetime.format("%Y-%m-%d %H:%M:%S.%f").to_string();
}

pub fn GetVentasPorHora(ventas: Vec<Venta>) -> Result<Vec<VentaPorHora>, String> {
    let ventasPorHora: Vec<VentaPorHora> = Vec::new();
    let mut ventasHashMap: HashMap<u32, f32> = HashMap::new();

    for venta in ventas.iter() {
        let fecha = EpochToDateTime(venta.createdAt.clone());
        let key = fecha.hour();

        let val = match ventasHashMap.entry(key) {
            Vacant(entry) => entry.insert(0.0),
            Occupied(entry) => entry.into_mut(),
        };
        *val += venta.precioVentaTotal;
    }
    println!("{:?}", ventasHashMap);
    Ok(ventasPorHora)
}
