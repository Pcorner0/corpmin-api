USE corpmin;

DROP TABLE IF EXISTS vendedores;

CREATE TABLE IF NOT EXISTS vendedores (
	idvendedor INT AUTO_INCREMENT PRIMARY KEY ,
    nombre VARCHAR(50) NOT NULL,
	primerapellido VARCHAR(50) DEFAULT NULL,
    segundoapellido VARCHAR(50) DEFAULT NULL,
    telefono VARCHAR(20) DEFAULT NULL,
    email VARCHAR(100) NOT NULL,
    rol VARCHAR(10) DEFAULT NULL,
    oficina VARCHAR(50) DEFAULT NULL,
    created_at DATETIME,
	updated_at DATETIME,
	deleted_at DATETIME
    )ENGINE=INNODB AUTO_INCREMENT=1000;

DROP TABLE IF EXISTS clientes;
CREATE TABLE IF NOT EXISTS clientes (
	id INT AUTO_INCREMENT KEY,
    rfc VARCHAR(50) NOT NULL,
    razonsocial VARCHAR(250) NOT NULL,
    calle VARCHAR(50) DEFAULT NULL,
    noext VARCHAR(50) DEFAULT NULL,
    noint VARCHAR(50) DEFAULT NULL,
    cp VARCHAR(20) NOT NULL,
    colonia VARCHAR(50) DEFAULT NULL,
    municipio VARCHAR(50) DEFAULT NULL,
    estado VARCHAR(50) DEFAULT NULL,
    contactonombre VARCHAR(50) NOT NULL,
	contactoprimerapellido VARCHAR(50) DEFAULT NULL,
    contactosegundoapellido VARCHAR(50) DEFAULT NULL,
    contactopuesto VARCHAR(50) DEFAULT NULL,
	email VARCHAR(100) DEFAULT NULL,
    telefonooficina VARCHAR(50) DEFAULT NULL, 
    telefonopersonal VARCHAR(50) DEFAULT NULL, 
    prospecto VARCHAR(50) DEFAULT NULL, 
    idvendedor INT DEFAULT NULL,
    estatus VARCHAR(50) DEFAULT NULL, 
	created_at DATETIME,
	updated_at DATETIME,
	deleted_at DATETIME
    )ENGINE=INNODB;

DROP TABLE IF EXISTS operaciones;
CREATE TABLE IF NOT EXISTS operaciones (
	i INT AUTO_INCREMENT KEY,
	id INT DEFAULT NULL,
    rfc VARCHAR(50) DEFAULT NULL,
    razonsocial VARCHAR(250) DEFAULT NULL,
    actividad VARCHAR(50) DEFAULT NULL,
    estatus VARCHAR(50) DEFAULT NULL,
    idvendedor INT DEFAULT NULL,
    observaciones VARCHAR(250) DEFAULT NULL,
	created_at DATETIME
    )ENGINE=INNODB;
    


SHOW columns FROM clientes;
SHOW columns FROM empleados;
SHOW columns FROM operaciones;

SELECT estatus, COUNT(*)
FROM operaciones
GROUP BY estatus
HAVING created_at > MAKEDATE(2019,31);

-- GROUP BY RANGEDATE AND ESTATUS
SELECT estatus, count(*)
FROM   operaciones o1
WHERE  created_at=(SELECT MAX(o2.created_at)
              FROM operaciones o2
              WHERE o1.rfc = o2.rfc)
GROUP BY estatus;

SELECT *
FROM users;

SELECT * FROM users WHERE email = 'pedro' AND deleted_at IS NULL LIMIT 1;

SELECT idvendedor, count(*) 
FROM operaciones
WHERE estatus = "CLIENTE"
GROUP BY idvendedor;

UPDATE operaciones SET vendedor = CONCAT(
	vendedores.nombre, 
    vendedores.primerapellido, 
    vendedores.segundoapellido)  
WHERE operaciones.idvendedor = vendedores.idvendedor ;

SELECT operaciones.*, CONCAT(nombre, primerapellido, segundoapellido)   
FROM operaciones 
INNER JOIN vendedores 
ON operaciones.idvendedor = vendedores.idvendedor 
ORDER BY created_at DESC;

SELECT 
  tbl.id,
  signin,
  signout
FROM tbl
  INNER JOIN (
    SELECT id, MAX(signin) AS maxsign FROM tbl GROUP BY id
  ) ms ON tbl.id = ms.id AND signin = maxsign
WHERE tbl.id=1;

SELECT * FROM operaciones INNER JOIN ( SELECT id AS idSecondary, MAX(created_at) AS maxsign FROM operaciones GROUP BY idSecondary) ms ON operaciones.id = ms.idSecondary AND created_at = maxsign WHERE operaciones.id=2

SELECT *, MAX(created_at) 
from operaciones
WHERE id = 2
group by razonsocial
HAVING created_at =  MAX(created_at);

SELECT *
FROM operaciones
WHERE id = 2;


SELECT MAX(created_at)
FROM operaciones
WHERE id = 2;

ALTER TABLE operaciones
DROP COLUMN vendedor;

SELECT *
FROM operaciones
ORDER BY created_at DESC;

SELECT * 
FROM clientes
ORDER BY created_at DESC;

 SELECT * FROM clientes 
 WHERE estatus != 'CLIENTE'
 AND deleted_at IS NULL;



SELECT *
FROM users;

SELECT *
FROM vendedores;

SELECT *
FROM clientes
ORDER BY created_at DESC;

SELECT *
FROM operaciones
ORDER BY created_at DESC;