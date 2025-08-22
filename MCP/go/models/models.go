package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Answer represents the Answer schema from the OpenAPI specification
type Answer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// StructuredValue represents the StructuredValue schema from the OpenAPI specification
type StructuredValue struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
}

// Places represents the Places schema from the OpenAPI specification
type Places struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
}

// WebWebAnswer represents the WebWebAnswer schema from the OpenAPI specification
type WebWebAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// RankingRankingResponse represents the RankingRankingResponse schema from the OpenAPI specification
type RankingRankingResponse struct {
	Sidebar RankingRankingGroup `json:"sidebar,omitempty"` // Defines a search results group, such as mainline.
	Mainline RankingRankingGroup `json:"mainline,omitempty"` // Defines a search results group, such as mainline.
	Pole RankingRankingGroup `json:"pole,omitempty"` // Defines a search results group, such as mainline.
}

// RelatedSearchesRelatedSearchAnswer represents the RelatedSearchesRelatedSearchAnswer schema from the OpenAPI specification
type RelatedSearchesRelatedSearchAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// News represents the News schema from the OpenAPI specification
type News struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Query represents the Query schema from the OpenAPI specification
type Query struct {
	Displaytext string `json:"displayText,omitempty"` // The display version of the query term. This version of the query term may contain special characters that highlight the search term found in the query string. The string contains the highlighting characters only if the query enabled hit highlighting
	Searchlink string `json:"searchLink,omitempty"`
	Text string `json:"text"` // The query string. Use this string as the query term in a new search request.
	Thumbnail ImageObject `json:"thumbnail,omitempty"` // Defines an image
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL that takes the user to the Bing search results page for the query.Only related search results include this field.
}

// Response represents the Response schema from the OpenAPI specification
type Response struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
}

// Article represents the Article schema from the OpenAPI specification
type Article struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// TimeZoneTimeZoneInformation represents the TimeZoneTimeZoneInformation schema from the OpenAPI specification
type TimeZoneTimeZoneInformation struct {
	Location string `json:"location"` // The name of the geographical location.For example, County; City; City, State; City, State, Country; or Time Zone.
	Time string `json:"time"` // The data and time specified in the form, YYYY-MM-DDThh;mm:ss.ssssssZ.
	Utcoffset string `json:"utcOffset"` // The offset from UTC. For example, UTC-7.
}

// Computation represents the Computation schema from the OpenAPI specification
type Computation struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
}

// SearchResultsAnswer represents the SearchResultsAnswer schema from the OpenAPI specification
type SearchResultsAnswer struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
}

// WebMetaTag represents the WebMetaTag schema from the OpenAPI specification
type WebMetaTag struct {
	Content string `json:"content,omitempty"` // The name of the metadata.
	Name string `json:"name,omitempty"` // The metadata.
}

// Intangible represents the Intangible schema from the OpenAPI specification
type Intangible struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
}

// Videos represents the Videos schema from the OpenAPI specification
type Videos struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// VideoObject represents the VideoObject schema from the OpenAPI specification
type VideoObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Hostpageurl string `json:"hostPageUrl,omitempty"` // URL of the page that hosts the media object.
}

// RankingRankingItem represents the RankingRankingItem schema from the OpenAPI specification
type RankingRankingItem struct {
	Resultindex int `json:"resultIndex,omitempty"` // A zero-based index of the item in the answer.If the item does not include this field, display all items in the answer. For example, display all news articles in the News answer.
	Screenshotindex int `json:"screenshotIndex,omitempty"`
	Textualindex int `json:"textualIndex,omitempty"`
	Value Identifiable `json:"value,omitempty"` // Defines the identity of a resource.
	Answertype string `json:"answerType"` // The answer that contains the item to display. Use the type to find the answer in the SearchResponse object. The type is the name of a SearchResponse field.
	Htmlindex int `json:"htmlIndex,omitempty"`
}

// WebWebGrouping represents the WebWebGrouping schema from the OpenAPI specification
type WebWebGrouping struct {
	Webpages []WebPage `json:"webPages"`
	TypeField string `json:"_type"`
}

// ResponseBase represents the ResponseBase schema from the OpenAPI specification
type ResponseBase struct {
	TypeField string `json:"_type"`
}

// SearchResponse represents the SearchResponse schema from the OpenAPI specification
type SearchResponse struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// MediaObject represents the MediaObject schema from the OpenAPI specification
type MediaObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// ImageObject represents the ImageObject schema from the OpenAPI specification
type ImageObject struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Height int `json:"height,omitempty"` // The height of the source media object, in pixels.
	Hostpageurl string `json:"hostPageUrl,omitempty"` // URL of the page that hosts the media object.
	Width int `json:"width,omitempty"` // The width of the source media object, in pixels.
	Contenturl string `json:"contentUrl,omitempty"` // Original URL to retrieve the source (file) for the media object (e.g the source URL for the image).
}

// QueryContext represents the QueryContext schema from the OpenAPI specification
type QueryContext struct {
	Originalquery string `json:"originalQuery"` // The query string as specified in the request.
	Adultintent bool `json:"adultIntent,omitempty"` // A Boolean value that indicates whether the specified query has adult intent. The value is true if the query has adult intent; otherwise, false.
	Alterationoverridequery string `json:"alterationOverrideQuery,omitempty"` // The query string to use to force Bing to use the original string. For example, if the query string is "saling downwind", the override query string will be "+saling downwind". Remember to encode the query string which results in "%2Bsaling+downwind". This field is included only if the original query string contains a spelling mistake.
	Alteredquery string `json:"alteredQuery,omitempty"` // The query string used by Bing to perform the query. Bing uses the altered query string if the original query string contained spelling mistakes. For example, if the query string is "saling downwind", the altered query string will be "sailing downwind". This field is included only if the original query string contains a spelling mistake.
	Askuserforlocation bool `json:"askUserForLocation,omitempty"` // A Boolean value that indicates whether Bing requires the user's location to provide accurate results. If you specified the user's location by using the X-MSEdge-ClientIP and X-Search-Location headers, you can ignore this field. For location aware queries, such as "today's weather" or "restaurants near me" that need the user's location to provide accurate results, this field is set to true. For location aware queries that include the location (for example, "Seattle weather"), this field is set to false. This field is also set to false for queries that are not location aware, such as "best sellers".
	Istransactional bool `json:"isTransactional,omitempty"`
}

// NewsArticle represents the NewsArticle schema from the OpenAPI specification
type NewsArticle struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
	Wordcount int `json:"wordCount,omitempty"` // The number of words in the text of the Article.
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Moredetails string `json:"moreDetails,omitempty"` // A description that provides additional information about the error.
	Parameter string `json:"parameter,omitempty"` // The parameter in the request that caused the error.
	Subcode string `json:"subCode,omitempty"` // The error code that further helps to identify the error.
	Value string `json:"value,omitempty"` // The parameter's value in the request that was not valid.
	Code string `json:"code"` // The error code that identifies the category of error.
	Message string `json:"message"` // A description of the error.
}

// TimeZone represents the TimeZone schema from the OpenAPI specification
type TimeZone struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// CreativeWork represents the CreativeWork schema from the OpenAPI specification
type CreativeWork struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
}

// RankingRankingGroup represents the RankingRankingGroup schema from the OpenAPI specification
type RankingRankingGroup struct {
	Items []RankingRankingItem `json:"items"` // A list of search result items to display in the group.
}

// WebPage represents the WebPage schema from the OpenAPI specification
type WebPage struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Description string `json:"description,omitempty"` // A short description of the item.
	Image ImageObject `json:"image,omitempty"` // Defines an image
	Name string `json:"name,omitempty"` // The name of the thing represented by this object.
	Url string `json:"url,omitempty"` // The URL to get more information about the thing represented by this object.
	Bingid string `json:"bingId,omitempty"` // An ID that uniquely identifies this item.
	Provider []Thing `json:"provider,omitempty"` // The source of the creative work.
	Text string `json:"text,omitempty"`
	Thumbnailurl string `json:"thumbnailUrl,omitempty"` // The URL to a thumbnail of the item.
}

// Images represents the Images schema from the OpenAPI specification
type Images struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Thing represents the Thing schema from the OpenAPI specification
type Thing struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
}

// SpellSuggestions represents the SpellSuggestions schema from the OpenAPI specification
type SpellSuggestions struct {
	TypeField string `json:"_type"`
	Id string `json:"id,omitempty"` // A String identifier.
	Websearchurl string `json:"webSearchUrl,omitempty"` // The URL To Bing's search result for this item.
	Followupqueries []Query `json:"followUpQueries,omitempty"`
	Totalestimatedmatches int64 `json:"totalEstimatedMatches,omitempty"` // The estimated number of webpages that are relevant to the query. Use this number along with the count and offset query parameters to page the results.
	Isfamilyfriendly bool `json:"isFamilyFriendly,omitempty"`
	Querycontext QueryContext `json:"queryContext,omitempty"` // Defines the query context that Bing used for the request.
}

// Identifiable represents the Identifiable schema from the OpenAPI specification
type Identifiable struct {
	TypeField string `json:"_type"`
}
