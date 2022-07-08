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
pub struct RevisionItemPage {
    #[serde(rename = "items", skip_serializing_if = "Option::is_none")]
    pub items: Option<Vec<crate::models::RevisionItem>>,
    #[serde(rename = "totalItems", skip_serializing_if = "Option::is_none")]
    pub total_items: Option<i32>,
    #[serde(rename = "pageIndex", skip_serializing_if = "Option::is_none")]
    pub page_index: Option<i32>,
    #[serde(rename = "pageSize", skip_serializing_if = "Option::is_none")]
    pub page_size: Option<i32>,
    #[serde(rename = "isLastPage", skip_serializing_if = "Option::is_none")]
    pub is_last_page: Option<bool>,
}

impl RevisionItemPage {
    pub fn new() -> RevisionItemPage {
        RevisionItemPage {
            items: None,
            total_items: None,
            page_index: None,
            page_size: None,
            is_last_page: None,
        }
    }
}


