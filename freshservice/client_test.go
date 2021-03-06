package freshservice

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCreateChange(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fsKey := "key"

	fs := NewClient(fsURL, fsKey)

	resJSON := `{
		"status": true,
		"item": {
			"itil_change": {
				"id": 1,
				"display_id": 1,
				"requester_id": 1,
				"owner_id": null,
				"group_id": null,
				"priority": 1,
				"impact": 1,
				"status": 1,
				"risk": 1,
				"change_type": 1,
				"approval_status": 4,
				"deleted": false,
				"subject": "change for support",
				"created_at": "2018-11-01T15:58:49+01:00",
				"updated_at": "2018-11-01T15:58:49+01:00",
				"cc_email": {},
				"planned_start_date": "2018-01-01T01:00:00+01:00",
				"planned_end_date": "2018-01-01T01:00:00+01:00",
				"import_id": null,
				"department_id": null,
				"email_config_id": null,
				"project_id": null,
				"approval_type": null,
				"wf_event_id": null,
				"state_flow_id": null,
				"state_traversal": [
					1
				],
				"status_name": "Open",
				"impact_name": "Low",
				"priority_name": "Low",
				"requester_name": "Hulk",
				"owner_name": null,
				"group_name": null,
				"risk_type": "Low",
				"change_type_name": "Minor",
				"approval_status_name": "Not Requested",
				"description": "change description",
				"assoc_release_id": null,
				"associated_assets": [],
				"attachments": [],
				"notes": [],
				"custom_field_values": {}
			}
		},
		"redirect": null
	}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fsURL+"/itil/changes.json",
		httpmock.NewStringResponder(200, resJSON))

	c := Change{
		Email:            "hulk@outerspace.com",
		Subject:          "change for support",
		DescriptionHTML:  "change description",
		Status:           1,
		Priority:         1,
		ChangeType:       1,
		Risk:             1,
		Impact:           1,
		PlannedStartDate: "2018-01-01T00:00:00.00Z",
		PlannedEndDate:   "2018-01-01T00:00:00.00Z",
	}

	res, err := fs.CreateChange(&c)
	if err != nil {
		t.Error(err)
	}

	var response ItilChange

	err = json.Unmarshal([]byte(resJSON), &response)
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(res.Item, response.Item) != true {
		t.Error(err)
	}
}

func TestUpdateChangeStatus(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fsKey := "key"

	fs := NewClient(fsURL, fsKey)

	changeID := 1
	changeStatus := 2
	resJSON := fmt.Sprintf(`{
		"status": true,
		"item": {
			"itil_change": {
				"id": 1,
				"display_id": 1,
				"requester_id": 1,
				"owner_id": null,
				"group_id": null,
				"priority": 1,
				"impact": 1,
				"status": %d,
				"risk": 1,
				"change_type": 1,
				"approval_status": 4,
				"deleted": false,
				"subject": "change for support",
				"created_at": "2018-11-01T15:58:49+01:00",
				"updated_at": "2018-11-01T15:58:49+01:00",
				"cc_email": {},
				"planned_start_date": "2018-01-01T01:00:00+01:00",
				"planned_end_date": "2018-01-01T01:00:00+01:00",
				"import_id": null,
				"department_id": null,
				"email_config_id": null,
				"project_id": null,
				"approval_type": null,
				"wf_event_id": null,
				"state_flow_id": null,
				"state_traversal": [
					1
				],
				"status_name": "Open",
				"impact_name": "Low",
				"priority_name": "Low",
				"requester_name": "Hulk",
				"owner_name": null,
				"group_name": null,
				"risk_type": "Low",
				"change_type_name": "Minor",
				"approval_status_name": "Not Requested",
				"description": "change description",
				"assoc_release_id": null,
				"associated_assets": [],
				"attachments": [],
				"notes": [],
				"custom_field_values": {}
			}
		},
		"redirect": null
	}`, changeStatus)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", fsURL+"/itil/changes/"+strconv.Itoa(changeID)+".json",
		httpmock.NewStringResponder(200, resJSON))

	res, err := fs.UpdateChangeStatus(changeID, changeStatus)
	if err != nil {
		t.Error(err)
	}

	if changeStatus != res.Item.ItilChange.Status {
		t.Error(err)
	}
}

func TestAddChangeNote(t *testing.T) {
	fsURL := "https://foo.freshservice.com"
	fsKey := "key"

	fs := NewClient(fsURL, fsKey)

	changeID := 1
	changeNote := "Hi Hulk, Still Angry"
	resJSON := fmt.Sprintf(`{
		"status": true,
		"item": {
			"note": {
				"id": 1,
				"body": "%s",
				"body_html": "<div>%s</div>",
				"to_emails": null,
				"cc_emails": null,
				"deleted": false,
				"notable_type": "Itil::Change",
				"notable_id": 1,
				"user_id": 1,
				"account_id": 1,
				"created_at": "2018-11-16T19:59:33+01:00",
				"updated_at": "2018-11-16T19:59:33+01:00",
				"header_info": null
			}
		},
		"redirect": null
	}`, changeNote, changeNote)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fsURL+"/itil/changes/"+strconv.Itoa(changeID)+"/notes.json",
		httpmock.NewStringResponder(200, resJSON))

	res, err := fs.AddChangeNote(changeID, changeNote)
	if err != nil {
		t.Error(err)
	}

	if (changeNote != res.Item.Note.Body) && ("<div>"+changeNote+"<div>" != res.Item.Note.BodyHTML) {
		t.Error(err)
	}
}
