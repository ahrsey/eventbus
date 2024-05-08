# Eventbus

- publish
- subscribe
- unsubscribe

# Publish
Publish events to event list

# Subscribe
Subscribe to event list with additional event filtering

# Event filtering
Specify which events we'll subscribe to, this should create a new event filter.
Each event filter is sessentially a publisher with scoped events. Each event
can be pushed to multiple event filters.

# Methods
- publish event
- filter events
- subscribe events or filtered events
- trigger subscribers

# Reading
https://medium.com/globant/pub-sub-in-golang-an-introduction-8be4c65eafd4
https://ably.com/blog/pubsub-golang
https://cloud.google.com/pubsub/docs/overview
https://blog.logrocket.com/building-pub-sub-service-go/
