package search

const (
	RankSearchIndex = "rank"
)

const (
	RankIndexMapping = `
{
  "mappings": {
    "properties": {
      "Placement": {
        "type": "integer"
      },
      "Login": {
        "type": "keyword"
      },
      "Id": {
        "type": "long"
      },
      "AvatarURL": {
        "type": "text",
        "index": false
      },
      "HTMLURL": {
        "type": "text",
        "index": false
      },
      "Location": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "Company": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "Score": {
        "type": "integer"
      },
      "Level": {
        "type": "keyword"
      },
      "LocationClassification": {
        "type": "keyword"
      },
      "Tech": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}
`
)
