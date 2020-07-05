package event_service

// TODO:
// I would love to unit test handleFuncs!! Although, request bodies when inserted and consumed in a reader (such is the one used to build requests) escapes the given json.
// This will accumulate several "\" characters
// Spent some time trying to change this on the test side (don't think that adding logic in the handler itself, to support test cases, is a good practice)
