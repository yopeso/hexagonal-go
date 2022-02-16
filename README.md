# Hexagonal Go Workshop #

[Presentation](https://www.google.com)

## Structure ##

* `src/shortener` - "business logic", defines the adapters (interfaces) and contains the logic to use them (service)
* `src/repository`- implementations for different repositories that implement the required interface
* `src/serializer`- implementations for different serializers that implement the interface
* `src/metrics` - counters for prometheus metrics (might also be a part of `shortener` as they can be business logic metrics)
* `src/api` - api handlers that create handlers for exposed endpoints to create/use URLs

Prerequisites:

* Docker
* VSCode with REST Client extension required to use the `test.http` file to interact with the API
