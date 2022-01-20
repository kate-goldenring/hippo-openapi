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
pub struct RevisionDto {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename = "appId", skip_serializing_if = "Option::is_none")]
    pub app_id: Option<String>,
    #[serde(rename = "revisionNumber", skip_serializing_if = "Option::is_none")]
    pub revision_number: Option<String>,
}

impl RevisionDto {
    pub fn new() -> RevisionDto {
        RevisionDto {
            id: None,
            app_id: None,
            revision_number: None,
        }
    }
}


