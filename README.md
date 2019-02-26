# mkr-services-delete

```bash
mkr-services-delete $(mkr services | jq -r '.[].name | select(startswith("build"))' | tr '\n' ' ')
```
