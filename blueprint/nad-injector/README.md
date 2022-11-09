# nad-injector

## Description
nad-injector controller

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] nad-injector`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree nad-injector`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init nad-injector
kpt live apply nad-injector --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
