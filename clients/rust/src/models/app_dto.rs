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
pub struct AppDto {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "storageId", skip_serializing_if = "Option::is_none")]
    pub storage_id: Option<String>,
    #[serde(rename = "channels", skip_serializing_if = "Option::is_none")]
    pub channels: Option<Vec<crate::models::ChannelDto>>,
    #[serde(rename = "revisions", skip_serializing_if = "Option::is_none")]
    pub revisions: Option<Vec<crate::models::RevisionDto>>,
}

impl AppDto {
    pub fn new() -> AppDto {
        AppDto {
            id: None,
            name: None,
            storage_id: None,
            channels: None,
            revisions: None,
        }
    }
}


