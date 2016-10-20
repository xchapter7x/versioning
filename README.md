# versioning

will simply get your tags and output a new tag string based on the action given

available actions
- bump_patch
- bump_minor
- bump_major

```
$> go get github.com/xchapter7x/versioning
# if latest tag is v1.0.1
$> versioning bump_patch
v1.0.2

# if latest tag is v1.0.1
$> versioning bump_minor
v1.1.0

# if latest tag is v1.0.1
$> versioning bump_major
v2.0.0
```
