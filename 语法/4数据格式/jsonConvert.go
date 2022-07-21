package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	TestJson2Map()
}

func TestJson2Map() {
	jsonStr := `
		{
		  "itemName": "test_item10",
		  "conditions": [
		    {
		      "field": "createtime",
		      "op": ">=",
		      "value": "2020-07-17 00:00:00"
		    },
		    {
		      "field": "createtime",
		      "op": "<",
		      "value": "2020-07-18 00:00:00"
		    }
		  ],
		  "child_filter": {
		    "conditions": [
		      {
		        "field": "CampaignFormID",
		        "op": "=",
		        "value": "172"
		      },
		      {
		        "field": "CampaignFormName",
		        "op": "=",
		        "value": "2323"
		      }
		    ],
		    "joint": "OR"
		  },
		  "joint": "AND",
		  "order_by": {
		    "field": "xxx",
		    "order": "desc"
		  }
		}
    `
	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapResult)
}
