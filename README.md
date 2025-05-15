# 35963

## Current behavior

The following warning is being produced:

```shell
WARN: Custom registries are not allowed for this datasource and will be ignored
{
  "datasource": "go"
  "registryUrls": [
    "https://github.com/onsi/ginkgo"
  ]
  "defaultRegistryUrls": null
}
```

## Expected behavior

No warning should be produced.

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/discussions/35963
