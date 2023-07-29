# Reusable

The purpose of `interface`s is to create functions that can be used on multiple types. All structs that implement these need to change if the interface ever changes and are not expected to happen often.

## Best practices

* Make the interface in it's own file describing what it is supposed to do
* Usually put all implimentations of that interface in child packages of the interface
* In test files add `var _ <interface name> = &<struct name>{}` to make sure the struct fulfils the instance
