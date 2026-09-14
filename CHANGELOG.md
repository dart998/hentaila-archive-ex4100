# Changelog

## [0.1.3] - 2026-09-14

### Corregido
- La descarga masiva indexa bajo demanda las series abiertas directamente desde el catálogo aunque no estén en las listas personales ni se hayan indexado previamente en SQLite.
- El crawler envía la cookie de sesión guardada al solicitar la ficha de una serie, para que el indexado bajo demanda use la misma sesión de HentaiLA.

## [0.1.2] - 2026-09-14

### Cambiado
- Se elimina del sitio espejado la tarjeta informativa «Evita los baneos», incluso si HentaiLA la reinserta durante la hidratación.

## [0.1.1] - 2026-09-14

### Cambiado
- La cookie de HentaiLA se muestra en texto visible y permanece cargada en el formulario después de guardarla.
- Se elimina el control para borrar la cookie desde el panel de administración.
- Se elimina el banner promocional superior del sitio espejado.

## [0.1.0] - 2026-09-14

### Añadido

- Primer mirror y crawler independiente de HentaiLA para WD My Cloud EX4100 (`linux/arm/v7`).
- Navegación local, panel de administración, caché incremental, filtrado publicitario y recursos del CDN.
- Sincronización de sesión, listas, favoritos y episodios vistos de HentaiLA.
- Matching y reproducción desde una carpeta de vídeos independiente en el RAID.
- Detección de Mega, YourUpload, StreamWish, MP4Upload, VidHide y Voe.
- Descarga masiva mediante Mega, seguimiento de progreso y cancelación segura.
- Piloto configurable de tres series; `MIRROR_SERIES_LIMIT=0` habilita el catálogo completo.
- Stack de Portainer separado en el puerto `8091` e imagen Docker ARMv7 propia.
