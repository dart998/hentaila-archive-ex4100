# Changelog

## [0.1.8] - 2026-09-17

### Corregido
- Una descarga finalizada deja de bloquear una nueva comprobación de episodios: al volver a pulsar descargar se reescanea la biblioteca y se inicia un lote nuevo con el estado actual.

### Diagnóstico
- Se añaden logs `[MEGA DEBUG]` para registrar de forma segura cómo se extraen e interpretan las URL de Mega, ocultando la clave de descifrado y sin registrar cookies.

## [0.1.7] - 2026-09-17

### Corregido
- Elimina directamente el bloque superior rojo `header > div.bg-[#f87171].text-white`, evitando que el banner promocional quede vacío ocupando 56 px tras retirar su contenido.

## [0.1.6] - 2026-09-17

### Corregido
- Se retira la interceptación global de `http.DefaultTransport` introducida en 0.1.5.
- La limpieza de anuncios queda limitada a los wrappers exactos de los iframes `Ads 728x90` y `Ads 300x250` de `adtng.com`, sin ascender por contenedores del layout.
- Se elimina `collapseEpisodeGap()` para evitar modificar bloques legítimos de la página.
- El banner superior promocional vuelve a eliminarse por contenido sin depender de su posición vertical.

## [0.1.5] - 2026-09-17

### Corregido
- Los banners publicitarios de `adtng.com`, incluidos los formatos `Ads 728x90` y `Ads 300x250`, eliminan también su contenedor para no dejar huecos vacíos en las páginas.
- La limpieza se vuelve a aplicar tras la hidratación mediante `MutationObserver`, cubriendo anuncios que HentaiLA inserte dinámicamente.

## [0.1.4] - 2026-09-17

### Cambiado
- `MIRROR_SERIES_LIMIT` pasa a `0` por defecto, eliminando el límite piloto de tres series y permitiendo espejar el catálogo completo.

### Corregido
- El matching local conserva y compara el número/secuela del título antes de aceptar coincidencias por stem o alias, evitando que series como `Onichichi 2` reutilicen por error la carpeta de `Onichichi`.
- Las páginas de episodio colapsan contenedores vacíos residuales para evitar huecos dejados por bloques publicitarios filtrados.

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
