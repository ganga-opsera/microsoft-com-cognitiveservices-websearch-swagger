package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/web-search-client/mcp-server/config"
	"github.com/web-search-client/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Web_searchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["answerCount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("answerCount=%v", val))
		}
		if val, ok := args["cc"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cc=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["freshness"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("freshness=%v", val))
		}
		if val, ok := args["mkt"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mkt=%v", val))
		}
		if val, ok := args["offset"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("offset=%v", val))
		}
		if val, ok := args["promote"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("promote=%v", val))
		}
		if val, ok := args["q"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("q=%v", val))
		}
		if val, ok := args["responseFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("responseFilter=%v", val))
		}
		if val, ok := args["safeSearch"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("safeSearch=%v", val))
		}
		if val, ok := args["setLang"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("setLang=%v", val))
		}
		if val, ok := args["textDecorations"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("textDecorations=%v", val))
		}
		if val, ok := args["textFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("textFormat=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/search%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Ocp-Apim-Subscription-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-BingApis-SDK"]; ok {
			req.Header.Set("X-BingApis-SDK", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Accept-Language"]; ok {
			req.Header.Set("Accept-Language", fmt.Sprintf("%v", val))
		}
		if val, ok := args["Pragma"]; ok {
			req.Header.Set("Pragma", fmt.Sprintf("%v", val))
		}
		if val, ok := args["User-Agent"]; ok {
			req.Header.Set("User-Agent", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-MSEdge-ClientID"]; ok {
			req.Header.Set("X-MSEdge-ClientID", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-MSEdge-ClientIP"]; ok {
			req.Header.Set("X-MSEdge-ClientIP", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Search-Location"]; ok {
			req.Header.Set("X-Search-Location", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.SearchResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateWeb_searchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_search",
		mcp.WithDescription("The Web Search API lets you send a search query to Bing and get back search results that include links to webpages, images, and more."),
		mcp.WithString("X-BingApis-SDK", mcp.Required(), mcp.Description("Activate swagger compliance")),
		mcp.WithString("Accept", mcp.Description("The default media type is application/json. To specify that the response use [JSON-LD](http://json-ld.org/), set the Accept header to application/ld+json.")),
		mcp.WithString("Accept-Language", mcp.Description("A comma-delimited list of one or more languages to use for user interface strings. The list is in decreasing order of preference. For additional information, including expected format, see [RFC2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). This header and the setLang query parameter are mutually exclusive; do not specify both. If you set this header, you must also specify the cc query parameter. Bing will use the first supported language it finds from the list, and combine that language with the cc parameter value to determine the market to return results for. If the list does not include a supported language, Bing will find the closest language and market that supports the request, and may use an aggregated or default market for the results instead of a specified one. You should use this header and the cc query parameter only if you specify multiple languages; otherwise, you should use the mkt and setLang query parameters. A user interface string is a string that's used as a label in a user interface. There are very few user interface strings in the JSON response objects. Any links in the response objects to Bing.com properties will apply the specified language.")),
		mcp.WithString("Pragma", mcp.Description("By default, Bing returns cached content, if available. To prevent Bing from returning cached content, set the Pragma header to no-cache (for example, Pragma: no-cache).")),
		mcp.WithString("User-Agent", mcp.Description("The user agent originating the request. Bing uses the user agent to provide mobile users with an optimized experience. Although optional, you are strongly encouraged to always specify this header. The user-agent should be the same string that any commonly used browser would send. For information about user agents, see [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).")),
		mcp.WithString("X-MSEdge-ClientID", mcp.Description("Bing uses this header to provide users with consistent behavior across Bing API calls. Bing often flights new features and improvements, and it uses the client ID as a key for assigning traffic on different flights. If you do not use the same client ID for a user across multiple requests, then Bing may assign the user to multiple conflicting flights. Being assigned to multiple conflicting flights can lead to an inconsistent user experience. For example, if the second request has a different flight assignment than the first, the experience may be unexpected. Also, Bing can use the client ID to tailor web results to that client ID’s search history, providing a richer experience for the user. Bing also uses this header to help improve result rankings by analyzing the activity generated by a client ID. The relevance improvements help with better quality of results delivered by Bing APIs and in turn enables higher click-through rates for the API consumer. IMPORTANT: Although optional, you should consider this header required. Persisting the client ID across multiple requests for the same end user and device combination enables 1) the API consumer to receive a consistent user experience, and 2) higher click-through rates via better quality of results from the Bing APIs. Each user that uses your application on the device must have a unique, Bing generated client ID. If you do not include this header in the request, Bing generates an ID and returns it in the X-MSEdge-ClientID response header. The only time that you should NOT include this header in a request is the first time the user uses your app on that device. Use the client ID for each Bing API request that your app makes for this user on the device. Persist the client ID. To persist the ID in a browser app, use a persistent HTTP cookie to ensure the ID is used across all sessions. Do not use a session cookie. For other apps such as mobile apps, use the device's persistent storage to persist the ID. The next time the user uses your app on that device, get the client ID that you persisted. Bing responses may or may not include this header. If the response includes this header, capture the client ID and use it for all subsequent Bing requests for the user on that device. If you include the X-MSEdge-ClientID, you must not include cookies in the request.")),
		mcp.WithString("X-MSEdge-ClientIP", mcp.Description("The IPv4 or IPv6 address of the client device. The IP address is used to discover the user's location. Bing uses the location information to determine safe search behavior. Although optional, you are encouraged to always specify this header and the X-Search-Location header. Do not obfuscate the address (for example, by changing the last octet to 0). Obfuscating the address results in the location not being anywhere near the device's actual location, which may result in Bing serving erroneous results.")),
		mcp.WithString("X-Search-Location", mcp.Description("A semicolon-delimited list of key/value pairs that describe the client's geographical location. Bing uses the location information to determine safe search behavior and to return relevant local content. Specify the key/value pair as <key>:<value>. The following are the keys that you use to specify the user's location. lat (required): The latitude of the client's location, in degrees. The latitude must be greater than or equal to -90.0 and less than or equal to +90.0. Negative values indicate southern latitudes and positive values indicate northern latitudes. long (required): The longitude of the client's location, in degrees. The longitude must be greater than or equal to -180.0 and less than or equal to +180.0. Negative values indicate western longitudes and positive values indicate eastern longitudes. re (required): The radius, in meters, which specifies the horizontal accuracy of the coordinates. Pass the value returned by the device's location service. Typical values might be 22m for GPS/Wi-Fi, 380m for cell tower triangulation, and 18,000m for reverse IP lookup. ts (optional): The UTC UNIX timestamp of when the client was at the location. (The UNIX timestamp is the number of seconds since January 1, 1970.) head (optional): The client's relative heading or direction of travel. Specify the direction of travel as degrees from 0 through 360, counting clockwise relative to true north. Specify this key only if the sp key is nonzero. sp (optional): The horizontal velocity (speed), in meters per second, that the client device is traveling. alt (optional): The altitude of the client device, in meters. are (optional): The radius, in meters, that specifies the vertical accuracy of the coordinates. Specify this key only if you specify the alt key. Although many of the keys are optional, the more information that you provide, the more accurate the location results are. Although optional, you are encouraged to always specify the user's geographical location. Providing the location is especially important if the client's IP address does not accurately reflect the user's physical location (for example, if the client uses VPN). For optimal results, you should include this header and the X-MSEdge-ClientIP header, but at a minimum, you should include this header.")),
		mcp.WithNumber("answerCount", mcp.Description("The number of answers that you want the response to include. The answers that Bing returns are based on ranking. For example, if Bing returns webpages, images, videos, and relatedSearches for a request and you set this parameter to two (2), the response includes webpages and images.If you included the responseFilter query parameter in the same request and set it to webpages and news, the response would include only webpages.")),
		mcp.WithString("cc", mcp.Description("A 2-character country code of the country where the results come from. This API supports only the United States market. If you specify this query parameter, it must be set to us. If you set this parameter, you must also specify the Accept-Language header. Bing uses the first supported language it finds from the languages list, and combine that language with the country code that you specify to determine the market to return results for. If the languages list does not include a supported language, Bing finds the closest language and market that supports the request, or it may use an aggregated or default market for the results instead of a specified one. You should use this query parameter and the Accept-Language query parameter only if you specify multiple languages; otherwise, you should use the mkt and setLang query parameters. This parameter and the mkt query parameter are mutually exclusive—do not specify both.")),
		mcp.WithNumber("count", mcp.Description("The number of search results to return in the response. The default is 10 and the maximum value is 50. The actual number delivered may be less than requested.Use this parameter along with the offset parameter to page results.For example, if your user interface displays 10 search results per page, set count to 10 and offset to 0 to get the first page of results. For each subsequent page, increment offset by 10 (for example, 0, 10, 20). It is possible for multiple pages to include some overlap in results.")),
		mcp.WithString("freshness", mcp.Description("Filter search results by the following age values: Day—Return webpages that Bing discovered within the last 24 hours. Week—Return webpages that Bing discovered within the last 7 days. Month—Return webpages that discovered within the last 30 days. This filter applies only to webpage results and not to the other results such as news and images.")),
		mcp.WithString("mkt", mcp.Description("The market where the results come from. Typically, mkt is the country where the user is making the request from. However, it could be a different country if the user is not located in a country where Bing delivers results. The market must be in the form <language code>-<country code>. For example, en-US. The string is case insensitive. If known, you are encouraged to always specify the market. Specifying the market helps Bing route the request and return an appropriate and optimal response. If you specify a market that is not listed in Market Codes, Bing uses a best fit market code based on an internal mapping that is subject to change. This parameter and the cc query parameter are mutually exclusive—do not specify both.")),
		mcp.WithNumber("offset", mcp.Description("The zero-based offset that indicates the number of search results to skip before returning results. The default is 0. The offset should be less than (totalEstimatedMatches - count). Use this parameter along with the count parameter to page results. For example, if your user interface displays 10 search results per page, set count to 10 and offset to 0 to get the first page of results. For each subsequent page, increment offset by 10 (for example, 0, 10, 20). it is possible for multiple pages to include some overlap in results.")),
		mcp.WithArray("promote", mcp.Description("A comma-delimited list of answers that you want the response to include regardless of their ranking. For example, if you set answerCount) to two (2) so Bing returns the top two ranked answers, but you also want the response to include news, you'd set promote to news. If the top ranked answers are webpages, images, videos, and relatedSearches, the response includes webpages and images because news is not a ranked answer. But if you set promote to video, Bing would promote the video answer into the response and return webpages, images, and videos. The answers that you want to promote do not count against the answerCount limit. For example, if the ranked answers are news, images, and videos, and you set answerCount to 1 and promote to news, the response contains news and images. Or, if the ranked answers are videos, images, and news, the response contains videos and news. Possible values are Computation, Images, News, RelatedSearches, SpellSuggestions, TimeZone, Videos, Webpages. Use only if you specify answerCount.")),
		mcp.WithString("q", mcp.Required(), mcp.Description("The user's search query term. The term may not be empty. The term may contain Bing Advanced Operators. For example, to limit results to a specific domain, use the site: operator.")),
		mcp.WithArray("responseFilter", mcp.Description("A comma-delimited list of answers to include in the response. If you do not specify this parameter, the response includes all search answers for which there's relevant data. Possible filter values are Computation, Images, News, RelatedSearches, SpellSuggestions, TimeZone, Videos, Webpages. Although you may use this filter to get a single answer, you should instead use the answer-specific endpoint in order to get richer results. For example, to receive only images, send the request to one of the Image Search API endpoints. The RelatedSearches and SpellSuggestions answers do not support a separate endpoint like the Image Search API does (only the Web Search API returns them). To include answers that would otherwise be excluded because of ranking, see the promote query parameter.")),
		mcp.WithString("safeSearch", mcp.Description("A filter used to filter adult content. Off: Return webpages with adult text, images, or videos. Moderate: Return webpages with adult text, but not adult images or videos. Strict: Do not return webpages with adult text, images, or videos. The default is Moderate. If the request comes from a market that Bing's adult policy requires that safeSearch is set to Strict, Bing ignores the safeSearch value and uses Strict. If you use the site: query operator, there is the chance that the response may contain adult content regardless of what the safeSearch query parameter is set to. Use site: only if you are aware of the content on the site and your scenario supports the possibility of adult content.")),
		mcp.WithString("setLang", mcp.Description("The language to use for user interface strings. Specify the language using the ISO 639-1 2-letter language code. For example, the language code for English is EN. The default is EN (English). Although optional, you should always specify the language. Typically, you set setLang to the same language specified by mkt unless the user wants the user interface strings displayed in a different language. This parameter and the Accept-Language header are mutually exclusive; do not specify both. A user interface string is a string that's used as a label in a user interface. There are few user interface strings in the JSON response objects. Also, any links to Bing.com properties in the response objects apply the specified language.")),
		mcp.WithBoolean("textDecorations", mcp.Description("A Boolean value that determines whether display strings should contain decoration markers such as hit highlighting characters. If true, the strings may include markers. The default is false. To specify whether to use Unicode characters or HTML tags as the markers, see the textFormat query parameter.")),
		mcp.WithString("textFormat", mcp.Description("The type of markers to use for text decorations (see the textDecorations query parameter). Possible values are Raw—Use Unicode characters to mark content that needs special formatting. The Unicode characters are in the range E000 through E019. For example, Bing uses E000 and E001 to mark the beginning and end of query terms for hit highlighting. HTML—Use HTML tags to mark content that needs special formatting. For example, use <b> tags to highlight query terms in display strings. The default is Raw. For display strings that contain escapable HTML characters such as <, >, and &, if textFormat is set to HTML, Bing escapes the characters as appropriate (for example, < is escaped to &lt;).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Web_searchHandler(cfg),
	}
}
