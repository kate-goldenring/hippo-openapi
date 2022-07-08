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
pub struct CertificateItem {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "publicKey")]
    pub public_key: String,
    #[serde(rename = "privateKey")]
    pub private_key: String,
    #[serde(rename = "channels")]
    pub channels: Vec<crate::models::ChannelItem>,
}

impl CertificateItem {
    pub fn new(id: String, name: String, public_key: String, private_key: String, channels: Vec<crate::models::ChannelItem>) -> CertificateItem {
        CertificateItem {
            id,
            name,
            public_key,
            private_key,
            channels,
        }
    }
}


