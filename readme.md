# Flag Fairway

Work in progress

Flag Fairway is ment to be an ultra light weight feature flag service for containers. Using a KV database (Badger) and a high performant language (Go) the product is designed to be fast, reliable, and easy to use feature flag configuration tool behind the firewall of a system (ment to be internal at the moment)

Flag Fairway's UI will be built in Svelte, and hosted as a static site from within the main Go app itself.

what this is intended for

- having feature flags (server side) for backend applications

What this is lacking at the momment
- Security (please keep this behind thne firewall or add authentication)
- scaleablity (this is designed for small to medium size projects at the moment)