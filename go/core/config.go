package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Resend",
			"slug": "resend",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"now": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api.resend.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"add_contact_to_segment_response_success": map[string]any{},
				"add_inbox_thread_labels_response_success": map[string]any{},
				"api_key": map[string]any{},
				"audience": map[string]any{},
				"automation": map[string]any{},
				"automation_run": map[string]any{},
				"automation_run_list_item": map[string]any{},
				"batch_add_suppressions_response_success": map[string]any{},
				"batch_remove_suppressions_response_success": map[string]any{},
				"broadcast": map[string]any{},
				"contact": map[string]any{},
				"contact_import_response_success": map[string]any{},
				"contact_property": map[string]any{},
				"contact_topics_response_success": map[string]any{},
				"create_api_key": map[string]any{},
				"create_batch_email": map[string]any{},
				"create_broadcast_response_success": map[string]any{},
				"create_inbox_draft_response_success": map[string]any{},
				"create_inbox_label_response_success": map[string]any{},
				"domain": map[string]any{},
				"domain_claim": map[string]any{},
				"email": map[string]any{},
				"emails_metric": map[string]any{},
				"event": map[string]any{},
				"inbox": map[string]any{},
				"inbox_agent_response_success": map[string]any{},
				"inbox_draft_list_item": map[string]any{},
				"inbox_draft_response_success": map[string]any{},
				"inbox_label": map[string]any{},
				"inbox_thread": map[string]any{},
				"inbox_thread_email": map[string]any{},
				"inbox_thread_email_response_success": map[string]any{},
				"inbox_thread_response_success": map[string]any{},
				"list_attachment": map[string]any{},
				"list_broadcast_clicked_links_response_success": map[string]any{},
				"list_broadcast_recipients_response_success": map[string]any{},
				"list_contact_segments_response_success": map[string]any{},
				"list_contacts_response_success": map[string]any{},
				"list_webhook_event": map[string]any{},
				"list_webhook_event_attempt": map[string]any{},
				"log": map[string]any{},
				"received_email": map[string]any{},
				"remove_audience_response_success": map[string]any{},
				"remove_broadcast_response_success": map[string]any{},
				"remove_contact_from_segment_response_success": map[string]any{},
				"remove_contact_property_response_success": map[string]any{},
				"remove_contact_response_success": map[string]any{},
				"remove_event": map[string]any{},
				"remove_inbox_draft_response_success": map[string]any{},
				"remove_inbox_label_response_success": map[string]any{},
				"remove_inbox_response_success": map[string]any{},
				"remove_inbox_thread_label_response_success": map[string]any{},
				"remove_inbox_thread_response_success": map[string]any{},
				"remove_segment_response_success": map[string]any{},
				"remove_suppression_response_success": map[string]any{},
				"remove_template_response_success": map[string]any{},
				"remove_topic_response_success": map[string]any{},
				"retrieved_attachment": map[string]any{},
				"revoke_o_auth_grant": map[string]any{},
				"rotate": map[string]any{},
				"segment": map[string]any{},
				"set_inbox_thread_labels_response_success": map[string]any{},
				"suppression": map[string]any{},
				"template": map[string]any{},
				"topic": map[string]any{},
				"update_api_key": map[string]any{},
				"update_broadcast_response_success": map[string]any{},
				"update_contact_property_response_success": map[string]any{},
				"update_contact_response_success": map[string]any{},
				"update_contact_topics_response_success": map[string]any{},
				"update_domain_response_success": map[string]any{},
				"update_email_option": map[string]any{},
				"update_event": map[string]any{},
				"update_inbox_draft_response_success": map[string]any{},
				"update_inbox_label_response_success": map[string]any{},
				"update_inbox_response_success": map[string]any{},
				"update_inbox_thread_response_success": map[string]any{},
				"update_segment_response_success": map[string]any{},
				"update_template_response_success": map[string]any{},
				"update_topic_response_success": map[string]any{},
				"update_webhook": map[string]any{},
				"usage": map[string]any{},
				"webhook": map[string]any{},
				"webhook_event": map[string]any{},
			},
		},
		"entity": map[string]any{
			"add_contact_to_segment_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "contact_id",
						"title": "Contact Id",
						"type": "`$STRING`",
						"short": "The ID of the contact.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
					map[string]any{
						"name": "segment_id",
						"title": "Segment Id",
						"type": "`$STRING`",
						"short": "The ID of the segment.",
					},
				},
				"name": "add_contact_to_segment_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/contacts/{contact_id}/segments/{segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "contact_id",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
								},
								"parts": []any{
									"contacts",
									"{contact_id}",
									"segments",
									"{segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "contact_id",
											"orig": "contact_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"contact_id",
										"segment_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.contact",
							"$.main.kit.entity.segment",
						},
					},
				},
			},
			"add_inbox_thread_labels_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Every label on the thread after the change.",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The IDs of the labels to apply to the thread.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "Type of the response object.",
					},
				},
				"name": "add_inbox_thread_labels_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/labels",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"api_key": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the API key was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the API key.",
					},
					map[string]any{
						"name": "last_used_at",
						"title": "Last Used At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the API key was last used.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the API key.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "api_key",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/api-keys",
								"segments": []any{
									map[string]any{
										"lit": "api-keys",
									},
								},
								"parts": []any{
									"api-keys",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/api-keys/{api_key_id}",
								"segments": []any{
									map[string]any{
										"lit": "api-keys",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"api-keys",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"api_key_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "api_key_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"audience": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date that the object was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the audience.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The name of the audience.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object of the audience.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "audience",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/audiences",
								"segments": []any{
									map[string]any{
										"lit": "audiences",
									},
								},
								"parts": []any{
									"audiences",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/audiences",
								"segments": []any{
									map[string]any{
										"lit": "audiences",
									},
								},
								"parts": []any{
									"audiences",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/audiences/{id}",
								"segments": []any{
									map[string]any{
										"lit": "audiences",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"audiences",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"automation": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "connections",
						"title": "Connections",
						"type": "`$ARRAY`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"short": "The connections between steps in the active version of the automation.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the automation was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the automation.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The name of the automation.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The current status of the automation.",
					},
					map[string]any{
						"name": "steps",
						"title": "Steps",
						"type": "`$ARRAY`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"short": "The steps in the active version of the automation.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "The date and time the automation was last updated.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "automation",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/automations/{automation_id}/duplicate",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "duplicate",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
									"duplicate",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "duplicate",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/automations/{automation_id}/stop",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "stop",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
									"stop",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "stop",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/automations",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
								},
								"parts": []any{
									"automations",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/automations",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
								},
								"parts": []any{
									"automations",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/automations/{automation_id}",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/automations/{automation_id}",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/automations/{automation_id}",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"automation_run": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "completed_at",
						"title": "Completed At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the run completed.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the run was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the automation run.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "started_at",
						"title": "Started At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the run started.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The current status of the automation run.",
					},
					map[string]any{
						"name": "steps",
						"title": "Steps",
						"type": "`$ARRAY`",
						"short": "The steps executed in this run, sorted in graph order.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "automation_run",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/automations/{automation_id}/runs/{run_id}",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "automation_id",
									},
									map[string]any{
										"lit": "runs",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"automations",
									"{automation_id}",
									"runs",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"run_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "automation_id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "id",
											"orig": "run_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"automation_id",
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.automation",
						},
					},
				},
			},
			"automation_run_list_item": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "completed_at",
						"title": "Completed At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the run completed.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the run was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the automation run.",
					},
					map[string]any{
						"name": "started_at",
						"title": "Started At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the run started.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The current status of the automation run.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "automation_run_list_item",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/automations/{automation_id}/runs",
								"segments": []any{
									map[string]any{
										"lit": "automations",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "runs",
									},
								},
								"parts": []any{
									"automations",
									"{id}",
									"runs",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"automation_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "automation_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "runs",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"batch_add_suppressions_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "emails",
						"title": "Emails",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Email addresses to suppress.",
					},
				},
				"name": "batch_add_suppressions_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/suppressions/batch/add",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"lit": "batch",
									},
									map[string]any{
										"lit": "add",
									},
								},
								"parts": []any{
									"suppressions",
									"batch",
									"add",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"batch_remove_suppressions_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"short": "Array containing the removed suppressions.",
					},
					map[string]any{
						"name": "emails",
						"title": "Emails",
						"type": "`$ARRAY`",
						"short": "Email addresses to remove from the suppression list.",
					},
					map[string]any{
						"name": "ids",
						"title": "Ids",
						"type": "`$ARRAY`",
						"short": "Suppression IDs to remove from the suppression list.",
					},
				},
				"name": "batch_remove_suppressions_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/suppressions/batch/remove",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"lit": "batch",
									},
									map[string]any{
										"lit": "remove",
									},
								},
								"parts": []any{
									"suppressions",
									"batch",
									"remove",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"broadcast": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Deprecated: use `segment_id` instead.",
						"deprecated": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the broadcast was created.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "The email address of the sender.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The HTML version of the broadcast content.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the broadcast.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of the broadcast.",
					},
					map[string]any{
						"name": "preview_text",
						"title": "Preview Text",
						"type": "`$STRING`",
						"short": "The preview text of the email.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "The email addresses to which replies should be sent.",
					},
					map[string]any{
						"name": "scheduled_at",
						"title": "Scheduled At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the broadcast is scheduled to be sent.",
					},
					map[string]any{
						"name": "segment_id",
						"title": "Segment Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Unique identifier of the segment this broadcast will be sent to.",
					},
					map[string]any{
						"name": "sent_at",
						"title": "Sent At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the broadcast was sent.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The status of the broadcast.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "The subject line of the email.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The plain text version of the broadcast content.",
					},
					map[string]any{
						"name": "topic_id",
						"title": "Topic Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The topic ID that the broadcast is scoped to.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "broadcast",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/broadcasts/{id}/cancel",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "cancel",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
									"cancel",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "cancel",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/broadcasts/{id}/duplicate",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "duplicate",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
									"duplicate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "duplicate",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/broadcasts/{id}/send",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "send",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
									"send",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "send",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/broadcasts",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
								},
								"parts": []any{
									"broadcasts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/broadcasts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the audience to which the contact belongs.",
						"deprecated": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the contact was created.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "Email address of the contact.",
					},
					map[string]any{
						"name": "first_name",
						"title": "First Name",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "First name of the contact.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the contact.",
					},
					map[string]any{
						"name": "last_name",
						"title": "Last Name",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Last name of the contact.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "properties",
						"title": "Properties",
						"type": "`$OBJECT`",
						"short": "A map of custom property keys and values.",
					},
					map[string]any{
						"name": "segments",
						"title": "Segments",
						"type": "`$ARRAY`",
						"short": "Array of segment IDs to add the contact to.",
					},
					map[string]any{
						"name": "topics",
						"title": "Topics",
						"type": "`$ARRAY`",
						"short": "Array of topic subscriptions for the contact.",
					},
					map[string]any{
						"name": "unsubscribed",
						"title": "Unsubscribed",
						"type": "`$BOOLEAN`",
						"short": "Indicates if the contact is unsubscribed.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "contact",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/contacts",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
								},
								"parts": []any{
									"contacts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
								},
								"parts": []any{
									"contacts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contacts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_import_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "column_map",
						"title": "Column Map",
						"type": "`$STRING`",
						"short": "JSON-encoded object mapping contact fields and custom property keys to CSV column names.",
					},
					map[string]any{
						"name": "completed_at",
						"title": "Completed At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Timestamp indicating when the contact import completed.",
					},
					map[string]any{
						"name": "counts",
						"title": "Counts",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the contact import was created.",
					},
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
						"short": "CSV file to import.",
						"format": "binary",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the contact import.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "on_conflict",
						"title": "On Conflict",
						"type": "`$STRING`",
						"short": "Strategy to use when an imported contact already exists.",
					},
					map[string]any{
						"name": "segments",
						"title": "Segments",
						"type": "`$STRING`",
						"short": "JSON-encoded array of segments to add imported contacts to.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "Current status of the contact import.",
					},
					map[string]any{
						"name": "topics",
						"title": "Topics",
						"type": "`$STRING`",
						"short": "JSON-encoded array of topic subscriptions to apply to imported contacts.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "contact_import_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/contacts/imports",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"lit": "imports",
									},
								},
								"parts": []any{
									"contacts",
									"imports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"body": map[string]any{
									"fields": []any{
										map[string]any{
											"name": "column_map",
										},
										map[string]any{
											"binary": true,
											"name": "file",
										},
										map[string]any{
											"name": "on_conflict",
										},
										map[string]any{
											"name": "segments",
										},
										map[string]any{
											"name": "topics",
										},
									},
									"kind": "multipart",
									"media": "multipart/form-data",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts/imports",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"lit": "imports",
									},
								},
								"parts": []any{
									"contacts",
									"imports",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts/imports/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"lit": "imports",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contacts",
									"imports",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_property": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the contact property was created.",
					},
					map[string]any{
						"name": "fallback_value",
						"title": "Fallback Value",
						"type": "`$ANY`",
						"short": "The default value when the property is not set for a contact.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the contact property.",
					},
					map[string]any{
						"name": "key",
						"title": "Key",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The property key.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The property type.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "contact_property",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/contact-properties",
								"segments": []any{
									map[string]any{
										"lit": "contact-properties",
									},
								},
								"parts": []any{
									"contact-properties",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contact-properties",
								"segments": []any{
									map[string]any{
										"lit": "contact-properties",
									},
								},
								"parts": []any{
									"contact-properties",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contact-properties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contact-properties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contact-properties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"contact_topics_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "contact_topics_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts/{contact_id}/topics",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "topics",
									},
								},
								"parts": []any{
									"contacts",
									"{id}",
									"topics",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"contact_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "contact_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "topics",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"create_api_key": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "domain_id",
						"title": "Domain Id",
						"type": "`$STRING`",
						"short": "Restrict an API key to send emails only from a specific domain.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The API key name.",
					},
					map[string]any{
						"name": "permission",
						"title": "Permission",
						"type": "`$STRING`",
						"short": "The API key can have full access to Resend’s API or be only restricted to send emails.",
					},
				},
				"name": "create_api_key",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/api-keys",
								"segments": []any{
									map[string]any{
										"lit": "api-keys",
									},
								},
								"parts": []any{
									"api-keys",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"create_batch_email": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "send_emails",
						"title": "Send Emails",
						"type": "`$ARRAY`",
					},
				},
				"name": "create_batch_email",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/emails/batch",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "batch",
									},
								},
								"parts": []any{
									"emails",
									"batch",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata.send_emails`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"create_broadcast_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "Use `segment_id` instead.",
						"deprecated": true,
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"req": true,
						"short": "The email address of the sender.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "The HTML version of the message.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of the broadcast.",
					},
					map[string]any{
						"name": "preview_text",
						"title": "Preview Text",
						"type": "`$STRING`",
						"short": "The preview text of the email.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "The email addresses to which replies should be sent.",
					},
					map[string]any{
						"name": "scheduled_at",
						"title": "Scheduled At",
						"type": "`$STRING`",
						"short": "Schedule time to send the broadcast.",
					},
					map[string]any{
						"name": "segment_id",
						"title": "Segment Id",
						"type": "`$STRING`",
						"req": true,
						"short": "Unique identifier of the segment this broadcast will be sent to.",
					},
					map[string]any{
						"name": "send",
						"title": "Send",
						"type": "`$BOOLEAN`",
						"short": "Whether to send the broadcast immediately or keep it as a draft.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"req": true,
						"short": "The subject line of the email.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text version of the message.",
					},
					map[string]any{
						"name": "topic_id",
						"title": "Topic Id",
						"type": "`$STRING`",
						"short": "The topic ID that the broadcast will be scoped to.",
					},
				},
				"name": "create_broadcast_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/broadcasts",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
								},
								"parts": []any{
									"broadcasts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"create_inbox_draft_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "Files to store on the draft.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ANY`",
						"short": "The BCC recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ANY`",
						"short": "The CC recipients.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The HTML body.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the draft.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "reply_to_email_id",
						"title": "Reply To Email Id",
						"type": "`$STRING`",
						"short": "The Email ID to reply to, as returned in `data[].id` by the list thread emails endpoint.",
						"format": "uuid",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The plain text body.",
					},
					map[string]any{
						"name": "thread_id",
						"title": "Thread Id",
						"type": "`$STRING`",
						"short": "The Thread ID to reply in.",
						"format": "uuid",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ANY`",
						"short": "The recipients.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "create_inbox_draft_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/drafts",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "drafts",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"drafts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"create_inbox_label_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "color",
						"title": "Color",
						"type": "`$STRING`",
						"short": "The color of the label, as a `#RRGGBB` hex color.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name of the label.",
					},
				},
				"name": "create_inbox_label_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/labels",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"domain": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "capabilities",
						"title": "Capabilities",
						"type": "`$OBJECT`",
						"short": "Configure the domain capabilities for sending and receiving emails.",
					},
					map[string]any{
						"name": "click_tracking",
						"title": "Click Tracking",
						"type": "`$BOOLEAN`",
						"short": "Whether click tracking is enabled for this domain.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the domain was created.",
					},
					map[string]any{
						"name": "custom_return_path",
						"title": "Custom Return Path",
						"type": "`$STRING`",
						"short": "For advanced use cases, choose a subdomain for the Return-Path address.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the domain.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The name of the domain.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "open_tracking",
						"title": "Open Tracking",
						"type": "`$BOOLEAN`",
						"short": "Whether open tracking is enabled for this domain.",
					},
					map[string]any{
						"name": "records",
						"title": "Records",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "region",
						"title": "Region",
						"type": "`$STRING`",
						"short": "The region where the domain is hosted.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The status of the domain.",
					},
					map[string]any{
						"name": "tls",
						"title": "Tls",
						"type": "`$STRING`",
						"short": "TLS mode.",
					},
					map[string]any{
						"name": "tracking_subdomain",
						"title": "Tracking Subdomain",
						"type": "`$STRING`",
						"short": "The subdomain used for click and open tracking.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "domain",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/domains/{domain_id}/verify",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "verify",
									},
								},
								"parts": []any{
									"domains",
									"{id}",
									"verify",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "verify",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/domains",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/domains",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
								},
								"parts": []any{
									"domains",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"domains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"domains",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"domain_claim": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "blocked_reason",
						"title": "Blocked Reason",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Why the claim is currently blocked, if applicable.",
					},
					map[string]any{
						"name": "click_tracking",
						"title": "Click Tracking",
						"type": "`$BOOLEAN`",
						"short": "Track clicks within the body of each HTML email.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the claim was created.",
					},
					map[string]any{
						"name": "custom_return_path",
						"title": "Custom Return Path",
						"type": "`$STRING`",
						"short": "For advanced use cases, choose a subdomain for the Return-Path address.",
					},
					map[string]any{
						"name": "domain_id",
						"title": "Domain Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The ID of the placeholder domain created for the claim.",
					},
					map[string]any{
						"name": "expires_at",
						"title": "Expires At",
						"type": "`$STRING`",
						"short": "The date and time the claim expires if not verified.",
					},
					map[string]any{
						"name": "failure_reason",
						"title": "Failure Reason",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Why the claim failed, if applicable.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the claim.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The name of the domain being claimed.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "open_tracking",
						"title": "Open Tracking",
						"type": "`$BOOLEAN`",
						"short": "Track the open rate of each email.",
					},
					map[string]any{
						"name": "record",
						"title": "Record",
						"type": "`$OBJECT`",
						"short": "The TXT record to add to your DNS to prove ownership of the claimed domain.",
					},
					map[string]any{
						"name": "region",
						"title": "Region",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The region where the claimed domain will send from.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The status of the claim.",
					},
					map[string]any{
						"name": "tracking_subdomain",
						"title": "Tracking Subdomain",
						"type": "`$STRING`",
						"short": "The subdomain to use for click and open tracking.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "domain_claim",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/domains/{domain_id}/claim/verify",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
									map[string]any{
										"lit": "claim",
									},
									map[string]any{
										"lit": "verify",
									},
								},
								"parts": []any{
									"domains",
									"{domain_id}",
									"claim",
									"verify",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"domain_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/domains/claim",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"lit": "claim",
									},
								},
								"parts": []any{
									"domains",
									"claim",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/domains/{domain_id}/claim",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "claim",
									},
								},
								"parts": []any{
									"domains",
									"{id}",
									"claim",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"domain_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.domain",
						},
					},
				},
			},
			"email": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"short": "The email addresses of the blind carbon copy recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"short": "The email addresses of the carbon copy recipients.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the email was created.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The email address of the sender.",
					},
					map[string]any{
						"name": "headers",
						"title": "Headers",
						"type": "`$OBJECT`",
						"short": "Custom headers to add to the email.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "The HTML body of the email.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the email.",
					},
					map[string]any{
						"name": "last_event",
						"title": "Last Event",
						"type": "`$STRING`",
						"short": "The status of the email.",
					},
					map[string]any{
						"name": "message_id",
						"title": "Message Id",
						"type": "`$STRING`",
						"short": "The Message-ID header value of the email.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "The email addresses to which replies should be sent.",
					},
					map[string]any{
						"name": "scheduled_at",
						"title": "Scheduled At",
						"type": "`$STRING`",
						"short": "Schedule email to be sent later.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The subject line of the email.",
					},
					map[string]any{
						"name": "tags",
						"title": "Tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text body of the email.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ARRAY`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$ANY`",
							},
						},
						"short": "Recipient email address.",
					},
					map[string]any{
						"name": "topic_id",
						"title": "Topic Id",
						"type": "`$STRING`",
						"short": "The topic ID to scope the email to.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "email",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/emails/{email_id}/cancel",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "cancel",
									},
								},
								"parts": []any{
									"emails",
									"{id}",
									"cancel",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"email_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "cancel",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/emails/{email_id}/share",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "share",
									},
								},
								"parts": []any{
									"emails",
									"{id}",
									"share",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"email_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "share",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/emails",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
								},
								"parts": []any{
									"emails",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
								},
								"parts": []any{
									"emails",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/{email_id}",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"emails",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"email_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"emails_metric": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "broadcast_id",
						"title": "Broadcast Id",
						"type": "`$STRING`",
						"short": "Present when `broadcast` is in `dimensions`.",
						"format": "uuid",
					},
					map[string]any{
						"name": "broadcast_name",
						"title": "Broadcast Name",
						"type": "`$STRING`",
						"short": "Present when `broadcast` is in `dimensions`.",
					},
					map[string]any{
						"name": "domain_id",
						"title": "Domain Id",
						"type": "`$STRING`",
						"short": "Present when `domain` is in `dimensions`.",
						"format": "uuid",
					},
					map[string]any{
						"name": "domain_name",
						"title": "Domain Name",
						"type": "`$STRING`",
						"short": "Present when `domain` is in `dimensions`.",
					},
					map[string]any{
						"name": "email_id",
						"title": "Email Id",
						"type": "`$STRING`",
						"short": "Present when `email` is in `dimensions`.",
						"format": "uuid",
					},
					map[string]any{
						"name": "period",
						"title": "Period",
						"type": "`$STRING`",
						"short": "Present when `period` is in `dimensions`.",
					},
				},
				"name": "emails_metric",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/metrics",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "metrics",
									},
								},
								"parts": []any{
									"emails",
									"metrics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "broadcast_id",
											"orig": "broadcast_id",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "dimension",
											"orig": "dimensions",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "end_date",
											"orig": "end_date",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "granularity",
											"orig": "granularity",
											"type": "`$STRING`",
											"kind": "query",
											"example": "daily",
										},
										map[string]any{
											"name": "metric",
											"orig": "metrics",
											"type": "`$ARRAY`",
											"kind": "query",
										},
										map[string]any{
											"name": "start_date",
											"orig": "start_date",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "timezone",
											"orig": "timezone",
											"type": "`$STRING`",
											"kind": "query",
											"example": "UTC",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the event was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The event ID.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"short": "The event name.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "schema",
						"title": "Schema",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$OBJECT`",
								"`$NULL`",
							},
						},
						"short": "A flat key/type map defining the event payload schema.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the event was last updated.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "event",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/events",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
								},
								"parts": []any{
									"events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/events/send",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"lit": "send",
									},
								},
								"parts": []any{
									"events",
									"send",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{
									"$action": "send",
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/events",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
								},
								"parts": []any{
									"events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/events/{identifier}",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"events",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"identifier": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbox": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "ISO 8601 timestamp when the inbox was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "domain_id",
						"title": "Domain Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the domain that owns the address.",
						"format": "uuid",
					},
					map[string]any{
						"name": "drafts",
						"title": "Drafts",
						"type": "`$INTEGER`",
						"req": true,
						"short": "The number of unsent drafts.",
					},
					map[string]any{
						"name": "email_address",
						"title": "Email Address",
						"type": "`$STRING`",
						"req": true,
						"short": "The address of the inbox.",
					},
					map[string]any{
						"name": "forwarding",
						"title": "Forwarding",
						"type": "`$BOOLEAN`",
						"short": "When `true`, Resend provisions a receiving address so you can receive mail without adding an MX record.",
					},
					map[string]any{
						"name": "from_name",
						"title": "From Name",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The name recipients see when mail is sent from this inbox.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the inbox.",
						"format": "uuid",
					},
					map[string]any{
						"name": "last_received",
						"title": "Last Received",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "ISO 8601 timestamp of the newest received email in the Inbox folder.",
						"format": "date-time",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Internal name for the inbox.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
					map[string]any{
						"name": "receiving_address",
						"title": "Receiving Address",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The address to forward mail to when forwarding is enabled.",
					},
					map[string]any{
						"name": "unread",
						"title": "Unread",
						"type": "`$INTEGER`",
						"req": true,
						"short": "The number of unread threads in the Inbox folder.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"short": "ISO 8601 timestamp when the inbox was last updated.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
								},
								"parts": []any{
									"inboxes",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
								},
								"parts": []any{
									"inboxes",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/inboxes/{inbox_id}/agent",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "agent",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
									"agent",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"$action": "agent",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbox_agent_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_agent_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/agent",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "agent",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
									"agent",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.actions`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"$action": "agent",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbox_draft_list_item": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The files stored on the draft.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The BCC recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The CC recipients.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the draft.",
						"format": "uuid",
					},
					map[string]any{
						"name": "reply_to_email_id",
						"title": "Reply To Email Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Email ID the draft replies to.",
						"format": "uuid",
					},
					map[string]any{
						"name": "snippet",
						"title": "Snippet",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "A short preview of the body.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject.",
					},
					map[string]any{
						"name": "thread_id",
						"title": "Thread Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Thread ID when the draft is a reply.",
						"format": "uuid",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The recipients.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"req": true,
						"short": "`reply` when the draft replies to an email in an existing thread.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the draft was last saved.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_draft_list_item",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/drafts",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "drafts",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
									"drafts",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"$action": "drafts",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbox_draft_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The files stored on the draft.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The BCC recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The CC recipients.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the draft was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "email_id",
						"title": "Email Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The ID of the outbound email queued when the draft was sent.",
						"format": "uuid",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The HTML body.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the draft.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
					map[string]any{
						"name": "reply_to_email_id",
						"title": "Reply To Email Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Email ID the draft replies to.",
						"format": "uuid",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The plain text body.",
					},
					map[string]any{
						"name": "thread_id",
						"title": "Thread Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Thread ID when the draft is a reply.",
						"format": "uuid",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The recipients.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"req": true,
						"short": "`reply` when the draft replies to an email in an existing thread.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the draft was last saved.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_draft_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/drafts/{draft_id}/send",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "draft_id",
									},
									map[string]any{
										"lit": "send",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"drafts",
									"{draft_id}",
									"send",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "draft_id",
											"orig": "draft_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"$action": "send",
									"exist": []any{
										"draft_id",
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/drafts/{draft_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "draft_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"drafts",
									"{draft_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "draft_id",
											"orig": "draft_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"draft_id",
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"inbox_label": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "color",
						"title": "Color",
						"type": "`$STRING`",
						"req": true,
						"short": "The color of the label, as a `#RRGGBB` hex color in uppercase.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the label was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the label.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name of the label.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_label",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/labels",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
									"labels",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"inbox_thread": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The BCC recipients of the most recent email.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The CC recipients of the most recent email.",
					},
					map[string]any{
						"name": "email_count",
						"title": "Email Count",
						"type": "`$INTEGER`",
						"req": true,
						"short": "The number of emails in the thread.",
					},
					map[string]any{
						"name": "folder",
						"title": "Folder",
						"type": "`$STRING`",
						"req": true,
						"short": "The folder the thread is in.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The sender email address.",
					},
					map[string]any{
						"name": "has_attachment",
						"title": "Has Attachment",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the most recent email in the thread has an attachment.",
					},
					map[string]any{
						"name": "has_draft",
						"title": "Has Draft",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the thread has an unsent draft.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the thread.",
						"format": "uuid",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The labels attached to the thread.",
					},
					map[string]any{
						"name": "matched_email_id",
						"title": "Matched Email Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The ID of the email in the thread that matched the search.",
						"format": "uuid",
					},
					map[string]any{
						"name": "read",
						"title": "Read",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "True only when every email in the thread is read.",
					},
					map[string]any{
						"name": "received_at",
						"title": "Received At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the thread was last active.",
						"format": "date-time",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject of the thread.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The recipients of the most recent email.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_thread",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/threads",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "threads",
									},
								},
								"parts": []any{
									"inboxes",
									"{id}",
									"threads",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"inbox_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "folder",
											"orig": "folders",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"inbox",
												"sent",
											},
											"field": true,
										},
										map[string]any{
											"name": "label",
											"orig": "labels",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"7f9c1d2e-4a6b-4c3d-8e15-9b0a7c6d5e34",
											},
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "read",
											"orig": "read",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/threads/search",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"lit": "search",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"search",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "bcc",
											"orig": "bcc",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "cc",
											"orig": "cc",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "end_date",
											"orig": "end_date",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2026-09-30",
										},
										map[string]any{
											"name": "folder",
											"orig": "folders",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"inbox",
												"sent",
											},
											"field": true,
										},
										map[string]any{
											"name": "from",
											"orig": "from",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"isabella@example.com",
												"carolina@example.com",
											},
											"field": true,
										},
										map[string]any{
											"name": "has_attachment",
											"orig": "has_attachment",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "label",
											"orig": "labels",
											"type": "`$ARRAY`",
											"kind": "query",
											"example": []any{
												"7f9c1d2e-4a6b-4c3d-8e15-9b0a7c6d5e34",
											},
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "query",
											"orig": "query",
											"type": "`$STRING`",
											"kind": "query",
											"example": "invoice -draft",
										},
										map[string]any{
											"name": "read",
											"orig": "read",
											"type": "`$BOOLEAN`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "start_date",
											"orig": "start_date",
											"type": "`$STRING`",
											"kind": "query",
											"example": "2026-09-01",
										},
										map[string]any{
											"name": "to",
											"orig": "to",
											"type": "`$ARRAY`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"inbox_thread_email": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The attachments of the message.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The BCC recipients of the message.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The CC recipients of the message.",
					},
					map[string]any{
						"name": "direction",
						"title": "Direction",
						"type": "`$STRING`",
						"req": true,
						"short": "Whether the inbox received the message (`inbound`) or sent it (`outbound`).",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"req": true,
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The HTML body.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the message.",
						"format": "uuid",
					},
					map[string]any{
						"name": "message_id",
						"title": "Message Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Message-ID header of the message.",
					},
					map[string]any{
						"name": "read",
						"title": "Read",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the message has been read.",
					},
					map[string]any{
						"name": "received_at",
						"title": "Received At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the inbox received the message, or when the inbox sent it.",
						"format": "date-time",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The Reply-To addresses of the message.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject of the message.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The plain-text body.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The recipients of the message.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_thread_email",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/emails",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "emails",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"emails",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"inbox_thread_email_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The attachments of the message.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The BCC recipients of the message.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The CC recipients of the message.",
					},
					map[string]any{
						"name": "direction",
						"title": "Direction",
						"type": "`$STRING`",
						"req": true,
						"short": "Whether the inbox received the message (`inbound`) or sent it (`outbound`).",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"req": true,
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The HTML body.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the message.",
						"format": "uuid",
					},
					map[string]any{
						"name": "message_id",
						"title": "Message Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The Message-ID header of the message.",
					},
					map[string]any{
						"name": "read",
						"title": "Read",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "Whether the message has been read.",
					},
					map[string]any{
						"name": "received_at",
						"title": "Received At",
						"type": "`$STRING`",
						"req": true,
						"short": "Timestamp when the inbox received the message, or when the inbox sent it.",
						"format": "date-time",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The Reply-To addresses of the message.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject of the message.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The plain-text body.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The recipients of the message.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_thread_email_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/forward",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
									map[string]any{
										"lit": "forward",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"emails",
									"{email_id}",
									"forward",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "forward",
									"exist": []any{
										"email_id",
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}/reply",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
									map[string]any{
										"lit": "reply",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"emails",
									"{email_id}",
									"reply",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"header": []any{
										map[string]any{
											"name": "idempotency_key",
											"orig": "Idempotency-Key",
											"type": "`$STRING`",
											"kind": "header",
										},
									},
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "reply",
									"exist": []any{
										"email_id",
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/emails/{email_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"emails",
									"{email_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email_id",
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
							"$.main.kit.entity.email",
						},
					},
				},
			},
			"inbox_thread_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "folder",
						"title": "Folder",
						"type": "`$STRING`",
						"req": true,
						"short": "The folder the thread is in.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the thread.",
						"format": "uuid",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The labels attached to the thread.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
					map[string]any{
						"name": "read",
						"title": "Read",
						"type": "`$BOOLEAN`",
						"req": true,
						"short": "True only when every email in the thread is read.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject of the thread.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "inbox_thread_response_success",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"list_attachment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "content_disposition",
						"title": "Content Disposition",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "How the attachment should be displayed.",
					},
					map[string]any{
						"name": "content_id",
						"title": "Content Id",
						"type": "`$STRING`",
						"short": "The content ID for inline attachments.",
					},
					map[string]any{
						"name": "content_type",
						"title": "Content Type",
						"type": "`$STRING`",
						"short": "The MIME type of the attachment.",
					},
					map[string]any{
						"name": "download_url",
						"title": "Download Url",
						"type": "`$STRING`",
						"short": "Signed URL to download the attachment content.",
					},
					map[string]any{
						"name": "expires_at",
						"title": "Expires At",
						"type": "`$STRING`",
						"short": "Timestamp when the download URL expires.",
						"format": "date-time",
					},
					map[string]any{
						"name": "filename",
						"title": "Filename",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The filename of the attachment.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the attachment.",
						"format": "uuid",
					},
					map[string]any{
						"name": "size",
						"title": "Size",
						"type": "`$INTEGER`",
						"short": "Size of the attachment in bytes.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_attachment",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/{email_id}/attachments",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
									map[string]any{
										"lit": "attachments",
									},
								},
								"parts": []any{
									"emails",
									"{email_id}",
									"attachments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/receiving/{email_id}/attachments",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "receiving",
									},
									map[string]any{
										"var": "receiving_id",
									},
									map[string]any{
										"lit": "attachments",
									},
								},
								"parts": []any{
									"emails",
									"receiving",
									"{receiving_id}",
									"attachments",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"email_id": "receiving_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "receiving_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"receiving_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.email",
						},
					},
				},
			},
			"list_broadcast_clicked_links_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "clicks",
						"title": "Clicks",
						"type": "`$INTEGER`",
						"short": "Total number of clicks on this URL.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "An opaque cursor for this row, used only for pagination.",
					},
					map[string]any{
						"name": "unique_clicks",
						"title": "Unique Clicks",
						"type": "`$INTEGER`",
						"short": "Number of unique clicks on this URL.",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"short": "The URL that was clicked.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_broadcast_clicked_links_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/broadcasts/{id}/clicked-links",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "broadcast_id",
									},
									map[string]any{
										"lit": "clicked-links",
									},
								},
								"parts": []any{
									"broadcasts",
									"{broadcast_id}",
									"clicked-links",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "broadcast_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "broadcast_id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"broadcast_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.broadcast",
						},
					},
				},
			},
			"list_broadcast_recipients_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "bounce_type",
						"title": "Bounce Type",
						"type": "`$STRING`",
						"short": "The type of bounce.",
					},
					map[string]any{
						"name": "clicked_links",
						"title": "Clicked Links",
						"type": "`$ARRAY`",
						"short": "The links this recipient clicked.",
					},
					map[string]any{
						"name": "contact_id",
						"title": "Contact Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The ID of the contact associated with this recipient, if one exists.",
					},
					map[string]any{
						"name": "count",
						"title": "Count",
						"type": "`$INTEGER`",
						"short": "The number of times this recipient triggered the event.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "The recipient's email address.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Opaque cursor identifying this row, used for pagination.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_broadcast_recipients_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/broadcasts/{id}/recipients",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "broadcast_id",
									},
									map[string]any{
										"lit": "recipients",
									},
								},
								"parts": []any{
									"broadcasts",
									"{broadcast_id}",
									"recipients",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "broadcast_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "broadcast_id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "bounce_type",
											"orig": "bounce_type",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "email",
											"orig": "email",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"broadcast_id",
										"type",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.broadcast",
						},
					},
				},
			},
			"list_contact_segments_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the contact was added to the segment.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the segment.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of the segment.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_contact_segments_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/contacts/{contact_id}/segments",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "contact_id",
									},
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"contacts",
									"{contact_id}",
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "contact_id",
											"orig": "contact_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"contact_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.contact",
						},
					},
				},
			},
			"list_contacts_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the contact was created.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "Email address of the contact.",
					},
					map[string]any{
						"name": "first_name",
						"title": "First Name",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "First name of the contact.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the contact.",
					},
					map[string]any{
						"name": "last_name",
						"title": "Last Name",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Last name of the contact.",
					},
					map[string]any{
						"name": "unsubscribed",
						"title": "Unsubscribed",
						"type": "`$BOOLEAN`",
						"short": "Indicates if the contact is unsubscribed.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_contacts_response_success",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/segments/{id}/contacts",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
									map[string]any{
										"lit": "contacts",
									},
								},
								"parts": []any{
									"segments",
									"{segment_id}",
									"contacts",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "segment_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "segment_id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"segment_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.segment",
						},
					},
				},
			},
			"list_webhook_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the event was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the webhook event.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The delivery status of the event for this webhook.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"short": "The type of the event.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_webhook_event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/webhooks/{webhook_id}/events",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "webhook_id",
									},
									map[string]any{
										"lit": "events",
									},
								},
								"parts": []any{
									"webhooks",
									"{webhook_id}",
									"events",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "webhook_id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"webhook_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.webhook",
						},
					},
				},
			},
			"list_webhook_event_attempt": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "http_status_code",
						"title": "Http Status Code",
						"type": "`$INTEGER`",
						"short": "The HTTP status code returned by the webhook endpoint.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the webhook event attempt.",
					},
					map[string]any{
						"name": "response",
						"title": "Response",
						"type": "`$STRING`",
						"short": "The response body returned by the webhook endpoint.",
					},
					map[string]any{
						"name": "sent_at",
						"title": "Sent At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the attempt was sent.",
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "list_webhook_event_attempt",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/webhooks/{webhook_id}/events/{event_id}/attempts",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "webhook_id",
									},
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "event_id",
									},
									map[string]any{
										"lit": "attempts",
									},
								},
								"parts": []any{
									"webhooks",
									"{webhook_id}",
									"events",
									"{event_id}",
									"attempts",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "event_id",
											"orig": "event_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "webhook_id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"event_id",
										"webhook_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.webhook",
							"$.main.kit.entity.event",
						},
					},
				},
			},
			"log": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date the log was created.",
					},
					map[string]any{
						"name": "endpoint",
						"title": "Endpoint",
						"type": "`$STRING`",
						"short": "The API endpoint that was called.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The log ID.",
						"format": "uuid",
					},
					map[string]any{
						"name": "method",
						"title": "Method",
						"type": "`$STRING`",
						"short": "The HTTP method used.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "request_body",
						"title": "Request Body",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$OBJECT`",
								"`$NULL`",
							},
						},
						"short": "The request body sent to the API.",
					},
					map[string]any{
						"name": "response_body",
						"title": "Response Body",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$OBJECT`",
								"`$NULL`",
							},
						},
						"short": "The response body returned by the API.",
					},
					map[string]any{
						"name": "response_status",
						"title": "Response Status",
						"type": "`$INTEGER`",
						"short": "The HTTP status code of the response.",
					},
					map[string]any{
						"name": "user_agent",
						"title": "User Agent",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The user agent of the request.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "log",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/logs",
								"segments": []any{
									map[string]any{
										"lit": "logs",
									},
								},
								"parts": []any{
									"logs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/logs/{log_id}",
								"segments": []any{
									map[string]any{
										"lit": "logs",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"logs",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"log_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "log_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"received_email": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": "`$ARRAY`",
						"short": "Array of attachments.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "The BCC recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "The CC recipients.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp when the email was received.",
						"format": "date-time",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "The sender email address.",
					},
					map[string]any{
						"name": "headers",
						"title": "Headers",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$OBJECT`",
								"`$NULL`",
							},
						},
						"short": "The email headers.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The HTML content of the email.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the received email.",
						"format": "uuid",
					},
					map[string]any{
						"name": "message_id",
						"title": "Message Id",
						"type": "`$STRING`",
						"short": "The unique message ID from the email headers.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "received_for",
						"title": "Received For",
						"type": "`$ARRAY`",
						"short": "The recipient addresses the email was forwarded for, taken from the `for` clause of the message's `Received` headers.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "The reply-to addresses.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "The email subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The plain text content of the email.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ARRAY`",
						"short": "The recipient email addresses.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "received_email",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/receiving",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "receiving",
									},
								},
								"parts": []any{
									"emails",
									"receiving",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/receiving/{email_id}",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "receiving",
									},
									map[string]any{
										"var": "email_id",
									},
								},
								"parts": []any{
									"emails",
									"receiving",
									"{email_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_audience_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_audience_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/audiences/{id}",
								"segments": []any{
									map[string]any{
										"lit": "audiences",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"audiences",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_broadcast_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_broadcast_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/broadcasts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_contact_from_segment_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_contact_from_segment_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/contacts/{contact_id}/segments/{segment_id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "contact_id",
									},
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "segment_id",
									},
								},
								"parts": []any{
									"contacts",
									"{contact_id}",
									"segments",
									"{segment_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "contact_id",
											"orig": "contact_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "segment_id",
											"orig": "segment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"contact_id",
										"segment_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.contact",
							"$.main.kit.entity.segment",
						},
					},
				},
			},
			"remove_contact_property_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_contact_property_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/contact-properties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contact-properties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contact-properties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_contact_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_contact_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/contacts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contacts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_event",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/events/{identifier}",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"events",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"identifier": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_inbox_draft_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_inbox_draft_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/inboxes/{inbox_id}/drafts/{draft_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "draft_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"drafts",
									"{draft_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "draft_id",
											"orig": "draft_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"draft_id",
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"remove_inbox_label_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_inbox_label_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/inboxes/{inbox_id}/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "label_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"labels",
									"{label_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "label_id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"label_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"remove_inbox_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_inbox_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/inboxes/{inbox_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"remove_inbox_thread_label_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_inbox_thread_label_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "label_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"labels",
									"{label_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "label_id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"label_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"remove_inbox_thread_response_success": map[string]any{
				"fields": []any{},
				"name": "remove_inbox_thread_response_success",
				"op": map[string]any{
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"remove_segment_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "The ID of the audience this segment belongs to.",
						"deprecated": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the segment was created.",
					},
					map[string]any{
						"name": "filter",
						"title": "Filter",
						"type": "`$OBJECT`",
						"short": "Filter conditions for the segment.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the segment.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The name of the segment.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_segment_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/segments",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/segments",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
								},
								"parts": []any{
									"segments",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/segments/{id}",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"segments",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_suppression_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the suppression was created.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "Email address to suppress.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the suppression.",
					},
					map[string]any{
						"name": "origin",
						"title": "Origin",
						"type": "`$STRING`",
						"short": "Origin of the suppression.",
					},
					map[string]any{
						"name": "source_id",
						"title": "Source Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Identifier of the event that caused the suppression, such as the email that bounced or complained.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_suppression_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/suppressions",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
								},
								"parts": []any{
									"suppressions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/suppressions",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
								},
								"parts": []any{
									"suppressions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
										map[string]any{
											"name": "origin",
											"orig": "origin",
											"type": "`$STRING`",
											"kind": "query",
											"field": true,
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/suppressions/{suppression}",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"var": "suppression",
									},
								},
								"parts": []any{
									"suppressions",
									"{suppression}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "suppression",
											"orig": "suppression",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"suppression",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.suppression",
						},
					},
				},
			},
			"remove_template_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alias",
						"title": "Alias",
						"type": "`$STRING`",
						"short": "The alias of the template.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the template was created.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"req": true,
						"short": "The HTML version of the template.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the template.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The name of the template.",
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Timestamp indicating when the template was published.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "Reply-to email addresses.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The publication status of the template.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "Email subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text version of the template.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the template was last updated.",
					},
					map[string]any{
						"name": "variables",
						"title": "Variables",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_template_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/templates",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
								},
								"parts": []any{
									"templates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/templates",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
								},
								"parts": []any{
									"templates",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/templates/{id}",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"templates",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"remove_topic_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the topic was created.",
					},
					map[string]any{
						"name": "default_subscription",
						"title": "Default Subscription",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The default subscription status for the topic.",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "A description of the topic.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the topic.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The name of the topic.",
					},
					map[string]any{
						"name": "visibility",
						"title": "Visibility",
						"type": "`$STRING`",
						"short": "The visibility of the topic.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "remove_topic_response_success",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/topics",
								"segments": []any{
									map[string]any{
										"lit": "topics",
									},
								},
								"parts": []any{
									"topics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/topics",
								"segments": []any{
									map[string]any{
										"lit": "topics",
									},
								},
								"parts": []any{
									"topics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/topics/{id}",
								"segments": []any{
									map[string]any{
										"lit": "topics",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"topics",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"retrieved_attachment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "content_disposition",
						"title": "Content Disposition",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "How the attachment should be displayed.",
					},
					map[string]any{
						"name": "content_id",
						"title": "Content Id",
						"type": "`$STRING`",
						"short": "The content ID for inline attachments.",
					},
					map[string]any{
						"name": "content_type",
						"title": "Content Type",
						"type": "`$STRING`",
						"short": "The MIME type of the attachment.",
					},
					map[string]any{
						"name": "download_url",
						"title": "Download Url",
						"type": "`$STRING`",
						"short": "Signed URL to download the attachment content.",
					},
					map[string]any{
						"name": "expires_at",
						"title": "Expires At",
						"type": "`$STRING`",
						"short": "Timestamp when the download URL expires.",
						"format": "date-time",
					},
					map[string]any{
						"name": "filename",
						"title": "Filename",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The filename of the attachment.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the attachment.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "size",
						"title": "Size",
						"type": "`$INTEGER`",
						"short": "Size of the attachment in bytes.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "retrieved_attachment",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/{email_id}/attachments/{attachment_id}",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
									map[string]any{
										"lit": "attachments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"emails",
									"{email_id}",
									"attachments",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"attachment_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "id",
											"orig": "attachment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email_id",
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/emails/receiving/{email_id}/attachments/{attachment_id}",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"lit": "receiving",
									},
									map[string]any{
										"var": "receiving_id",
									},
									map[string]any{
										"lit": "attachments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"emails",
									"receiving",
									"{receiving_id}",
									"attachments",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"attachment_id": "id",
										"email_id": "receiving_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "attachment_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "receiving_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"receiving_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.email",
						},
					},
				},
			},
			"revoke_o_auth_grant": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "client",
						"title": "Client",
						"type": "`$OBJECT`",
						"short": "The OAuth client the grant was issued to.",
					},
					map[string]any{
						"name": "client_id",
						"title": "Client Id",
						"type": "`$STRING`",
						"short": "The ID of the OAuth client the grant was issued to.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "The date and time the OAuth grant was created.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the OAuth grant.",
					},
					map[string]any{
						"name": "revoked_at",
						"title": "Revoked At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The date and time the OAuth grant was revoked, or null if it is still active.",
					},
					map[string]any{
						"name": "revoked_reason",
						"title": "Revoked Reason",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The reason the OAuth grant was revoked, or null if it is still active.",
					},
					map[string]any{
						"name": "scopes",
						"title": "Scopes",
						"type": "`$ARRAY`",
						"short": "The scopes granted to the OAuth client.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "revoke_o_auth_grant",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/oauth/grants",
								"segments": []any{
									map[string]any{
										"lit": "oauth",
									},
									map[string]any{
										"lit": "grants",
									},
								},
								"parts": []any{
									"oauth",
									"grants",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/oauth/grants/{oauth_grant_id}",
								"segments": []any{
									map[string]any{
										"lit": "oauth",
									},
									map[string]any{
										"lit": "grants",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"oauth",
									"grants",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"oauth_grant_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "oauth_grant_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"rotate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the webhook.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "signing_secret",
						"title": "Signing Secret",
						"type": "`$STRING`",
						"short": "The new secret key used to verify webhook payloads.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "rotate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/webhooks/{webhook_id}/signing-secret/rotate",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "webhook_id",
									},
									map[string]any{
										"lit": "signing-secret",
									},
									map[string]any{
										"lit": "rotate",
									},
								},
								"parts": []any{
									"webhooks",
									"{webhook_id}",
									"signing-secret",
									"rotate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "webhook_id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"webhook_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.webhook",
						},
					},
				},
			},
			"segment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "The ID of the audience this segment belongs to.",
						"deprecated": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the segment was created.",
					},
					map[string]any{
						"name": "filter",
						"title": "Filter",
						"type": "`$OBJECT`",
						"short": "Filter conditions for the segment.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the segment.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the segment.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "segment",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/segments/{id}",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"segments",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"set_inbox_thread_labels_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "data",
						"title": "Data",
						"type": "`$ARRAY`",
						"req": true,
						"short": "Every label on the thread after the change.",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The IDs of the labels the thread carries after the change.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "Type of the response object.",
					},
				},
				"name": "set_inbox_thread_labels_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PUT",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}/labels",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
									map[string]any{
										"lit": "labels",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
									"labels",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"suppression": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the suppression was created.",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "Email address that is suppressed.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the suppression.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "origin",
						"title": "Origin",
						"type": "`$STRING`",
						"short": "Origin of the suppression.",
					},
					map[string]any{
						"name": "source_id",
						"title": "Source Id",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Identifier of the event that caused the suppression, such as the email that bounced or complained.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "suppression",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/suppressions/{suppression}",
								"segments": []any{
									map[string]any{
										"lit": "suppressions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"suppressions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"suppression": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "suppression",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alias",
						"title": "Alias",
						"type": "`$STRING`",
						"short": "The alias of the template.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the template was created.",
					},
					map[string]any{
						"name": "current_version_id",
						"title": "Current Version Id",
						"type": "`$STRING`",
						"short": "The ID of the current version of the template.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "has_unpublished_versions",
						"title": "Has Unpublished Versions",
						"type": "`$BOOLEAN`",
						"short": "Indicates whether the template has unpublished versions.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "The HTML version of the template.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the template.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the template.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "published_at",
						"title": "Published At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Timestamp indicating when the template was published.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "Reply-to email addresses.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The publication status of the template.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "Email subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text version of the template.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the template was last updated.",
					},
					map[string]any{
						"name": "variables",
						"title": "Variables",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "template",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/templates/{id}/duplicate",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "duplicate",
									},
								},
								"parts": []any{
									"templates",
									"{id}",
									"duplicate",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "duplicate",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/templates/{id}/publish",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "publish",
									},
								},
								"parts": []any{
									"templates",
									"{id}",
									"publish",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "publish",
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/templates/{id}",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"templates",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"topic": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the topic was created.",
					},
					map[string]any{
						"name": "default_subscription",
						"title": "Default Subscription",
						"type": "`$STRING`",
						"short": "The default subscription status for the topic.",
					},
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "A description of the topic.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the topic.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the topic.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
					map[string]any{
						"name": "visibility",
						"title": "Visibility",
						"type": "`$STRING`",
						"short": "The visibility of the topic.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "topic",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/topics/{id}",
								"segments": []any{
									map[string]any{
										"lit": "topics",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"topics",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_api_key": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the API key.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The API key name.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_api_key",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/api-keys/{api_key_id}",
								"segments": []any{
									map[string]any{
										"lit": "api-keys",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"api-keys",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"api_key_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "api_key_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_broadcast_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "audience_id",
						"title": "Audience Id",
						"type": "`$STRING`",
						"short": "Use `segment_id` instead.",
						"deprecated": true,
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "The email address of the sender.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "The HTML version of the message.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the broadcast.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "Name of the broadcast.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type of the response.",
					},
					map[string]any{
						"name": "preview_text",
						"title": "Preview Text",
						"type": "`$STRING`",
						"short": "The preview text of the email.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "The email addresses to which replies should be sent.",
					},
					map[string]any{
						"name": "segment_id",
						"title": "Segment Id",
						"type": "`$STRING`",
						"short": "Unique identifier of the segment this broadcast will be sent to.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "The subject line of the email.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text version of the message.",
					},
					map[string]any{
						"name": "topic_id",
						"title": "Topic Id",
						"type": "`$STRING`",
						"short": "The topic ID that the broadcast will be scoped to.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_broadcast_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/broadcasts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "broadcasts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"broadcasts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_contact_property_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "fallback_value",
						"title": "Fallback Value",
						"type": "`$ANY`",
						"short": "The default value to use when the property is not set for a contact.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the contact property.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_contact_property_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/contact-properties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contact-properties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contact-properties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_contact_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"short": "Email address of the contact.",
					},
					map[string]any{
						"name": "first_name",
						"title": "First Name",
						"type": "`$STRING`",
						"short": "First name of the contact.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "Unique identifier for the updated contact.",
					},
					map[string]any{
						"name": "last_name",
						"title": "Last Name",
						"type": "`$STRING`",
						"short": "Last name of the contact.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "properties",
						"title": "Properties",
						"type": "`$OBJECT`",
						"short": "A map of custom property keys and values to update.",
					},
					map[string]any{
						"name": "unsubscribed",
						"title": "Unsubscribed",
						"type": "`$BOOLEAN`",
						"short": "The Contact's global subscription status.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_contact_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/contacts/{id}",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"contacts",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_contact_topics_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "contact_id",
						"title": "Contact Id",
						"type": "`$STRING`",
						"short": "The ID of the contact.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
					map[string]any{
						"name": "topics",
						"title": "Topics",
						"type": "`$ARRAY`",
						"op": map[string]any{
							"update": map[string]any{
								"req": true,
								"type": "`$ARRAY`",
							},
						},
						"short": "Array of updated topic subscriptions.",
					},
				},
				"name": "update_contact_topics_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/contacts/{contact_id}/topics",
								"segments": []any{
									map[string]any{
										"lit": "contacts",
									},
									map[string]any{
										"var": "contact_id",
									},
									map[string]any{
										"lit": "topics",
									},
								},
								"parts": []any{
									"contacts",
									"{contact_id}",
									"topics",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "contact_id",
											"orig": "contact_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"contact_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.contact",
						},
					},
				},
			},
			"update_domain_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "capabilities",
						"title": "Capabilities",
						"type": "`$OBJECT`",
						"short": "Configure the domain capabilities for sending and receiving emails.",
					},
					map[string]any{
						"name": "click_tracking",
						"title": "Click Tracking",
						"type": "`$BOOLEAN`",
						"short": "Track clicks within the body of each HTML email.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the updated domain.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type representing the updated domain.",
					},
					map[string]any{
						"name": "open_tracking",
						"title": "Open Tracking",
						"type": "`$BOOLEAN`",
						"short": "Track the open rate of each email.",
					},
					map[string]any{
						"name": "tls",
						"title": "Tls",
						"type": "`$STRING`",
						"short": "enforced | opportunistic.",
					},
					map[string]any{
						"name": "tracking_subdomain",
						"title": "Tracking Subdomain",
						"type": "`$STRING`",
						"short": "The subdomain to use for click and open tracking.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_domain_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/domains/{domain_id}",
								"segments": []any{
									map[string]any{
										"lit": "domains",
									},
									map[string]any{
										"var": "domain_id",
									},
								},
								"parts": []any{
									"domains",
									"{domain_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "domain_id",
											"orig": "domain_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"domain_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.domain",
						},
					},
				},
			},
			"update_email_option": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "scheduled_at",
						"title": "Scheduled At",
						"type": "`$STRING`",
						"short": "Schedule email to be sent later.",
					},
				},
				"name": "update_email_option",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/emails/{email_id}",
								"segments": []any{
									map[string]any{
										"lit": "emails",
									},
									map[string]any{
										"var": "email_id",
									},
								},
								"parts": []any{
									"emails",
									"{email_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "email_id",
											"orig": "email_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.email",
						},
					},
				},
			},
			"update_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the updated event.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "Type of the response object.",
					},
					map[string]any{
						"name": "schema",
						"title": "Schema",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$OBJECT`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "A flat key/type map defining the event payload schema.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_event",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/events/{identifier}",
								"segments": []any{
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"events",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"identifier": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "identifier",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_inbox_draft_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "attachments",
						"title": "Attachments",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "Files to store on the draft.",
					},
					map[string]any{
						"name": "bcc",
						"title": "Bcc",
						"type": "`$ANY`",
						"short": "The BCC recipients.",
					},
					map[string]any{
						"name": "cc",
						"title": "Cc",
						"type": "`$ANY`",
						"short": "The CC recipients.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The HTML body.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the draft.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "The plain text body.",
					},
					map[string]any{
						"name": "to",
						"title": "To",
						"type": "`$ANY`",
						"short": "The recipients.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_inbox_draft_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/inboxes/{inbox_id}/drafts/{draft_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "drafts",
									},
									map[string]any{
										"var": "draft_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"drafts",
									"{draft_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "draft_id",
											"orig": "draft_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"draft_id",
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"update_inbox_label_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "color",
						"title": "Color",
						"type": "`$STRING`",
						"short": "A new color for the label, as a `#RRGGBB` hex color.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the updated label.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "A new name for the label.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "Type of the response object.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_inbox_label_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/inboxes/{inbox_id}/labels/{label_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "labels",
									},
									map[string]any{
										"var": "label_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"labels",
									"{label_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "label_id",
											"orig": "label_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"label_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"update_inbox_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "from_name",
						"title": "From Name",
						"type": "`$STRING`",
						"short": "The name recipients see when mail is sent from this inbox.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the inbox.",
						"format": "uuid",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "New internal name for the inbox.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_inbox_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/inboxes/{inbox_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"update_inbox_thread_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "folder",
						"title": "Folder",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The folder the thread is in.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"req": true,
						"short": "The ID of the thread.",
						"format": "uuid",
					},
					map[string]any{
						"name": "labels",
						"title": "Labels",
						"type": "`$ARRAY`",
						"req": true,
						"short": "The labels attached to the thread.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"req": true,
						"short": "The object type.",
					},
					map[string]any{
						"name": "read",
						"title": "Read",
						"type": "`$BOOLEAN`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$BOOLEAN`",
							},
						},
						"short": "True only when every email in the thread is read.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"req": true,
						"short": "The subject of the thread.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_inbox_thread_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/inboxes/{inbox_id}/threads/{thread_id}",
								"segments": []any{
									map[string]any{
										"lit": "inboxes",
									},
									map[string]any{
										"var": "inbox_id",
									},
									map[string]any{
										"lit": "threads",
									},
									map[string]any{
										"var": "thread_id",
									},
								},
								"parts": []any{
									"inboxes",
									"{inbox_id}",
									"threads",
									"{thread_id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "inbox_id",
											"orig": "inbox_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
											"example": "b3e2b2b6-3f0e-4c8e-9ad3-2f43a1e2c7f1",
										},
										map[string]any{
											"name": "thread_id",
											"orig": "thread_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"inbox_id",
										"thread_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.inbox",
						},
					},
				},
			},
			"update_segment_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the segment.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"short": "The name of the segment.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_segment_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/segments/{id}",
								"segments": []any{
									map[string]any{
										"lit": "segments",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"segments",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_template_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "alias",
						"title": "Alias",
						"type": "`$STRING`",
						"short": "The alias of the template.",
					},
					map[string]any{
						"name": "from",
						"title": "From",
						"type": "`$STRING`",
						"short": "Sender email address.",
					},
					map[string]any{
						"name": "html",
						"title": "Html",
						"type": "`$STRING`",
						"short": "The HTML version of the template.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the template.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the template.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type of the response.",
					},
					map[string]any{
						"name": "reply_to",
						"title": "Reply To",
						"type": "`$ARRAY`",
						"short": "Reply-to email addresses.",
					},
					map[string]any{
						"name": "subject",
						"title": "Subject",
						"type": "`$STRING`",
						"short": "Email subject.",
					},
					map[string]any{
						"name": "text",
						"title": "Text",
						"type": "`$STRING`",
						"short": "The plain text version of the template.",
					},
					map[string]any{
						"name": "variables",
						"title": "Variables",
						"type": "`$ARRAY`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_template_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/templates/{id}",
								"segments": []any{
									map[string]any{
										"lit": "templates",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"templates",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_topic_response_success": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "description",
						"title": "Description",
						"type": "`$STRING`",
						"short": "A description of the topic.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the topic.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"short": "The name of the topic.",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The object type.",
					},
					map[string]any{
						"name": "visibility",
						"title": "Visibility",
						"type": "`$STRING`",
						"short": "The visibility of the topic.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_topic_response_success",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/topics/{id}",
								"segments": []any{
									map[string]any{
										"lit": "topics",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"topics",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"update_webhook": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the webhook was created.",
					},
					map[string]any{
						"name": "endpoint",
						"title": "Endpoint",
						"type": "`$STRING`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "The URL where webhook events will be sent.",
					},
					map[string]any{
						"name": "events",
						"title": "Events",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"list": map[string]any{
								"type": []any{
									"`$ONE`",
									[]any{
										"`$ARRAY`",
										"`$NULL`",
									},
								},
							},
							"update": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"short": "Array of event types to subscribe to.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the updated webhook.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The status of the webhook.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "update_webhook",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/webhooks",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
								},
								"parts": []any{
									"webhooks",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/webhooks",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
								},
								"parts": []any{
									"webhooks",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.data`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "after",
											"orig": "after",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "before",
											"orig": "before",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/webhooks/{webhook_id}",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"webhooks",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"webhook_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"usage": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "ai_credits",
						"title": "Ai Credits",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "automation_runs",
						"title": "Automation Runs",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "broadcasts",
						"title": "Broadcasts",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "contacts",
						"title": "Contacts",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "domains",
						"title": "Domains",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "emails",
						"title": "Emails",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "rate_limit",
						"title": "Rate Limit",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "segments",
						"title": "Segments",
						"type": "`$OBJECT`",
					},
				},
				"name": "usage",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/usage",
								"segments": []any{
									map[string]any{
										"lit": "usage",
									},
								},
								"parts": []any{
									"usage",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"webhook": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the webhook was created.",
					},
					map[string]any{
						"name": "endpoint",
						"title": "Endpoint",
						"type": "`$STRING`",
						"short": "The URL where webhook events are sent.",
					},
					map[string]any{
						"name": "events",
						"title": "Events",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$ARRAY`",
								"`$NULL`",
							},
						},
						"short": "Array of event types subscribed to.",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the webhook.",
						"format": "uuid",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "signing_secret",
						"title": "Signing Secret",
						"type": "`$STRING`",
						"short": "The secret key used to verify webhook payloads.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The status of the webhook.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "webhook",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/webhooks/{webhook_id}",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"webhooks",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"webhook_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/webhooks/{webhook_id}",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"webhooks",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"webhook_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"webhook_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"short": "Timestamp indicating when the event was created.",
						"format": "date-time",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "The ID of the webhook event.",
					},
					map[string]any{
						"name": "next_attempt_at",
						"title": "Next Attempt At",
						"type": []any{
							"`$ONE`",
							[]any{
								"`$STRING`",
								"`$NULL`",
							},
						},
						"short": "Timestamp of the next scheduled delivery attempt, or null when none is scheduled.",
						"format": "date-time",
					},
					map[string]any{
						"name": "object",
						"title": "Object",
						"type": "`$STRING`",
						"short": "The type of object.",
					},
					map[string]any{
						"name": "payload",
						"title": "Payload",
						"type": "`$OBJECT`",
						"short": "The event payload sent to the webhook endpoint.",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "The delivery status of the event for this webhook.",
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"short": "The type of the event.",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "webhook_event",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/webhooks/{webhook_id}/events/{event_id}/replay",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "webhook_id",
									},
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "event_id",
									},
									map[string]any{
										"lit": "replay",
									},
								},
								"parts": []any{
									"webhooks",
									"{webhook_id}",
									"events",
									"{event_id}",
									"replay",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "event_id",
											"orig": "event_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "webhook_id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "replay",
									"exist": []any{
										"event_id",
										"webhook_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/webhooks/{webhook_id}/events/{event_id}",
								"segments": []any{
									map[string]any{
										"lit": "webhooks",
									},
									map[string]any{
										"var": "webhook_id",
									},
									map[string]any{
										"lit": "events",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"webhooks",
									"{webhook_id}",
									"events",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"event_id": "id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "event_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
										map[string]any{
											"name": "webhook_id",
											"orig": "webhook_id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"webhook_id",
									},
								},
								"response": map[string]any{
									"kind": "json",
									"media": "application/json",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"$.main.kit.entity.webhook",
						},
						[]any{
							"$.main.kit.entity.webhook",
							"$.main.kit.entity.event",
						},
					},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
