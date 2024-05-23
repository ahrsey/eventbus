# Eventbus

# Notes
[sqlite3](https://gosamples.dev/sqlite-intro/)
For sqlite3 implementation copy [save](../save/main.go) structure

# TODO
// TODO: Add sqlite3
// TODO: Add subscriber to the insert into db action which will emit the
// events, something along those lines, the idea is to use the pub/sub pattern
// here and try not to block on posting to the api
// TODO: Publish events
// TODO: Filter subscriber triggers
// TODO: Access events data in subscribers
// TODO: When you subscribe it should create an event that is possible to be
// trigger to be used to trigger that subscriber
// TODO: Cleanup, we have three main libs -> pub/sub, database and http, so I think
// it's time to tidy up before things get too messy https://github.com/Pungyeon/clean-go-article
// TODO: Emit event
// TODO: Pick up event and publish that event to required subscribers
// TODO: Subscribe to publisher which creates event groups that can be
// published to
