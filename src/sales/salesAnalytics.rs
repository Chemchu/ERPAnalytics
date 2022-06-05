#[derive(Serialize, Deserialize, Debug)]
struct Venta {
    _id: ID,
    productos: [ProductoVendido],
    dineroEntregadoEfectivo: i32,
    dineroEntregadoTarjeta: i32,
    precioVentaTotalSinDto: i32,
    precioVentaTotal: i32,
    cambio: i32,
    cliente: Cliente,
    vendidoPor: Empleado,
    modificadoPor: Empleado,
    tipo: String,
    descuentoEfectivo: i32,
    descuentoPorcentaje: i32,
    tpv: ID,
    updatedAt: u64,
    createdAt: u64,
}

struct ProductoVendido {
    _id: ID,
    nombre: String,
    precioCompra: i32,
    precioVenta: i32,
    cantidadVendida: i32,
    dto: i32,
    iva: i32,
    ean: i32,
}

struct Cliente {
    _id: ID,
    nombre: String,
    precioCompra: i32,
    precioVenta: i32,
    cantidadVendida: i32,
    dto: i32,
    iva: i32,
    ean: i32,
    variables: Option<Value>,
}

struct Empleado {
    nombre: String,
    precioCompra: i32,
    precioVenta: i32,
    cantidadVendida: i32,
    dto: i32,
    iva: i32,
    ean: i32,
    _id: ID,
    variables: Option<Value>,
}

fn RequestSales() -> Result<(), Box<dyn std::error::Error>> {
    let mut res = reqwest::blocking::get("http://localhost:8080/graphql")?;
    let mut body = String::new();
    res.read_to_string(&mut body)?;

    println!("Status: {}", res.status());
    println!("Headers:\n{:#?}", res.headers());
    println!("Body:\n{}", body);

    Ok(())
}

fn main() {
    RequestSales()
}
