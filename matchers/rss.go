package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/shivani-1521/RSS-feeds-Matcher/search"
)


type (
	item struct{
		XMLName			xml.Name	'xml:"item"'
		Title			string		'xml:"title"'
		Description		string		'xml:"description"'
		Link			string		'xml:"link"'
		Pubdate			string		'xml:"pubDate"'
		GUID			string		'xml:"guid"'
		GeoRssPoint		string		'xml:"georss:point"'
	}

	image struct{
		XMLName			xml.Name	'xml:"image"'
		Title			string		'xml:"title"'
		Link			string		'xml:"link"'
		URL				string		'xml:"url"'

	}

	channel struct{
		XMLName			xml.Name	'xml:"channel"'
		Title			string		'xml:"title"'
		Description		string		'xml:"description"'
		Link			string		'xml:"link"'
		Pubdate			string		'xml:"pubDate"'
		LastBuildDate	string		'xml:"lastBuildDate"'
		TTL             string		'xml:"ttl"'
		Language		string		'xml:"language"'
		ManagingEditor	string		'xml:"managingEditor"'
		WebMaster		string		'xml:"webMaster"'
		Image			Image       'xml:"image"'
		Item            []item      'xml:"link"'
	}

	rssDocument struct{
		XMLName			xml.Name	'xml:"rss"'
		Channel			channel 	'xml:"channel"'
	}
)


