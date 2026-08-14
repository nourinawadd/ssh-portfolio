# web

The HTTP side of `nourin.is-a.dev`. The SSH TUI (this repo's Go app, port 2323) is a separate
process and is **not** affected by anything in here.

| file | what it is |
|---|---|
| `index.html` | the current holding page — black screen, serif `UNDER CONSTRUCTION`, sparse glitch |
| `index.full.html` | the real landing page, captured from the live server before it was replaced |

## How it's served

`nourin.is-a.dev` has a single A record → `89.168.117.165` (Oracle VM), registered via
[is-a-dev/register](https://github.com/is-a-dev/register) — `domains/nourin.json`, not this repo.
nginx 1.18 on that box serves a plain static `index.html`. No build step, no CI, no framework.

Until now that file existed **only on the VM**. `index.full.html` is the first off-server copy.

## Deploy

Find the real docroot first (`/var/www/html` is the Ubuntu default, but confirm):

```bash
ssh ubuntu@89.168.117.165 "grep -rn 'root ' /etc/nginx/sites-enabled/"
```

Then upload and swap, keeping a copy of whatever was live:

```bash
scp web/index.html ubuntu@89.168.117.165:/tmp/index.html
ssh ubuntu@89.168.117.165
  sudo cp /var/www/html/index.html /var/www/html/index.full.html
  sudo mv /tmp/index.html /var/www/html/index.html
  sudo chown root:root /var/www/html/index.html
  sudo chmod 644 /var/www/html/index.html
```

No `nginx -s reload` — it's a static file, picked up on the next request.

## Restore the real site

```bash
sudo cp /var/www/html/index.full.html /var/www/html/index.html
```

If the VM is ever lost, `web/index.full.html` in this repo is the fallback.

## Check it took

```bash
curl -sI https://nourin.is-a.dev
```

`Content-Length` should be ~4.5 KB (the old page was 8836 bytes) and `Last-Modified` should be
today. In a browser use a hard refresh — nginx sends `ETag`/`Last-Modified`, so a soft refresh can
still show the cached old page.
