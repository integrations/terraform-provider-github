package models
import (
    "errors"
)
// The current status of the deployment.
type PagesDeploymentStatus_status int

const (
    DEPLOYMENT_IN_PROGRESS_PAGESDEPLOYMENTSTATUS_STATUS PagesDeploymentStatus_status = iota
    SYNCING_FILES_PAGESDEPLOYMENTSTATUS_STATUS
    FINISHED_FILE_SYNC_PAGESDEPLOYMENTSTATUS_STATUS
    UPDATING_PAGES_PAGESDEPLOYMENTSTATUS_STATUS
    PURGING_CDN_PAGESDEPLOYMENTSTATUS_STATUS
    DEPLOYMENT_CANCELLED_PAGESDEPLOYMENTSTATUS_STATUS
    DEPLOYMENT_FAILED_PAGESDEPLOYMENTSTATUS_STATUS
    DEPLOYMENT_CONTENT_FAILED_PAGESDEPLOYMENTSTATUS_STATUS
    DEPLOYMENT_ATTEMPT_ERROR_PAGESDEPLOYMENTSTATUS_STATUS
    DEPLOYMENT_LOST_PAGESDEPLOYMENTSTATUS_STATUS
    SUCCEED_PAGESDEPLOYMENTSTATUS_STATUS
)

func (i PagesDeploymentStatus_status) String() string {
    return []string{"deployment_in_progress", "syncing_files", "finished_file_sync", "updating_pages", "purging_cdn", "deployment_cancelled", "deployment_failed", "deployment_content_failed", "deployment_attempt_error", "deployment_lost", "succeed"}[i]
}
func ParsePagesDeploymentStatus_status(v string) (any, error) {
    result := DEPLOYMENT_IN_PROGRESS_PAGESDEPLOYMENTSTATUS_STATUS
    switch v {
        case "deployment_in_progress":
            result = DEPLOYMENT_IN_PROGRESS_PAGESDEPLOYMENTSTATUS_STATUS
        case "syncing_files":
            result = SYNCING_FILES_PAGESDEPLOYMENTSTATUS_STATUS
        case "finished_file_sync":
            result = FINISHED_FILE_SYNC_PAGESDEPLOYMENTSTATUS_STATUS
        case "updating_pages":
            result = UPDATING_PAGES_PAGESDEPLOYMENTSTATUS_STATUS
        case "purging_cdn":
            result = PURGING_CDN_PAGESDEPLOYMENTSTATUS_STATUS
        case "deployment_cancelled":
            result = DEPLOYMENT_CANCELLED_PAGESDEPLOYMENTSTATUS_STATUS
        case "deployment_failed":
            result = DEPLOYMENT_FAILED_PAGESDEPLOYMENTSTATUS_STATUS
        case "deployment_content_failed":
            result = DEPLOYMENT_CONTENT_FAILED_PAGESDEPLOYMENTSTATUS_STATUS
        case "deployment_attempt_error":
            result = DEPLOYMENT_ATTEMPT_ERROR_PAGESDEPLOYMENTSTATUS_STATUS
        case "deployment_lost":
            result = DEPLOYMENT_LOST_PAGESDEPLOYMENTSTATUS_STATUS
        case "succeed":
            result = SUCCEED_PAGESDEPLOYMENTSTATUS_STATUS
        default:
            return 0, errors.New("Unknown PagesDeploymentStatus_status value: " + v)
    }
    return &result, nil
}
func SerializePagesDeploymentStatus_status(values []PagesDeploymentStatus_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PagesDeploymentStatus_status) isMultiValue() bool {
    return false
}
