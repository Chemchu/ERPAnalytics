fn RequestSales(ventas: Vec<Venta>) -> Result<(), Box<dyn std::error::Error>> {
    // let mut res = reqwest::blocking::get("http://localhost:8080/graphql")?;
    // let mut body = String::new();
    // res.read_to_string(&mut body)?;

    // println!("Status: {}", res.status());
    // println!("Headers:\n{:#?}", res.headers());
    // println!("Body:\n{}", body);

    Ok(())
}

fn main() {
    RequestSales()
}
