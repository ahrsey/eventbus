# Eventbus

# Notes
[sqlite3](https://gosamples.dev/sqlite-intro/)
For sqlite3 implementation copy [save](../save/main.go) structure

# TODO
// TODO: Add sqlite3 for storing events incase of crashes?
// TODO: Add subscriber to the insert into db action which will emit the events,
// something along those lines, the idea is to use the pub/sub pattern here and
// try not to block on posting to the api
// TODO: Access events data in subscribers
// TODO: Cleanup https://github.com/Pungyeon/clean-go-article
// TODO: Queue
// TODO: Update handler signature to be able to work with more complex types
// TODO: Write tests
// TODO: Move logger out into reusable lib
// TODO: Update to use bus to do things itself
