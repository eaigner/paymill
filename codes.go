package paymill

var codes = map[int]string{
	10001: "General undefined response.",
	10002: "Still waiting on something.",
	20000: "General success response.",
	40000: "General problem with data.",
	40100: "Problem with creditcard data.",
	40101: "Problem with cvv.",
	40102: "Card expired or not yet valid.",
	40103: "Limit exceeded.",
	40104: "Card invalid.",
	40105: "expiry date not valid",
	40200: "Problem with bank account data.",
	40300: "Problem with 3d secure data.",
	40301: "currency / amount mismatch",
	40400: "Problem with input data.",
	40401: "Amount too low or zero.",
	40402: "Usage field too long.",
	40403: "Currency not allowed.",
	50000: "General problem with backend.",
	50001: "country blacklisted.",
	50100: "Technical error with credit card.",
	50101: "Error limit exceeded.",
	50102: "Card declined by authorization system.",
	50103: "Manipulation or stolen card.",
	50104: "Card restricted.",
	50105: "Invalid card configuration data.",
	50200: "Technical error with bank account.",
	50201: "Card blacklisted.",
	50300: "Technical error with 3D secure.",
	50400: "Decline because of risk issues.",
	50500: "General timeout.",
	50501: "Timeout on side of the acquirer.",
}
