### a vue-router history mode deploy serve
> fallback to `index.html`

### config file

```json
{
  "port": 8089,
  "path": "./dist",
  "routes": [
    {"src": "^/js/(.*)", "dest": "/js/$1"},
    {"src": "^/css/(.*)", "dest": "/css/$1"},
    {"src": "^/img/(.*)", "dest": "/img/$1"},
    {"src": "^/wb-assets/(.*)", "dest": "/wb-assets/$1"},
    {"src": "^/service-worker.js", "dest": "/service-worker.js"},
    {"src": "^/manifest.json", "dest": "/manifest.json"},
    {"src": ".*", "dest": "/index.html"}
  ]
}
```
* port: a net port
* path: a dist dir path