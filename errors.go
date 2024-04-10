package main

func checkStatus(data TrafficData) string {
	var msg string

	switch data.Status {
	case "INVALID_REQUEST":
		msg = "Please provide a valid request"

	case "MAX_ELEMENTS_EXCEEDED":
		msg = "Sorry, maximum elements per query limit exceeded!"

	case "MAX_DIMENSIONS_EXCEEDED":
		msg = "Sorry, maximum dimensions per query limit exceeded!"

	case "OVER_QUERY_LIMIT":
		msg = "too many requests from this your application!"

	case "REQUEST_DENIED":
		msg = "service denied! app can no longer access Distance Matrix service"

	case "UNKNOWN_ERROR":
		msg = "Sorry! server failed to respond"

	default:
		msg = ""
	}

	return msg
}

func checkElementStatus(data TrafficData) string {
	var msg string

	switch data.Rows[0].Elements[0].Status {
	case "NOT_FOUND":
		msg = "Please provide valid origin/destination!"

	case "ZERO_RESULTS":
		msg = "Route not found between origin and destination!"

	case "MAX_ROUTE_LENGTH_EXCEEDED":
		msg = "request route too long to be processed"

	default:
		msg = ""
	}
	return msg
}
