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

Either no warning should be produced or Renovate should correctly identify the package when no `registryUrls` are defined in the package rule.

Note that the following warning is produced when no `registryUrls` are defined:

```json
{
  "currentValue": "v2.23.4",
  "datasource": "go",
  "depName": "github.com/onsi/ginkgo/v2",
  "depType": "require",
  "managerData": {
    "lineNumber": 7,
    "multiLine": true
  },
  "packageName": "github.com/onsi/ginkgo/v2",
  "versioning": "semver",
  "warnings": [
    {
      "topic": "github.com/onsi/ginkgo/v2",
      "message": "Failed to look up go package github.com/onsi/ginkgo/v2"
    }
  ],
  "updates": []
}
```

## Link to the Renovate issue or Discussion

https://github.com/renovatebot/renovate/discussions/35963
