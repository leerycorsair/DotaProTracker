package postgres

const INITIAL_LOGIN string = "host=localhost port=5432 user=initial_login password=initial_login dbname=dotaprotracker sslmode=disable"

const DEFAULT_USER string = "host=localhost port=5432 user=default_user password=default_user dbname=dotaprotracker sslmode=disable"

const MODERATOR string = "host=localhost port=5432 user=moderator password=moderator dbname=dotaprotracker sslmode=disable"

const ADMINISTRATOR string = "host=localhost port=5432 user=administrator password=administrator dbname=dotaprotracker sslmode=disable"

const TEST string = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

var CONNECTIONS_POOL [3]string = [3]string{DEFAULT_USER, MODERATOR, ADMINISTRATOR}
