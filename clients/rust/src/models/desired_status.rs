/*
 * Hippo.Web
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 * Generated by: https://openapi-generator.tech
 */


/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum DesiredStatus {
    #[serde(rename = "Running")]
    Running,
    #[serde(rename = "Dead")]
    Dead,

}

impl ToString for DesiredStatus {
    fn to_string(&self) -> String {
        match self {
            Self::Running => String::from("Running"),
            Self::Dead => String::from("Dead"),
        }
    }
}

impl Default for DesiredStatus {
    fn default() -> DesiredStatus {
        Self::Running
    }
}




