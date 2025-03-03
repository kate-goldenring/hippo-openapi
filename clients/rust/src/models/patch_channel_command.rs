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
pub struct PatchChannelCommand {
    #[serde(rename = "channelId", skip_serializing_if = "Option::is_none")]
    pub channel_id: Option<uuid::Uuid>,
    #[serde(rename = "environmentVariables", skip_serializing_if = "Option::is_none")]
    pub environment_variables: Option<Box<crate::models::UpdateEnvironmentVariableDtoListField>>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<Box<crate::models::StringField>>,
    #[serde(rename = "domain", skip_serializing_if = "Option::is_none")]
    pub domain: Option<Box<crate::models::StringField>>,
    #[serde(rename = "revisionSelectionStrategy", skip_serializing_if = "Option::is_none")]
    pub revision_selection_strategy: Option<Box<crate::models::ChannelRevisionSelectionStrategyField>>,
    #[serde(rename = "rangeRule", skip_serializing_if = "Option::is_none")]
    pub range_rule: Option<Box<crate::models::StringField>>,
    #[serde(rename = "activeRevisionId", skip_serializing_if = "Option::is_none")]
    pub active_revision_id: Option<Box<crate::models::GuidNullableField>>,
    #[serde(rename = "certificateId", skip_serializing_if = "Option::is_none")]
    pub certificate_id: Option<Box<crate::models::GuidNullableField>>,
}

impl PatchChannelCommand {
    pub fn new() -> PatchChannelCommand {
        PatchChannelCommand {
            channel_id: None,
            environment_variables: None,
            name: None,
            domain: None,
            revision_selection_strategy: None,
            range_rule: None,
            active_revision_id: None,
            certificate_id: None,
        }
    }
}


