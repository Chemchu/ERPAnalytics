use rocket::serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Venta {
    pub _id: String,
    pub productos: Vec<ProductoVendido>,
    pub dineroEntregadoEfectivo: f32,
    pub dineroEntregadoTarjeta: f32,
    pub precioVentaTotalSinDto: f32,
    pub precioVentaTotal: f32,
    pub cambio: f32,
    pub cliente: Cliente,
    pub vendidoPor: Empleado,
    pub modificadoPor: Empleado,
    pub tipo: String,
    pub descuentoEfectivo: f32,
    pub descuentoPorcentaje: f32,
    pub tpv: String,
    pub updatedAt: String,
    pub createdAt: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProductoVendido {
    pub _id: String,
    pub nombre: String,
    pub precioCompra: f32,
    pub precioVenta: f32,
    pub cantidadVendida: f32,
    pub dto: f32,
    pub iva: f32,
    pub ean: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Cliente {
    pub _id: String,
    pub nombre: String,
    pub calle: String,
    pub cp: String,
    pub nif: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Empleado {
    pub _id: String,
    pub nombre: String,
    pub apellidos: String,
    pub email: String,
    pub rol: String,
}
