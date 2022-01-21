/*
 * Hippo.Web
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct Revision {
    #[serde(rename = "created", skip_serializing_if = "Option::is_none")]
    pub created: Option<String>,
    #[serde(rename = "createdBy", skip_serializing_if = "Option::is_none")]
    pub created_by: Option<String>,
    #[serde(rename = "lastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "lastModifiedBy", skip_serializing_if = "Option::is_none")]
    pub last_modified_by: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename = "revisionNumber", skip_serializing_if = "Option::is_none")]
    pub revision_number: Option<String>,
    #[serde(rename = "appId", skip_serializing_if = "Option::is_none")]
    pub app_id: Option<String>,
    #[serde(rename = "app", skip_serializing_if = "Option::is_none")]
    pub app: Option<Box<crate::models::App>>,
    #[serde(rename = "domainEvents", skip_serializing_if = "Option::is_none")]
    pub domain_events: Option<Vec<crate::models::DomainEvent>>,
}

impl Revision {
    pub fn new() -> Revision {
        Revision {
            created: None,
            created_by: None,
            last_modified: None,
            last_modified_by: None,
            id: None,
            revision_number: None,
            app_id: None,
            app: None,
            domain_events: None,
        }
    }
}

