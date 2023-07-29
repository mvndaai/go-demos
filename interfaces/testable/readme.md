# Testable

To avoid complex setup `interface`s are often created not for usablibity but just to simplify testing. These have a different purpose so treat them differently. These interfaces change with the change of the struct and are expected to happen often

## Best Practices

* Keep the interface in the same file as the `struct`
* Make the interface include only the public functions needed for tests
* If a `New` function for the struct exists make it return the interface instead of a pointer, i.e. `func New<struct>() <interface>`

