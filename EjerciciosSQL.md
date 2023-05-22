En el problema descrito, se requiere escribir consultas SQL para realizar diversas tareas relacionadas con las tablas "Usuarios" y "Pedidos" de una base de datos.

Imagina que tienes una base de datos con dos tablas: "Usuarios" y "Pedidos".

La tabla "Usuarios" contiene la siguiente estructura:
* ID (entero, clave primaria)
* Nombre (cadena de caracteres)
* Email (cadena de caracteres)
* Fecha de registro (fecha y hora)

La tabla "Pedidos" contiene la siguiente estructura:
* ID (entero, clave primaria)
* Usuario ID (entero, clave foránea de la tabla "Usuarios")
* Producto (cadena de caracteres)
* Cantidad (entero)
* Fecha de pedido (fecha y hora)

```sql
-- Creación de la tabla "Usuarios"
CREATE TABLE Usuarios (
    ID INT PRIMARY KEY,
    Nombre VARCHAR(255),
    Email VARCHAR(255),
    FechaRegistro DATETIME
);

-- Creación de la tabla "Pedidos"
CREATE TABLE Pedidos (
    ID INT PRIMARY KEY,
    UsuarioID INT,
    Producto VARCHAR(255),
    Cantidad INT,
    FechaPedido DATETIME,
    FOREIGN KEY (UsuarioID) REFERENCES Usuarios(ID)
);
```

A continuación, se detallan las consultas a realizar:

a) Obtener el nombre y el email de todos los usuarios registrados en orden alfabético:
```sql
SELECT Nombre, Email FROM Usuarios ORDER BY Nombre ASC;
```

b) Contar cuántos pedidos se han realizado en total:
```sql
SELECT COUNT(*) AS total_pedidos FROM Pedidos;
```

c) Obtener el nombre de los usuarios que han realizado al menos un pedido:
```sql
SELECT DISTINCT Usuarios.Nombre FROM Usuarios INNER JOIN Pedidos ON Usuarios.ID = Pedidos.UsuarioID;
```

d) Calcular la cantidad total de productos pedidos por cada usuario:
```sql
SELECT Usuarios.Nombre, SUM(Pedidos.Cantidad) AS total_productos_pedidos 
FROM Usuarios INNER JOIN Pedidos ON Usuarios.ID = Pedidos.UsuarioID 
GROUP BY Usuarios.Nombre;
```

e) Obtener la fecha y el producto del pedido más reciente realizado por cada usuario:
```sql
SELECT Usuarios.Nombre, Pedidos.FechaPedido, Pedidos.Producto 
FROM Usuarios INNER JOIN Pedidos ON Usuarios.ID = Pedidos.UsuarioID 
WHERE Pedidos.ID IN (
    SELECT MAX(ID) 
    FROM Pedidos 
    GROUP BY UsuarioID
);
```

f) Obtener el nombre y el email de los usuarios que han realizado pedidos en los últimos 7 días:
```sql
SELECT Usuarios.Nombre, Usuarios.Email 
FROM Usuarios INNER JOIN Pedidos ON Usuarios.ID = Pedidos.UsuarioID 
WHERE Pedidos.FechaPedido >= (CURRENT_DATE - INTERVAL '7 days');
```

g) Calcular el promedio de la cantidad de productos pedidos por usuario:
```sql
SELECT AVG(Pedidos.Cantidad) AS promedio_productos_pedidos 
FROM Pedidos;
```

Estas consultas SQL permiten obtener la información requerida de la base de datos según los criterios especificados. Los postulantes pueden utilizar un entorno de desarrollo SQL, como PgAdmin, para escribir y ejecutar estas consultas y obtener los resultados esperados.