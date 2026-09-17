# HentaiLA Archive para WD EX4100

Mirror y crawler local de [HentaiLA](https://hentaila.com/) para un WD My Cloud EX4100 (`linux/arm/v7`). Es un proyecto independiente de AnimeAV1 Archive: conserva las mismas funciones, pero usa su propio contenedor, puerto, SQLite, caché y carpeta de vídeos.

## Funciones

- mirror navegable de HentaiLA en `/` y administración en `/admin`;
- sesión de HentaiLA guardada solo en backend;
- sincronización de listas personales y prioridad **Viendo → Planeado → Completado → resto**;
- mirror incremental y reanudable de HTML, SvelteKit, imágenes y recursos del CDN;
- filtrado de publicidad y destinos externos mediante EasyList;
- indexado recursivo de la biblioteca montada en `/library`;
- matching por título, alias, franquicia, temporada y patrón de episodio;
- reproducción local cuando el navegador admite el archivo y fallback a reproductores online;
- proveedores HentaiLA: Mega, YourUpload, StreamWish, MP4Upload, VidHide y Voe;
- marcado de episodios vistos, favoritos/listas, refresco de fichas y descarga masiva mediante Mega;
- notificaciones de episodios nuevos para series en estado **Viendo**;
- imagen Docker específica para `linux/arm/v7` publicada mediante GitHub Actions.

La aplicación no mueve, renombra ni elimina automáticamente los vídeos existentes.

## Despliegue en Portainer CE

El stack completo está en [`docker-compose.yml`](docker-compose.yml). Sus valores predeterminados no colisionan con AnimeAV1:

| Recurso | HentaiLA Archive |
| --- | --- |
| Puerto web | `8091` |
| Datos y SQLite | `/mnt/HD/HD_a2/Public/hentaila-archive` |
| Vídeos | `/mnt/HD/HD_a2/Public/Anime/Hentai` |
| Imagen | `ovelayos/hentaila-archive-ex4100:0.1.9` |

Después de desplegar:

1. abre `http://<EX4100>:8091/admin`;
2. guarda la cookie de sesión de HentaiLA;
3. pulsa **Actualizar listas HLA**;
4. pulsa **Reindexar /library**;
5. pulsa **Descargar / actualizar sitio**;
6. valida el catálogo en `http://<EX4100>:8091/`.

## Catálogo completo

El stack usa `MIRROR_SERIES_LIMIT=0` por defecto, por lo que la sincronización del sitio puede admitir todos los slugs `/media/...` descubiertos. `CRAWLER_BATCH_SIZE=3` sigue limitando cuántas series procesa el crawler en cada lote para mantener una carga razonable en el EX4100.

## Persistencia

```text
/data/
|-- db/archive.sqlite
|-- site/                  mirror y caché
|-- metadata/
|-- images/
|-- logs/
`-- tmp/

/library/                  vídeos de HentaiLA
```

Se indexan `.mkv`, `.mp4`, `.avi`, `.webm`, `.m4v` y `.mov`. La cookie se muestra en texto visible únicamente en `/admin`; no se incrusta en el sitio espejado ni se reenvía al CDN.

## Variables principales

| Variable | Predeterminado | Uso |
| --- | --- | --- |
| `WEB_PORT` | `8091` | Puerto publicado |
| `TZ` | `Europe/Madrid` | Zona horaria |
| `HENTAILA_BASE_URL` | `https://hentaila.com` | Origen del mirror |
| `MIRROR_SERIES_LIMIT` | `0` | Sin límite de series del catálogo |
| `CRAWLER_ENABLED` | `true` | Activa el crawler periódico |
| `CRAWLER_INTERVAL` | `30m` | Intervalo del crawler |
| `CRAWLER_BATCH_SIZE` | `3` | Series procesadas por lote |
| `PROVIDER_ORDER` | `mega,yourupload,streamwish,mp4upload,vidhide,voe` | Prioridad de fuentes |
| `DOWNLOAD_VIDEOS` | `false` | Descarga automática de vídeos |
| `VIDEO_DOWNLOAD_AUTHORIZED` | `false` | Confirmación explícita para descargas automáticas |

## Build y publicación

El workflow `.github/workflows/docker.yml` comprueba el changelog, compila para ARMv7 y publica las etiquetas `0.1.9` y `latest` en Docker Hub. El repositorio necesita los secretos `DOCKERHUB_USERNAME`, `DOCKERHUB_TOKEN` y `PORTAINER_WEBHOOK_URL`.

Build local con Docker Buildx:

```sh
./build.sh 0.1.9
```

## Diagnóstico de Mega

Los logs con prefijo `[MEGA DEBUG]` permiten revisar cómo se interpretan las URL de Mega sin registrar cookies ni claves de descifrado. La descarga admite tanto `/file/HANDLE#KEY` como enlaces legacy `/file/!HANDLE!KEY`.

## Seguridad y alcance

Este mirror está pensado para uso privado en la red local. No evita DRM, tokens ni controles de acceso de proveedores. El propietario del despliegue debe respetar las condiciones del sitio de origen y tener autorización para conservar o descargar el contenido.
