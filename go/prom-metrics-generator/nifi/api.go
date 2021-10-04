package nifi

import (
        "prom-metrics-generator/logger"
        "encoding/json"
        "strconv"
)

func NifiFlowAbout(
	url string,        
) string{
	// Cluster Level stats
	nifiFlowAboutUrl := "http://"+url+"/nifi-api/flow/about"

	type nifiFlowAboutResponse struct {
			About struct {
					Title            string `json:"title"`
					Version          string `json:"version"`
					URI              string `json:"uri"`
					ContentViewerURL string `json:"contentViewerUrl"`
					Timezone         string `json:"timezone"`
					BuildTag         string `json:"buildTag"`
					BuildTimestamp   string `json:"buildTimestamp"`
			} `json:"about"`
	}

	var nifiFlowAbout nifiFlowAboutResponse
	if err := json.Unmarshal(GetCall(nifiFlowAboutUrl), &nifiFlowAbout); err != nil {   // Parse []byte to go struct pointer
			clog.Error.Println("Can not unmarshal JSON")
			clog.Error.Println("NIFI cluster appears to be down.")
	}
	//clog.Info.Println(PrettyPrint(NifiFlowStatus))
	return nifiFlowAbout.About.URI
}

func NifiFlowStatus(
	url string,
	clusterUrl string,
	pgID string,
) {
	// Cluster Level stats
	nifiFlowStatusUrl := "http://"+url+"/nifi-api/flow/status"

	COMMON_FIELD := "{ nifi_uri=\""+clusterUrl+"\", pg=\""+pgID+"\"}"
	NIFI_CLUSTER_STATUS := "1"
	NIFI_ACTIVE_THREAD := "-1"
	NIFI_TOTAL_QUEUED_DATA := "-1"
	NIFI_RUNNING_COMPONENTS := "-1"
	NIFI_STOPPED_COMPONENTS := "-1"
	NIFI_INVALID_COMPONENTS := "-1"
	NIFI_DISABLED_COMPONENTS := "-1"
	NIFI_UP_TO_DATE_VERSIONED_PROCESS_GROUPS := "-1"
	NIFI_LOCALLY_MODIFIED_VERSIONED_PROCESS_GROUPS := "-1"
	NIFI_STALE_VERSIONED_PROCESS_GROUPS := "-1"
	NIFI_LOCALLY_MODIFIED_AND_STALE_VERSIONED_PROCESS_GROUPS := "-1"
	NIFI_SYNC_FAILURE_VERSIONED_PROCESS_GROUPS := "-1"

	
	type nifiFlowStatusResponse struct {
			ControllerStatus struct {
					ActiveThreadCount            int    `json:"activeThreadCount"`
					TerminatedThreadCount        int    `json:"terminatedThreadCount"`
					Queued                       string `json:"queued"`
					FlowFilesQueued              int    `json:"flowFilesQueued"`
					BytesQueued                  int    `json:"bytesQueued"`
					RunningCount                 int    `json:"runningCount"`
					StoppedCount                 int    `json:"stoppedCount"`
					InvalidCount                 int    `json:"invalidCount"`
					DisabledCount                int    `json:"disabledCount"`
					ActiveRemotePortCount        int    `json:"activeRemotePortCount"`
					InactiveRemotePortCount      int    `json:"inactiveRemotePortCount"`
					UpToDateCount                int    `json:"upToDateCount"`
					LocallyModifiedCount         int    `json:"locallyModifiedCount"`
					StaleCount                   int    `json:"staleCount"`
					LocallyModifiedAndStaleCount int    `json:"locallyModifiedAndStaleCount"`
					SyncFailureCount             int    `json:"syncFailureCount"`
			} `json:"controllerStatus"`
	}

	var nifiFlowStatus nifiFlowStatusResponse
	if err := json.Unmarshal(GetCall(nifiFlowStatusUrl), &nifiFlowStatus); err != nil {   // Parse []byte to go struct pointer
			clog.Error.Println("Can not unmarshal JSON")
			clog.Error.Println("NIFI cluster appears to be down.")
	} else{
			NIFI_CLUSTER_STATUS = "0"
			NIFI_ACTIVE_THREAD = strconv.Itoa(nifiFlowStatus.ControllerStatus.ActiveThreadCount)
			NIFI_TOTAL_QUEUED_DATA = strconv.Itoa(nifiFlowStatus.ControllerStatus.FlowFilesQueued)
			NIFI_RUNNING_COMPONENTS = strconv.Itoa(nifiFlowStatus.ControllerStatus.RunningCount)
			NIFI_STOPPED_COMPONENTS = strconv.Itoa(nifiFlowStatus.ControllerStatus.StoppedCount)
			NIFI_INVALID_COMPONENTS = strconv.Itoa(nifiFlowStatus.ControllerStatus.InvalidCount)
			NIFI_DISABLED_COMPONENTS = strconv.Itoa(nifiFlowStatus.ControllerStatus.DisabledCount)
			NIFI_UP_TO_DATE_VERSIONED_PROCESS_GROUPS = strconv.Itoa(nifiFlowStatus.ControllerStatus.UpToDateCount)
			NIFI_LOCALLY_MODIFIED_VERSIONED_PROCESS_GROUPS = strconv.Itoa(nifiFlowStatus.ControllerStatus.LocallyModifiedCount)
			NIFI_STALE_VERSIONED_PROCESS_GROUPS = strconv.Itoa(nifiFlowStatus.ControllerStatus.StaleCount)
			NIFI_LOCALLY_MODIFIED_AND_STALE_VERSIONED_PROCESS_GROUPS = strconv.Itoa(nifiFlowStatus.ControllerStatus.LocallyModifiedAndStaleCount)
			NIFI_SYNC_FAILURE_VERSIONED_PROCESS_GROUPS = strconv.Itoa(nifiFlowStatus.ControllerStatus.SyncFailureCount)
	}

	writeMetrics("nifi_cluster_status", "0=UP, 1=DOWN gauge", "gauge", "{ nifi_uri=\""+clusterUrl+"\"}", "{ nifi_uri=\""+clusterUrl+"\"}", NIFI_CLUSTER_STATUS)
	writeMetrics("nifi_active_thread", "nifi_active_thread gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_ACTIVE_THREAD)
	writeMetrics("nifi_total_queued_data", "nifi_total_queued_data gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_TOTAL_QUEUED_DATA)
	writeMetrics("nifi_running_components", "nifi_running_components gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_RUNNING_COMPONENTS)
	writeMetrics("nifi_stopped_components", "nifi_stopped_components gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_STOPPED_COMPONENTS)
	writeMetrics("nifi_invalid_components", "nifi_invalid_components gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_INVALID_COMPONENTS)
	writeMetrics("nifi_disabled_components", "nifi_disabled_components gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_DISABLED_COMPONENTS)
	writeMetrics("nifi_up_to_date_versioned_process_groups", "nifi_up_to_date_versioned_process_groups gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_UP_TO_DATE_VERSIONED_PROCESS_GROUPS)
	writeMetrics("nifi_locally_modified_versioned_process_groups", "nifi_locally_modified_versioned_process_groups gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_LOCALLY_MODIFIED_VERSIONED_PROCESS_GROUPS)
	writeMetrics("nifi_stale_versioned_process_groups", "nifi_stale_versioned_process_groups gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_STALE_VERSIONED_PROCESS_GROUPS)
	writeMetrics("nifi_locally_modified_and_stale_versioned_process_groups", "nifi_locally_modified_and_stale_versioned_process_groups gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_LOCALLY_MODIFIED_AND_STALE_VERSIONED_PROCESS_GROUPS)
	writeMetrics("nifi_sync_failure_versioned_process_groups", "nifi_sync_failure_versioned_process_groups gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_SYNC_FAILURE_VERSIONED_PROCESS_GROUPS)
}

func NifiFlowClusterSummary(
	url string,
	clusterUrl string,
){
	// Cluster Level stats
	nifiFlowClusterSummaryUrl := "http://"+url+"/nifi-api/flow/cluster/summary"

	COMMON_FIELD := "{ nifi_uri=\""+clusterUrl+"\" }"
	NIFI_CLUSTER_CONNECTED_NODE := "0"
	NIFI_CLUSTER_TOTAL_NODE := "0"
	
	type nifiFlowClusterSummaryResponse struct {
			ClusterSummary struct {
					ConnectedNodes     string `json:"connectedNodes"`
					ConnectedNodeCount int    `json:"connectedNodeCount"`
					TotalNodeCount     int    `json:"totalNodeCount"`
					ConnectedToCluster bool   `json:"connectedToCluster"`
					Clustered          bool   `json:"clustered"`
			} `json:"clusterSummary"`
	}

	var nifiFlowClusterSummary nifiFlowClusterSummaryResponse
	if err := json.Unmarshal(GetCall(nifiFlowClusterSummaryUrl), &nifiFlowClusterSummary); err != nil {   // Parse []byte to go struct pointer
			clog.Error.Println("Can not unmarshal JSON")
			clog.Error.Println("NIFI cluster appears to be down.")
	}else{
			NIFI_CLUSTER_CONNECTED_NODE = strconv.Itoa(nifiFlowClusterSummary.ClusterSummary.ConnectedNodeCount)
			NIFI_CLUSTER_TOTAL_NODE = strconv.Itoa(nifiFlowClusterSummary.ClusterSummary.TotalNodeCount)
	}
	
	writeMetrics("nifi_cluster_connected_node", "nifi_cluster_connected_node gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_CLUSTER_CONNECTED_NODE)
	writeMetrics("nifi_cluster_total_node", "nifi_cluster_total_node gauge", "gauge", COMMON_FIELD, COMMON_FIELD, NIFI_CLUSTER_TOTAL_NODE)
}

func NifiFlowPgRoot(
	url string,
	clusterUrl string,      
) string{
	// Root PG Level stats
	nifiFlowPgRootUrl := "http://"+url+"/nifi-api/flow/process-groups/root"

	type nifiFlowPgRootResponse struct {
			Permissions struct {
					CanRead  bool `json:"canRead"`
					CanWrite bool `json:"canWrite"`
			} `json:"permissions"`
			ProcessGroupFlow struct {
					ID         string `json:"id"`
					URI        string `json:"uri"`
					Breadcrumb struct {
							ID          string `json:"id"`
							Permissions struct {
									CanRead  bool `json:"canRead"`
									CanWrite bool `json:"canWrite"`
							} `json:"permissions"`
							Breadcrumb struct {
									ID   string `json:"id"`
									Name string `json:"name"`
							} `json:"breadcrumb"`
					} `json:"breadcrumb"`
					Flow struct {
							ProcessGroups []struct {
									Revision struct {
											ClientID string `json:"clientId"`
											Version  int    `json:"version"`
									} `json:"revision"`
									ID       string `json:"id"`
									URI      string `json:"uri"`
									Position struct {
											X float64 `json:"x"`
											Y float64 `json:"y"`
									} `json:"position"`
									Permissions struct {
											CanRead  bool `json:"canRead"`
											CanWrite bool `json:"canWrite"`
									} `json:"permissions"`
									Bulletins []struct {
											ID          int    `json:"id"`
											GroupID     string `json:"groupId"`
											SourceID    string `json:"sourceId"`
											Timestamp   string `json:"timestamp"`
											NodeAddress string `json:"nodeAddress"`
											CanRead     bool   `json:"canRead"`
											Bulletin    struct {
													ID          int    `json:"id"`
													NodeAddress string `json:"nodeAddress"`
													Category    string `json:"category"`
													GroupID     string `json:"groupId"`
													SourceID    string `json:"sourceId"`
													SourceName  string `json:"sourceName"`
													Level       string `json:"level"`
													Message     string `json:"message"`
													Timestamp   string `json:"timestamp"`
											} `json:"bulletin"`
									} `json:"bulletins"`
									Component struct {
											ID                   string `json:"id"`
											VersionedComponentID string `json:"versionedComponentId"`
											ParentGroupID        string `json:"parentGroupId"`
											Position             struct {
													X float64 `json:"x"`
													Y float64 `json:"y"`
											} `json:"position"`
											Name      string `json:"name"`
											Comments  string `json:"comments"`
											Variables struct {
											} `json:"variables"`
											VersionControlInformation struct {
													GroupID          string `json:"groupId"`
													RegistryID       string `json:"registryId"`
													RegistryName     string `json:"registryName"`
													BucketID         string `json:"bucketId"`
													BucketName       string `json:"bucketName"`
													FlowID           string `json:"flowId"`
													FlowName         string `json:"flowName"`
													FlowDescription  string `json:"flowDescription"`
													Version          int    `json:"version"`
													State            string `json:"state"`
													StateExplanation string `json:"stateExplanation"`
											} `json:"versionControlInformation"`
											RunningCount                 int `json:"runningCount"`
											StoppedCount                 int `json:"stoppedCount"`
											InvalidCount                 int `json:"invalidCount"`
											DisabledCount                int `json:"disabledCount"`
											ActiveRemotePortCount        int `json:"activeRemotePortCount"`
											InactiveRemotePortCount      int `json:"inactiveRemotePortCount"`
											UpToDateCount                int `json:"upToDateCount"`
											LocallyModifiedCount         int `json:"locallyModifiedCount"`
											StaleCount                   int `json:"staleCount"`
											LocallyModifiedAndStaleCount int `json:"locallyModifiedAndStaleCount"`
											SyncFailureCount             int `json:"syncFailureCount"`
											LocalInputPortCount          int `json:"localInputPortCount"`
											LocalOutputPortCount         int `json:"localOutputPortCount"`
											PublicInputPortCount         int `json:"publicInputPortCount"`
											PublicOutputPortCount        int `json:"publicOutputPortCount"`
											InputPortCount               int `json:"inputPortCount"`
											OutputPortCount              int `json:"outputPortCount"`
									} `json:"component"`
									Status struct {
											ID                 string `json:"id"`
											Name               string `json:"name"`
											StatsLastRefreshed string `json:"statsLastRefreshed"`
											AggregateSnapshot  struct {
													ID                                string        `json:"id"`
													Name                              string        `json:"name"`
													ConnectionStatusSnapshots         []interface{} `json:"connectionStatusSnapshots"`
													ProcessorStatusSnapshots          []interface{} `json:"processorStatusSnapshots"`
													RemoteProcessGroupStatusSnapshots []interface{} `json:"remoteProcessGroupStatusSnapshots"`
													InputPortStatusSnapshots          []interface{} `json:"inputPortStatusSnapshots"`
													OutputPortStatusSnapshots         []interface{} `json:"outputPortStatusSnapshots"`
													VersionedFlowState                string        `json:"versionedFlowState"`
													FlowFilesIn                       int           `json:"flowFilesIn"`
													BytesIn                           int           `json:"bytesIn"`
													Input                             string        `json:"input"`
													FlowFilesQueued                   int           `json:"flowFilesQueued"`
													BytesQueued                       int           `json:"bytesQueued"`
													Queued                            string        `json:"queued"`
													QueuedCount                       string        `json:"queuedCount"`
													QueuedSize                        string        `json:"queuedSize"`
													BytesRead                         int           `json:"bytesRead"`
													Read                              string        `json:"read"`
													BytesWritten                      int           `json:"bytesWritten"`
													Written                           string        `json:"written"`
													FlowFilesOut                      int           `json:"flowFilesOut"`
													BytesOut                          int           `json:"bytesOut"`
													Output                            string        `json:"output"`
													FlowFilesTransferred              int           `json:"flowFilesTransferred"`
													BytesTransferred                  int           `json:"bytesTransferred"`
													Transferred                       string        `json:"transferred"`
													BytesReceived                     int           `json:"bytesReceived"`
													FlowFilesReceived                 int           `json:"flowFilesReceived"`
													Received                          string        `json:"received"`
													BytesSent                         int           `json:"bytesSent"`
													FlowFilesSent                     int           `json:"flowFilesSent"`
													Sent                              string        `json:"sent"`
													ActiveThreadCount                 int           `json:"activeThreadCount"`
													TerminatedThreadCount             int           `json:"terminatedThreadCount"`
											} `json:"aggregateSnapshot"`
									} `json:"status"`
									RunningCount                 int    `json:"runningCount"`
									StoppedCount                 int    `json:"stoppedCount"`
									InvalidCount                 int    `json:"invalidCount"`
									DisabledCount                int    `json:"disabledCount"`
									ActiveRemotePortCount        int    `json:"activeRemotePortCount"`
									InactiveRemotePortCount      int    `json:"inactiveRemotePortCount"`
									VersionedFlowState           string `json:"versionedFlowState"`
									UpToDateCount                int    `json:"upToDateCount"`
									LocallyModifiedCount         int    `json:"locallyModifiedCount"`
									StaleCount                   int    `json:"staleCount"`
									LocallyModifiedAndStaleCount int    `json:"locallyModifiedAndStaleCount"`
									SyncFailureCount             int    `json:"syncFailureCount"`
									LocalInputPortCount          int    `json:"localInputPortCount"`
									LocalOutputPortCount         int    `json:"localOutputPortCount"`
									PublicInputPortCount         int    `json:"publicInputPortCount"`
									PublicOutputPortCount        int    `json:"publicOutputPortCount"`
									InputPortCount               int    `json:"inputPortCount"`
									OutputPortCount              int    `json:"outputPortCount"`
							} `json:"processGroups"`
							RemoteProcessGroups []interface{} `json:"remoteProcessGroups"`
							Processors          []interface{} `json:"processors"`
							InputPorts          []interface{} `json:"inputPorts"`
							OutputPorts         []interface{} `json:"outputPorts"`
							Connections         []interface{} `json:"connections"`
							Labels              []interface{} `json:"labels"`
							Funnels             []interface{} `json:"funnels"`
					} `json:"flow"`
					LastRefreshed string `json:"lastRefreshed"`
			} `json:"processGroupFlow"`
	}

	var nifiFlowPgRoot nifiFlowPgRootResponse
	if err := json.Unmarshal(GetCall(nifiFlowPgRootUrl), &nifiFlowPgRoot); err != nil {   // Parse []byte to go struct pointer
			clog.Error.Println("Can not unmarshal JSON")
			clog.Error.Println("NIFI cluster appears to be down.")
	}
	return nifiFlowPgRoot.ProcessGroupFlow.ID
	//clog.Info.Println(PrettyPrint(NifiFlowPgRoot))
}
